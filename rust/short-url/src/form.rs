use serde::Deserialize;

#[derive(Deserialize)]
pub struct CreateUrl {
    pub url: String,
    pub email: String,
}
#[derive(Deserialize)]
pub struct UpdateUrl {
    pub id: String,
    pub url: String,
    pub email: String,
}