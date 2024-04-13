use std::sync::Arc;
use sqlx::{PgPool, Row};
use tonic::{Request, Response, Status};
use blog_proto::category_service_server::CategoryService;
use blog_proto::{CategoryExistsReply, CategoryExistsRequest, CreateCategoryReply, CreateCategoryRequest, EditCategoryReply, EditCategoryRequest, GetCategoryReply, GetCategoryRequest, ListCategoryReply, ListCategoryRequest, ToggleCategoryReply, ToggleCategoryRequest};
use blog_proto::category_exists_request::Condition;

pub struct Category {
    pool: Arc<PgPool>,
}

impl Category {
    pub fn new(pool: PgPool) -> Self {
        Self {
            pool: Arc::new(pool)
        }
    }
}

#[tonic::async_trait]
impl CategoryService for Category {
    async fn create_category(&self, request: Request<CreateCategoryRequest>) -> Result<Response<CreateCategoryReply>, Status> {
        let CreateCategoryRequest { name } = request.into_inner();
        let exists_request = tonic::Request::new(
            CategoryExistsRequest {
                condition: Some(Condition::Name(name.clone().into())),
            }
        );
        let exists_reply = self.category_exists(exists_request).await?.into_inner();
        if exists_reply.exists {
            return Err(tonic::Status::already_exists("分类已存在"));
        }
        let res = sqlx::query("INSERT INTO categories (name) VALUES ($1) RETURNING id")
            .bind(name)
            .fetch_one(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?;
        let reply = CreateCategoryReply {
            id: res.get("id")
        };
        Ok(Response::new(reply))
    }

    async fn edit_category(&self, request: Request<EditCategoryRequest>) -> Result<Response<EditCategoryReply>, Status> {
        let EditCategoryRequest { id, name } = request.into_inner();
        let row = sqlx::query("SELECT COUNT(*) FROM categories WHERE name=$1 AND id<>$2")
            .bind(&name)
            .bind(id)
            .fetch_one(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?;
        let count: i64 = row.get(0);
        if count > 0i64 {
            return Err(tonic::Status::already_exists("分类已存在"));
        }

        let rows_affected = sqlx::query("UPDATE categories SET name=$1 WHERE id=$2")
            .bind(&name)
            .bind(id)
            .execute(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?
            .rows_affected();
        let reply = EditCategoryReply {
            id,
            ok: rows_affected > 0,
        };
        Ok(Response::new(reply))
    }

    async fn list_category(&self, request: Request<ListCategoryRequest>) -> Result<Response<ListCategoryReply>, Status> {
        let ListCategoryRequest { name, is_del } = request.into_inner();
        let query = match name {
            Some(name) => {
                let name = format!("%{}%", name);
                match is_del {
                    None => sqlx::query("SELECT id,name,is_del FROM categories WHERE name ILIKE $1  ORDER BY id").bind(name),
                    Some(is_del) => {
                        sqlx::query("SELECT id,name,is_del FROM categories WHERE name ILIKE $1 AND is_del=$2 ORDER BY id")
                            .bind(name.clone())
                            .bind(is_del)
                    }
                }
            }
            None => match is_del {
                None => sqlx::query("SELECT id,name,is_del FROM categories ORDER BY id"),
                Some(is_del) => {
                    sqlx::query("SELECT id,name,is_del FROM categories WHERE is_del=$1 ORDER BY id")
                        .bind(is_del)
                }
            }
        };
        let rows = query.fetch_all(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?;
        if rows.is_empty() {
            return Err(Status::not_found("没有符合条件的分类"));
        }
        let mut categories = Vec::with_capacity(rows.len());
        for row in rows {
            categories.push(blog_proto::Category {
                id: row.get("id"),
                name: row.get("name"),
                is_del: row.get("is_del"),
            });
        }
        let reply = ListCategoryReply {
            categories
        };
        Ok(Response::new(reply))
    }

    async fn toggle_category(&self, request: Request<ToggleCategoryRequest>) -> Result<Response<ToggleCategoryReply>, Status> {
        let ToggleCategoryRequest { id } = request.into_inner();
        let row = sqlx::query("UPDATE categories SET is_del=(NOT is_del) WHERE id=$1 RETURNING is_del")
            .bind(id)
            .fetch_optional(&*self.pool)
            .await
            .map_err(|err| Status::internal(err.to_string()))?;

        if let Some(row) = row {
            return Ok(tonic::Response::new(ToggleCategoryReply {
                id,
                is_del: row.get(0),
            }));
        }
        Err(Status::not_found("分类不存在"))
    }

    async fn category_exists(&self, request: Request<CategoryExistsRequest>) -> Result<Response<CategoryExistsReply>, Status> {
        let request = request.into_inner();
        let condition = request
            .condition
            .ok_or(tonic::Status::invalid_argument("参数错误"))?;
        let query = match condition {
            Condition::Name(name) => {
                sqlx::query("select count(*) from categories where name = $1").bind(name)
            }
            Condition::Id(id) => {
                sqlx::query("select count(*) from categories where id = $1").bind(id)
            }
        };
        let row = query
            .fetch_one(&*self.pool)
            .await
            .map_err(|e| tonic::Status::internal(e.to_string()))?;

        let count: i64 = row.get(0);
        let reply = CategoryExistsReply {
            exists: count > 0,
        };
        Ok(tonic::Response::new(reply))
    }

    async fn get_category(&self, request: Request<GetCategoryRequest>) -> Result<Response<GetCategoryReply>, Status> {
        let GetCategoryRequest { id, is_del } = request.into_inner();
        let query = match is_del {
            None => sqlx::query("SELECT id,name,is_del FROM categories WHERE id=$1").bind(id),
            Some(id_del) => {
                sqlx::query("SELECT id,name,is_del FROM categories WHERE id=$1 AND is_del=$2").bind(id).bind(id_del)
            }
        };

        let row = query
            .fetch_optional(&*self.pool)
            .await
            .map_err(|e| Status::internal(e.to_string()))?;

        let reply = match row {
            None => GetCategoryReply { category: None },
            Some(row) => GetCategoryReply {
                category: Some(blog_proto::Category {
                    id: row.get("id"),
                    name: row.get("name"),
                    is_del: row.get("id_del"),
                })
            }
        };
        Ok(Response::new(reply))
    }
}