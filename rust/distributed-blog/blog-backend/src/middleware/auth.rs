use std::sync::Arc;
use axum::async_trait;
use axum::extract::{FromRequest, RequestParts};
use blog_auth::Claims;
use crate::handler::cookie;
use crate::model;

pub struct Auth(Claims);

#[async_trait]
impl <B> FromRequest<B> for Auth where B:Send,{
    type Rejection = String;

    async fn from_request(req: &mut RequestParts<B>) -> Result<Self, Self::Rejection> {
        let state = req.extensions().get::<Arc<model::AppState>>().unwrap();
        let headers = req.headers();
        let claims = match cookie::get(headers, "axum_rs_token") {
            Some(token) => state
                .jwt
                .verify_and_get(&token)
                .map_err(|err| err.to_string())?,
            None => return Err("请登录".to_string()),
        };
        Ok(Self(claims))
    }
}