//! Configuration for the DVN off-chain workflow.

use alloy::primitives::Address;
use config::Config;
use eyre::Result;
use serde::Deserialize;

const CONFIG_PATH: &str = "./workers_config";

#[derive(Debug, Deserialize)]
pub struct WorkerConfig {
    /// The Websocket RPC URL to connect to the Ethereum network.
    pub ws_rpc_url: String,
    /// The HTTP RPC URL to connect to the Ethereum network.
    pub http_rpc_url: String,
    /// The LayerZero endpoint address.
    pub l0_endpoint_addr: Address,
    /// The SendLib Ultra Light Node 302 address.
    pub sendlib_uln302_addr: Address,
    /// The ReceiveLib Ultra Light Node 302 address.
    pub receivelib_uln302_addr: Address,
    /// The SendLib Ultra Light Node 301 address.
    pub sendlib_uln301_addr: Address,
    /// The ReceiveLib Ultra Light Node 301 address.
    pub receivelib_uln301_addr: Address,
    /// The Ethereum network ID.
    pub target_network_eid: u64,
    /// Own DVN address. Used to check when the DVN is assigned to a task.
    pub dvn_addr: Address,
    /// NFFL Aggregator URL
    pub aggregator_url: String,
}

impl WorkerConfig {
    /// Load environment variables.
    pub fn load_from_env() -> Result<Self> {
        let settings = Config::builder()
            .add_source(config::File::with_name(CONFIG_PATH))
            .build()?;
        Ok(settings.try_deserialize::<Self>()?)
    }
}

/// Useful events for the DVN workflow.
pub enum LayerZeroEvent {
    PacketSent,
    DVNFeePaid,
    ExecutorFeePaid,
    PacketVerified,
}

impl AsRef<str> for LayerZeroEvent {
    fn as_ref(&self) -> &str {
        match self {
            LayerZeroEvent::PacketSent => "PacketSent(bytes,bytes,address)",
            LayerZeroEvent::DVNFeePaid => "DVNFeePaid(address[],address[],uint256[])",
            LayerZeroEvent::ExecutorFeePaid => "ExecutorFeePaid(address,uint256)",
            LayerZeroEvent::PacketVerified => "PacketVerified(address,bytes,uint256,bytes32)",
        }
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
        assert!(conf.ws_rpc_url.starts_with("ws://") || conf.ws_rpc_url.starts_with("wss://"));

        assert!(conf.http_rpc_url.starts_with("http://") || conf.http_rpc_url.starts_with("https://"));

        assert!(conf.target_network_eid > 0);
    }
}
