use std::sync::Arc;
use chrono::{Datelike, DateTime, Local, Timelike, TimeZone};
use sqlx::{PgPool, Row};
use tonic::{Request, Response, Status};
use blog_proto::{CreateTopicReply, CreateTopicRequest, EditTopicReply, EditTopicRequest, GetTopicReply, GetTopicRequest, ListTopicReply, ListTopicRequest, ToggleTopicReply, ToggleTopicRequest};
use blog_proto::topic_service_server::TopicService;

pub struct Topic {
    pool: Arc<PgPool>,
}

impl Topic {
    pub fn new(pool: PgPool) -> Self {
        Self {
            pool: Arc::new(pool),
        }
    }
}

fn get_summary(content: &str) -> String {
    if content.len() <= 255 {
        return String::from(content);
    }
    content.chars().into_iter().take(255).collect()
}

fn dt_convert(dt: &DateTime<Local>) -> Option<prost_types::Timestamp> {
    if let Ok(dt) = prost_types::Timestamp::date_time(
        dt.year().into(),
        dt.month() as u8,
        dt.day() as u8,
        dt.hour() as u8,
        dt.minute() as u8,
        dt.second() as u8,
    ) {
        Some(dt)
    } else {
        None
    }
}

fn tm_cover(tm: Option<prost_types::Timestamp>) -> Option<DateTime<Local>> {
    match tm {
        Some(tm) => Some(Local.timestamp_opt(tm.seconds, 0).unwrap()),
        None => None,
    }
}
#[tonic::async_trait]
impl TopicService for Topic {
    async fn create_topic(&self, request: Request<CreateTopicRequest>) -> Result<Response<CreateTopicReply>, Status> {
        let CreateTopicRequest { title, category_id, content, summary } = request.into_inner();
        let summary = match summary {
            None => get_summary(&content),
            Some(content) => content,
        };
        let row = sqlx::query("INSERT INTO topics (title,category_id,content,summary) VALUES($1,$2,$3,$4) RETURNING id")
            .bind(title)
            .bind(category_id)
            .bind(content)
            .bind(summary)
            .fetch_one(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?;
        let reply = CreateTopicReply { id: row.get("id") };
        Ok(Response::new(reply))
    }

    async fn edit_topic(&self, request: Request<EditTopicRequest>) -> Result<Response<EditTopicReply>, Status> {
        let EditTopicRequest { id, title, category_id, content, summary } = request.into_inner();
        // summary 如果为空 使用content 开头内容
        let summary = match summary {
            None => get_summary(&content),
            Some(content) => content,
        };
        let rows_affected = sqlx::query("UPDATE topics SET title=$1,content=$2,summary=$3,category_id=$4 WHERE id=$5")
            .bind(title)
            .bind(content)
            .bind(summary)
            .bind(category_id)
            .bind(id)
            .execute(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?
            .rows_affected();
        Ok(Response::new(EditTopicReply {
            id,
            ok: rows_affected > 0,
        }))
    }

    async fn list_topic(&self, request: Request<ListTopicRequest>) -> Result<Response<ListTopicReply>, Status> {
        let ListTopicRequest {  page,
            category_id,
            keyword,
            is_del,
            dateline_range, } = request.into_inner();
        let page = page.unwrap_or(0);
        let page_size = 30;
        let offset = page * page_size;
        let mut start = None;
        let mut end = None;
        if let Some(dr) = dateline_range {
            start = tm_cover(dr.start);
            end = tm_cover(dr.end);
        }
        let row = sqlx::query(
            r#"
            SELECT
                COUNT(*)
            FROM
            topics
            WHERE 1=1
                AND ($1::int IS NULL OR category_id = $1::int)
                AND ($2::text IS NULL OR title ILIKE CONCAT('%',$2::text,'%'))
                AND ($3::boolean IS NULL OR is_del = $3::boolean)
                AND (
                    ($4::TIMESTAMPTZ IS NULL OR $5::TIMESTAMPTZ IS NULL)
                    OR
                    (dateline BETWEEN $4::TIMESTAMPTZ AND $5::TIMESTAMPTZ)
                )"#,
        )
            .bind(&category_id)
            .bind(&keyword)
            .bind(&is_del)
            .bind(&start)
            .bind(&end)
            .fetch_one(&*self.pool)
            .await
            .map_err(|e| Status::internal(e.to_string()))?;
        let record_total:i64 = row.get(0);
        let page_total = f64::ceil(record_total as f64 / page_size as f64) as i64;
        let rows = sqlx::query(
            r#"
        SELECT
            id,title,content,summary,is_del,category_id,dateline,hit FROM topics
         WHERE 1=1
            AND ($3::int IS NULL OR category_id = $3::int)
            AND ($4::text IS NULL OR title ILIKE CONCAT('%',$4::text,'%'))
            AND ($5::boolean IS NULL OR is_del = $5::boolean)
            AND (
                ($6::TIMESTAMPTZ IS NULL OR $7::TIMESTAMPTZ IS NULL)
                OR
                (dateline BETWEEN $6::TIMESTAMPTZ AND $7::TIMESTAMPTZ)
            )
        ORDER BY
            id DESC
        LIMIT
            $1
        OFFSET
            $2
        "#,
        )
            .bind(page_size)
            .bind(offset)
            .bind(&category_id)
            .bind(&keyword)
            .bind(&is_del)
            .bind(&start)
            .bind(&end)
            .fetch_all(&*self.pool)
            .await
            .map_err(|err| tonic::Status::internal(err.to_string()))?;
        let mut topics = Vec::with_capacity(rows.len());

        for row in rows {
            let dt: DateTime<Local> = row.get("dateline");
            let dateline = dt_convert(&dt);
            topics.push(blog_proto::Topic {
                id: row.get("id"),
                title: row.get("title"),
                category_id: row.get("category_id"),
                content: row.get("content"),
                summary: row.get("summary"),
                hit: row.get("hit"),
                is_del: row.get("is_del"),
                dateline,
            });
        }

        Ok(tonic::Response::new(ListTopicReply {
            page,
            page_size,
            topics,
            record_total,
            page_total,
        }))
    }

    async fn toggle_topic(&self, request: Request<ToggleTopicRequest>) -> Result<Response<ToggleTopicReply>, Status> {
        let ToggleTopicRequest { id } = request.into_inner();
        let row = sqlx::query("UPDATE topics SET is_del=(NOT is_del) WHERE id=$1 RETURNING is_del")
            .bind(id)
            .fetch_optional(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?;
        if row.is_none() {
            return Err(Status::not_found("文章不存在"));
        }
        Ok(Response::new(ToggleTopicReply {
            id,
            is_del: row.unwrap().get("is_del"),
        }))
    }
    async fn get_topic(&self, request: Request<GetTopicRequest>) -> Result<Response<GetTopicReply>, Status> {
        let GetTopicRequest { id, is_del, inc_hit } = request.into_inner();
        let inc_hit = inc_hit.unwrap_or(false);

        // 增加点击量
        if inc_hit {
            sqlx::query("UPDATE topics SET hit=hit+1 WHERE id=$1")
                .bind(id)
                .execute(&*self.pool)
                .await
                .map_err(|err| Status::internal(err.to_string()))?;
        }
        let query = match is_del {
            None => sqlx::query("SELECT id,title,content,summary,is_del,category_id,dateline,hit FROM topics WHERE id=$1")
                .bind(id),
            Some(is_del) => sqlx::query("SELECT id,title,content,summary,is_del,category_id,dateline,hit FROM topics WHERE id=$1 AND is_del=$2")
                .bind(id).bind(is_del)
        };
        let row = query
            .fetch_optional(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?;
        if row.is_none() {
            return Err(Status::not_found("文章不存在"));
        }
        let row = row.unwrap();
        let dt: DateTime<Local> = row.get("dateline");
        let dateline = dt_convert(&dt);
        Ok(Response::new(GetTopicReply {
            topic: Some(blog_proto::Topic {
                id: row.get("id"),
                title: row.get("title"),
                category_id: row.get("category_id"),
                content: row.get("content"),
                summary: row.get("summary"),
                hit: row.get("hit"),
                is_del: row.get("is_del"),
                dateline,
            })
        }))
    }
}
