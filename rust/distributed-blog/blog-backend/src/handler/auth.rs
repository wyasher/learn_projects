use std::sync::Arc;
use axum::{Extension, Form};
use axum::response::Html;
use tera::Context;
use tonic::codegen::http::{HeaderMap, StatusCode};
use blog_proto::get_admin_request::ByAuth;
use crate::form;
use crate::handler::redirect_with_cookie;
use crate::model::AppState;

pub async fn login_ui(Extension(state): Extension<Arc<AppState>>) -> Result<Html<String>, String> {
    let ctx = Context::new();
    let out = state
        .tera
        .render("login.html", &ctx)
        .map_err(|err| err.to_string())?;
    Ok(Html(out))
}

pub async fn login(Extension(state): Extension<Arc<AppState>>,
                   Form(frm): Form<form::Login>,
) -> Result<(StatusCode, HeaderMap), String> {
    let condition = blog_proto::get_admin_request::Condition::ByAuth(ByAuth {
        email: frm.email,
        password: frm.password,
    });
    let mut admin = state.admin.clone();
    let resp = admin
        .get_admin(tonic::Request::new(blog_proto::GetAdminRequest {
            condition: Some(condition),
        }))
        .await
        .map_err(|err| err.to_string())?;
    let repl = resp.into_inner();
    let logined_admin = match repl.admin {
        Some(la) => la,
        None => return Err("登录失败".to_string()),
    };
    let claims = state.jwt.new_claims(logined_admin.id, logined_admin.email);
    let token = state.jwt.token(&claims).map_err(|err| err.to_string())?;
    let cookie = format!("axum_rs_token={}", &token);
    Ok(redirect_with_cookie("/m/cate", Some(&cookie)))
}

pub async fn logout() -> Result<(StatusCode, HeaderMap), String> {
    Ok(redirect_with_cookie("/login", Some("axum_rs_token=")))
}