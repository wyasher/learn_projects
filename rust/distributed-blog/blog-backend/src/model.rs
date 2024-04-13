use tera::Tera;
use tonic::transport::Channel;
use blog_auth::Jwt;
use blog_proto::admin_service_client::AdminServiceClient;
use blog_proto::category_service_client::CategoryServiceClient;
use blog_proto::topic_service_client::TopicServiceClient;

pub struct AppState{
    pub tera: Tera,
    pub cate:CategoryServiceClient<Channel>,
    pub topic:TopicServiceClient<Channel>,
    pub admin:AdminServiceClient<Channel>,
    pub jwt: Jwt,
}