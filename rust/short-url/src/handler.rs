use askama::Template;
use axum::{
    extract::{Extension, Form, Path, Query},
    http::{HeaderMap, StatusCode},
    response::Html,
};
use deadpool_postgres::Client;
use serde::Deserialize;

use crate::{
    core, db,
    error::AppError,
    form,
    html::{IndexTemplate, MsgTemplate, RankTemplate},
    model::AppState,
    types::{HandlerHtmlResult, HandlerRedirectResult, RedirectResponse},
    Result,
};

/// 跳转
fn redirect_with_msg(url: &str, args: Option<&MsgArgs>) -> RedirectResponse {
    let url = match args {
        Some(args) => {
            format!("{}?{}", url, args.to_string())
        }
        None => url.to_string(),
    };
    let mut headers = HeaderMap::new();
    headers.insert(axum::http::header::LOCATION, url.as_str().parse().unwrap());
    (StatusCode::FOUND, headers, ())
}
fn redirect(url: &str) -> RedirectResponse {
    redirect_with_msg(url, None)
}
/// 渲染模型
fn render<T: Template>(t: T) -> HandlerHtmlResult {
    let html = t.render().map_err(AppError::from)?;
    Ok(Html(html))
}

/// 从连接池中获取数据库连接
async fn get_client(state: &AppState, handler_name: &str) -> Result<Client> {
    state.pool.get().await.map_err(|err| {
        tracing::error!("{}: {:?}", handler_name, err);
        AppError::db_error(err)
    })
}

/// 记录错误
fn log_error(handler_name: String) -> Box<dyn Fn(AppError) -> AppError> {
    Box::new(move |err| {
        tracing::error!("{}: {:?}", handler_name, err);
        err
    })
}

/// 处理提交的内容
pub async fn index_action(
    Extension(state): Extension<AppState>,
    Form(cu): Form<form::CreateUrl>,
) -> HandlerRedirectResult {
    let id = core::short_url(&cu.url);
    if (&state).short_url_cfg.in_reserved_words(&id) {
        return Err(AppError::reserved_word(&id));
    };
    let handler_name = "index_action";
    let client = get_client(&state, handler_name).await?;
    let result = db::create(&client, cu, id)
        .await
        .map_err(log_error(handler_name.to_string()))?;
    let redirect_url = format!("/?id={}", result.id);
    Ok(redirect(&redirect_url))
}

#[derive(Deserialize)]
pub struct IndexArgs {
    pub id: Option<String>,
}

/// 首页
pub async fn index(
    Extension(state): Extension<AppState>,
    Query(args): Query<IndexArgs>,
) -> HandlerHtmlResult {
    let handler_name = "index";
    let tmpl = IndexTemplate {
        id: args.id.clone(),
        short_url_domain: state.short_url_cfg.domain.clone(),
    };
    render(tmpl).map_err(log_error(handler_name.to_string()))
}

/// 跳转到目标URL
pub async fn goto_url(
    Extension(state): Extension<AppState>,
    Path(id): Path<String>,
) -> HandlerRedirectResult {
    let handler_name = "goto_url";
    let client = get_client(&state, handler_name).await?;
    let result = db::goto_url(&client, id)
        .await
        .map_err(log_error(handler_name.to_string()))?;
    Ok(redirect(result.url.as_str()))
}

/// 排行
pub async fn rank(Extension(state): Extension<AppState>) -> HandlerHtmlResult {
    let handler_name = "rank";
    let client = get_client(&state, handler_name).await?;
    let result = db::rank(&client)
        .await
        .map_err(log_error(handler_name.to_string()))?;
    let tmpl = RankTemplate {
        urls: result,
        short_url_domain: state.short_url_cfg.domain.clone(),
    };
    render(tmpl).map_err(log_error(handler_name.to_string()))
}

#[derive(Deserialize)]
pub struct MsgArgs {
    pub ok: Option<String>,
    pub err: Option<String>,
    pub target: Option<String>,
}
impl ToString for MsgArgs {
    fn to_string(&self) -> String {
        let mut r: Vec<String> = vec![];
        if let Some(target) = self.target.clone() {
            r.push(format!("target={}", target));
        }
        if let Some(msg) = self.ok.clone() {
            r.push(format!("ok={}", msg));
        }
        if let Some(msg) = self.err.clone() {
            r.push(format!("err={}", msg));
        }
        r.join("&")
    }
}
impl Into<MsgTemplate> for MsgArgs {
    fn into(self) -> MsgTemplate {
        let mut tmpl = MsgTemplate::default();
        tmpl.target_url = self.target.clone();
        match self {
            MsgArgs { ok: Some(msg), .. } => {
                tmpl.is_ok = true;
                tmpl.msg = msg.clone();
            }
            MsgArgs { err: Some(msg), .. } => {
                tmpl.is_ok = false;
                tmpl.msg = msg.clone();
            }
            _ => {}
        }
        tmpl
    }
}
pub async fn msg(Query(args): Query<MsgArgs>) -> HandlerHtmlResult {
    let handler_name = "err";
    let tmpl: MsgTemplate = args.into();
    render(tmpl).map_err(log_error(handler_name.to_string()))
}
