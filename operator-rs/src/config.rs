use crate::NFFLNodeConfig;
use std::path::PathBuf;
use anyhow::Result;


// TODO: Add checks on the validation of the config file
pub fn load_config(path: PathBuf) -> Result<NFFLNodeConfig> {
    let config_str = std::fs::read_to_string(path)?;
    let config: NFFLNodeConfig = serde_yaml::from_str(&config_str)?;
    Ok(config)
}