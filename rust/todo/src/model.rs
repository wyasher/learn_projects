use serde::Serialize;
use tokio_pg_mapper_derive::PostgresMapper;

#[derive(Clone)]
pub struct AppState {
    // psql 连接池
    pub pool: deadpool_postgres::Pool,
}

#[derive(PostgresMapper, Serialize)]
#[pg_mapper(table = "todo_list")]
pub struct TodoList {
    pub id: i32,
    pub title: String,
}

#[derive(PostgresMapper, Serialize)]
#[pg_mapper(table = "todo_item")]
pub struct TodoListID{
    pub id: i32,
}

#[derive(PostgresMapper, Serialize)]
#[pg_mapper(table = "todo_item")]
pub struct TodoItem{
    pub id:i32,
    pub title:String,
    pub checked:bool,
    pub list_id:i32,
}

#[derive(PostgresMapper, Serialize)]
#[pg_mapper(table = "todo_item")]
pub struct TodoItemID{
    pub id:i32,
}