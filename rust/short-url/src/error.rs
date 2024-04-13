use std::convert::Infallible;

use askama::Template;
use axum::{
    body::{Bytes, Full},
    response::{Html, IntoResponse},
};

use crate::html::MsgTemplate;

#[derive(Debug)]
pub enum AppErrorType {
    DbError,
    NotFound,
    TemplateError,
    ReservedWord,
}

#[derive(Debug)]
pub struct AppError {
    pub message: Option<String>,
    pub cause: Option<String>,
    pub error_type: AppErrorType,
}

impl AppError {
    pub fn new(message: Option<String>, cause: Option<String>, error_type: AppErrorType) -> Self {
        Self {
            message,
            cause,
            error_type,
        }
    }
    pub fn from_err(err: impl ToString, error_type: AppErrorType) -> Self {
        Self::new(None, Some(err.to_string()), error_type)
    }
    pub fn from_str(msg: &str, error_type: AppErrorType) -> Self {
        Self::new(Some(msg.to_string()), None, error_type)
    }
    pub fn db_error(err: impl ToString) -> Self {
        Self::from_err(err, AppErrorType::DbError)
    }
    pub fn not_found_with_msg(msg: &str) -> Self {
        Self::from_str(msg, AppErrorType::NotFound)
    }
    pub fn not_found() -> Self {
        Self::not_found_with_msg("没有找到符合条件的数据")
    }
    pub fn tmpl_error(err: impl ToString) -> Self {
        Self::from_err(err, AppErrorType::TemplateError)
    }
    pub fn tmpl_error_from_str(msg: &str) -> Self {
        Self::from_str(msg, AppErrorType::TemplateError)
    }
    pub fn is_not_found(&self) -> bool {
        match self.error_type {
            AppErrorType::NotFound => true,
            _ => false,
        }
    }
    pub fn reserved_word(word: &str) -> Self {
        let msg = format!("{}是保留字", word);
        Self::from_str(&msg, AppErrorType::ReservedWord)
    }
}

impl std::error::Error for AppError {}
impl std::fmt::Display for AppError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}", self)
    }
}

impl From<deadpool_postgres::PoolError> for AppError {
    fn from(err: deadpool_postgres::PoolError) -> Self {
        Self::db_error(err)
    }
}

impl From<tokio_postgres::Error> for AppError {
    fn from(err: tokio_postgres::Error) -> Self {
        Self::db_error(err)
    }
}

impl From<askama::Error> for AppError {
    fn from(err: askama::Error) -> Self {
        Self::tmpl_error(err)
    }
}

impl IntoResponse for AppError {
    type Body = Full<Bytes>;
    type BodyError = Infallible;
    fn into_response(self) -> axum::http::Response<Self::Body> {
        let msg = self.message.unwrap_or("有错误发生".to_string());
        let tmpl = MsgTemplate::err(msg.clone());
        let html = tmpl.render().unwrap_or(msg);
        Html(html).into_response()
    }
}
