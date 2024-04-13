use axum::{
    error_handling::HandleErrorExt,
    http::StatusCode,
    routing::{get, service_method_routing as service},
    AddExtensionLayer, Router,
};
use dotenv::dotenv;
use error::AppError;
use tower_http::services::ServeDir;

use crate::model::AppState;

mod config;
mod core;
mod db;
mod error;
mod form;
mod handler;
mod html;
mod model;
mod types;

pub use types::Result;

#[tokio::main]
async fn main() {
    // 初始化日志
    if std::env::var_os("RUST_LOG").is_none() {
        std::env::set_var("RUST_LOG", "short_url=debug");
    }
    tracing_subscriber::fmt::init();

    // 解析 .env 文件
    dotenv().ok();

    let cfg = config::Config::from_env().expect("初始化配置失败");
    let pool = cfg
        .pg
        .create_pool(tokio_postgres::NoTls)
        .expect("初始化数据库连接池失败");

    let app = Router::new()
        .route("/", get(handler::index).post(handler::index_action))
        .route("/rank", get(handler::rank))
        .route("/msg", get(handler::msg))
        .route("/:id", get(handler::goto_url))
        .nest(
            "/static",
            service::get(ServeDir::new("static")).handle_error(|err| {
                (
                    StatusCode::INTERNAL_SERVER_ERROR,
                    format!("处理静态资源出错：{:?}", err),
                )
            }),
        )
        .layer(AddExtensionLayer::new(AppState {
            pool,
            short_url_cfg: cfg.short_url.clone(),
        }));

    tracing::info!("服务器监听于：{}", &cfg.web.addr);

    // 绑定到配置文件设置的地址
    axum::Server::bind(&cfg.web.addr.parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}
