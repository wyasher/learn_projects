//! 配置文件

use serde::Deserialize;

/// Web 配置
#[derive(Deserialize)]
pub struct WebConfig {
    pub addr: String,
}

#[derive(Deserialize, Clone)]
pub struct ShortUrlConfig {
    pub reserved_words: String,
    pub domain: String,
}

impl ShortUrlConfig {
    pub fn reserved_words(&self) -> Vec<&str> {
        self.reserved_words.split(',').collect()
    }
    pub fn in_reserved_words(&self, word: &str) -> bool {
        for w in self.reserved_words() {
            if w == word {
                return true;
            }
        }
        false
    }
}

/// 应用配置
#[derive(Deserialize)]
pub struct Config {
    pub web: WebConfig,
    pub pg: deadpool_postgres::Config,
    pub short_url: ShortUrlConfig,
}

impl Config {
    /// 从环境变量中初始化配置
    pub fn from_env() -> Result<Self, config::ConfigError> {
        let mut cfg = config::Config::new();
        cfg.merge(config::Environment::new())?;
        cfg.try_into()
    }
}
