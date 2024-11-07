//! Configuration for the DVN off-chain workflow.

use alloy::primitives::Address;
use config::Config;
use eyre::{eyre, Result};
use serde::Deserialize;
use std::path::PathBuf;

const CONFIG_PATH: &str = "offchain/workers_config";

#[derive(Debug, Deserialize)]
pub struct WorkerConfig {
    /// The Websocket RPC URL to connect to the Ethereum network for the source chain.
    pub source_ws_rpc_url: String,
    /// The HTTP RPC URL to connect to the Ethereum network for the source chain.
    pub source_http_rpc_url: String,
    /// The Websocket RPC URL to connect to the Ethereum network for the target chain.
    pub target_ws_rpc_url: String,
    /// The HTTP RPC URL to connect to the Ethereum network for the target chain.
    pub target_http_rpc_url: String,
    /// The LayerZero endpoint address on the source chain.
    pub source_endpoint: Address,
    /// The LayerZero endpoint address on the target chain.
    pub target_endpoint: Address,
    /// The SendLib Ultra Light Node 302 address on the source chain.
    pub source_sendlib: Address,
    /// The ReceiveLib Ultra Light Node 302 address on the target chain.
    pub target_receivelib: Address,
    /// The Ethereum network ID of the target chain.
    pub target_network_eid: u64,
    /// The address of the source DVN. Used to check when the DVN is assigned.
    pub source_dvn: Address,
    /// NFFL Aggregator URL
    pub aggregator_url: String,
}

impl WorkerConfig {
    /// Load environment variables.
    pub fn load_from_env() -> Result<Self> {
        let path = project_root::get_project_root()?.join(PathBuf::from(CONFIG_PATH));
        let settings = Config::builder().add_source(config::File::from(path)).build()?;
        settings
            .try_deserialize::<Self>()
            .map_err(|e| eyre!("Something happened with the worker's config: {:?}", e))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn load_config_from_env() {
        let _conf = WorkerConfig::load_from_env().unwrap();
    }

    #[test]
    fn test_valid_config() {
        let conf = WorkerConfig::load_from_env().unwrap();

        assert!(conf.source_ws_rpc_url.starts_with("ws://") || conf.source_ws_rpc_url.starts_with("wss://"));
        assert!(conf.source_http_rpc_url.starts_with("http://") || conf.source_http_rpc_url.starts_with("https://"));
        assert!(conf.target_ws_rpc_url.starts_with("ws://") || conf.target_ws_rpc_url.starts_with("wss://"));
        assert!(conf.target_http_rpc_url.starts_with("http://") || conf.target_http_rpc_url.starts_with("https://"));
        assert!(conf.target_network_eid > 0);
    }
}
