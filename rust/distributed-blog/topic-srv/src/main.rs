use std::env;
use sqlx::PgPool;
use blog_proto::topic_service_server::TopicServiceServer;

mod server;

#[tokio::main]
async fn main() {
    let addr = "127.0.0.1:29527";
    println!("Listening on http://{}", addr);
    let dsn = env::var("PG_DSN").unwrap_or("postgres://postgres:12345678@127.0.0.1:5432/blog".to_string());
    let pool = PgPool::connect(&dsn).await.unwrap();
    let srv = server::Topic::new(pool);
    tonic::transport::Server
    ::builder()
        .add_service(TopicServiceServer::new(srv))
        .serve(addr.parse().unwrap())
        .await
        .unwrap();
}