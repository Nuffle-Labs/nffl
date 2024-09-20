use serde::Deserialize;
use std::fs;

#[derive(Debug, Deserialize)]
pub struct RelayerConfig {
    pub rpc_url: String,
    pub key_path: String,
    pub da_account_id: String,
    pub network: String,
}

pub fn load_config() -> anyhow::Result<RelayerConfig> {
    let config_str = fs::read_to_string("config.json")?;
    let config: RelayerConfig = serde_json::from_str(&config_str)?;
    Ok(config)
}