use std::env;
use sqlx::PgPool;
use blog_proto::admin_service_server::AdminServiceServer;

mod server;

#[tokio::main]
async fn main() {
    let addr = "127.0.0.1:49527";
    println!("admin-srv run at: {}", addr);
    let dsn = env::var("PG_DSN").unwrap_or("postgres://postgres:12345678@127.0.0.1:5432/blog".to_string());
    let pool = PgPool::connect(&dsn).await.unwrap();
    let srv = server::Admin::new(pool);
    tonic::transport::Server::builder()
        .add_service(AdminServiceServer::new(srv))
        .serve(addr.parse().unwrap())
        .await
        .unwrap();
}