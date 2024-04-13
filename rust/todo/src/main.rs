use axum::{AddExtensionLayer, Router};
use axum::routing::{get, post};
use dotenv::dotenv;

mod config;
mod error;
mod response;
mod model;
mod form;
mod db;
mod handler;

pub use response::Response;
use crate::model::AppState;

type Result<T> = std::result::Result<T, error::AppError>;
#[tokio::main]
async fn main() {
    if std::env::var_os("RUST_LOG").is_none(){
        std::env::set_var("RUST_LOG","todo=debug");
    }
    tracing_subscriber::fmt::init();
    dotenv().ok();
    let cfg = config::Config::from_env().unwrap();
    tracing::info!("服务器监听于：{}", &cfg.web.addr);
    let pool = cfg.pg
        .create_pool(tokio_postgres::NoTls)
        .expect("Failed to create pool");

    let app = Router::new()
        .route("/", get(handler::usage::usage))
        .route(
            "/todo",
            get(handler::todo_list::all).post(handler::todo_list::create),
        )
        .route(
            "/todo/:list_id",
            get(handler::todo_list::find)
                .put(handler::todo_list::update)
                .delete(handler::todo_list::delete),
        )
        .route(
            "/todo/:list_id/items",
            get(handler::todo_item::all).post(handler::todo_item::create),
        )
        .route(
            "/todo/:list_id/items/:item_id",
            get(handler::todo_item::find)
                .put(handler::todo_item::check)
                .delete(handler::todo_item::delete),
        )
        .layer(AddExtensionLayer::new(AppState { pool }));

    axum::Server::bind(&cfg.web.addr.parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap()
    ;
}
