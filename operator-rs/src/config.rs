use std::path::PathBuf;
use anyhow::Result;
use serde::Deserialize;
use crate::types::NodeConfig; // Changed from crate::operator::types

pub fn load_config(path: PathBuf) -> Result<NodeConfig> {
    let config_str = std::fs::read_to_string(path)?;
    let config: NodeConfig = serde_yaml::from_str(&config_str)?;
    Ok(config)
}