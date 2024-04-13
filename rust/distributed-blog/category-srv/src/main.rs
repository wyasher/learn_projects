use std::env;
use sqlx::PgPool;
use blog_proto::category_service_server::CategoryServiceServer;

mod server;

#[tokio::main]
async fn main() {
    let addr = "127.0.0.1:19527";
    println!("run at {}",addr);
    let dsn = env::var("PG_DSN").unwrap_or("postgres://postgres:12345678@127.0.0.1:5432/blog".to_string());
    let pool = PgPool::connect(&dsn).await.unwrap();
    let category_srv  = server::Category::new(pool);
    tonic::transport::Server::builder()
        .add_service(CategoryServiceServer::new(category_srv))
        .serve(addr.parse().unwrap())
        .await
        .unwrap();
}
