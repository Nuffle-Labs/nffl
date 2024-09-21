use std::path::PathBuf;

use serde::Deserialize;
#[derive(Debug, Deserialize)]
pub struct RelayerConfig {
    pub production: bool,
    pub rpc_url: String,
    pub da_account_id: String,
    pub key_path: String,
    pub network: String,
    pub metrics_ip_port_addr: Option<String>,
}

impl RelayerConfig {
    pub fn compile_cmd(&self) -> Vec<String> {
        let mut cmd = vec!["run-args".to_string()];

        if self.production {
            cmd.push("--production".to_string());
        }

        cmd.extend_from_slice(&[
            "--key-path".to_string(),
            self.key_path.clone(),
            "--rpc-url".to_string(),
            self.rpc_url.clone(),
            "--da-account-id".to_string(),
            self.da_account_id.clone(),
            "--network".to_string(),
            self.network.clone(),
        ]);

        if let Some(metrics_addr) = &self.metrics_ip_port_addr {
            cmd.extend_from_slice(&[
                "--metrics-ip-port-address".to_string(),
                metrics_addr.clone(),
            ]);
        }

        cmd
    }
}

pub fn load_config(path: PathBuf) -> eyre::Result<RelayerConfig> {
    let config_str = std::fs::read_to_string(path)?;
    let config: RelayerConfig = serde_yaml::from_str(&config_str)?;
    Ok(config)
}