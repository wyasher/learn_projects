use tokio_postgres::Client;

use crate::{
    form,
    model::{Url, UrlID, UrlTarget},
    Result,
};

pub async fn create(client: &Client, cu: form::CreateUrl, id: String) -> Result<UrlID> {
    // 是否存在
    let result = super::query_one(client, "SELECT id FROM url WHERE id=$1", &[&id]).await;
    match result {
        // 如果已存在，直接返回
        Ok(result) => return Ok(result),
        // 如果不是“未找到”的错误，直接返回
        Err(err) if !err.is_not_found() => return Err(err),
        // 如果不存在，什么也不做，继续下面的代码
        _ => {}
    };
    let result = super::query_one(
        client,
        "INSERT INTO url(id, url, email) VALUES ($1,$2,$3) RETURNING id",
        &[&id, &cu.url, &cu.email],
    )
        .await?;
    Ok(result)
}

pub async fn goto_url(client: &Client, id: String) -> Result<UrlTarget> {
    let result = super::query_one(
        client,
        "UPDATE url SET visit=visit+1 WHERE id=$1 RETURNING url",
        &[&id],
    )
        .await?;
    Ok(result)
}

pub async fn rank(client: &Client) -> Result<Vec<Url>> {
    let result = super::query(client, "SELECT id, url,email,visit,is_del FROM url WHERE  is_del=false ORDER BY visit DESC LIMIT 100", &[]).await?;
    Ok(result)
}
