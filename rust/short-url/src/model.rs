use serde::Serialize;
use tokio_pg_mapper_derive::PostgresMapper;

use crate::config;

/// 应用状态共享
#[derive(Clone)]
pub struct AppState {
    /// PostgreSQL 连接池
    pub pool: deadpool_postgres::Pool,
    pub short_url_cfg: config::ShortUrlConfig,
}

#[derive(PostgresMapper, Serialize)]
#[pg_mapper(table = "url")]
pub struct Url {
    pub id: String,
    pub url: String,
    pub email: String,
    pub visit: i32,
    pub is_del: bool,
}

#[derive(PostgresMapper, Serialize)]
#[pg_mapper(table = "url")]
pub struct UrlID {
    pub id: String,
}
#[derive(PostgresMapper, Serialize)]
#[pg_mapper(table = "url")]
pub struct UrlTarget {
    pub url: String,
}
