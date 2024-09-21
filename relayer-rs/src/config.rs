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

pub fn load_config(path: PathBuf) -> Result<RelayerConfig, eyre::Report> {
    let config_str = std::fs::read_to_string(path)?;
    let config: RelayerConfig = serde_yaml::from_str(&config_str)?;
    Ok(config)
}

#[cfg(test)]
mod tests {
    use super::*;
    use tempfile::NamedTempFile;
    use std::io::Write;

    #[test]
    fn test_relayer_config_compile_cmd() {
        let config = RelayerConfig {
            production: true,
            rpc_url: "http://example.com".to_string(),
            da_account_id: "test.near".to_string(),
            key_path: "/path/to/key".to_string(),
            network: "testnet".to_string(),
            metrics_ip_port_addr: Some("127.0.0.1:8080".to_string()),
        };

        let cmd = config.compile_cmd();

        assert_eq!(cmd[0], "run-args");
        assert!(cmd.contains(&"--production".to_string()));
        assert!(cmd.contains(&"--rpc-url".to_string()));
        assert!(cmd.contains(&"http://example.com".to_string()));
        assert!(cmd.contains(&"--da-account-id".to_string()));
        assert!(cmd.contains(&"test.near".to_string()));
        assert!(cmd.contains(&"--key-path".to_string()));
        assert!(cmd.contains(&"/path/to/key".to_string()));
        assert!(cmd.contains(&"--network".to_string()));
        assert!(cmd.contains(&"testnet".to_string()));
        assert!(cmd.contains(&"--metrics-ip-port-address".to_string()));
        assert!(cmd.contains(&"127.0.0.1:8080".to_string()));
    }

    #[test]
    fn test_load_config() -> eyre::Result<()> {
        let config_content = r#"
        production: true
        rpc_url: "http://example.com"
        da_account_id: "test.near"
        key_path: "/path/to/key"
        network: "testnet"
        metrics_ip_port_addr: "127.0.0.1:8080"
        "#;

        let temp_file = NamedTempFile::new()?;
        write!(temp_file.as_file(), "{}", config_content)?;

        let config = load_config(temp_file.path().to_path_buf())?;

        assert!(config.production);
        assert_eq!(config.rpc_url, "http://example.com");
        assert_eq!(config.da_account_id, "test.near");
        assert_eq!(config.key_path, "/path/to/key");
        assert_eq!(config.network, "testnet");
        assert_eq!(config.metrics_ip_port_addr, Some("127.0.0.1:8080".to_string()));

        Ok(())
    }
}