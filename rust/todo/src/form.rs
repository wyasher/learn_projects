use serde::Deserialize;

#[derive(Deserialize)]
pub struct CreateTodoList{
    pub title:String,
}

#[derive(Deserialize)]
pub struct UpdateTodoList{
    pub id:i32,
    pub title: String
}

#[derive(Deserialize)]
pub struct CreateTodoItem{
    pub title:String,
    pub list_id:i32,
}