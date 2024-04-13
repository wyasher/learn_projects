use std::sync::Arc;
use axum::{Extension, Router};
use axum::routing::get;
use tera::Tera;
use blog_proto::category_service_client::CategoryServiceClient;
use blog_proto::topic_service_client::TopicServiceClient;
use crate::model::AppState;

mod model;
mod handler;

#[tokio::main]
async fn main() {
    let addr = "0.0.0.0:39527";
    let cate = CategoryServiceClient::connect("http://127.0.0.1:19527")
        .await
        .unwrap();
    let topic = TopicServiceClient::connect("http://127.0.0.1:29527")
        .await
        .unwrap();
    let tera = Tera::new("blog-frontend/templates/**/*.html").unwrap();
    let app = Router::new()
        .route("/", get(handler::index))
        .route("/detail/:id",get(handler::detail))
        .layer(Extension(Arc::new(AppState::new(cate,topic,tera))));

    axum::Server::bind(&addr.parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();

}