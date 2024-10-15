//! Configuration for the DVN offchain workflow.

use config::Config;
use eyre::Result;
use serde::Deserialize;

#[derive(Default, Deserialize)]
pub struct DVNConfig {
    /// The Websocket RPC URL to connect to the Ethereum network.
    pub ws_rpc_url: String,
    /// The HTTP RPC URL to connect to the Ethereum network.
    pub http_rpc_url: String,
    /// The LayerZero endpoint address.
    pub l0_endpoint_addr: String,
    /// The SendLib Ultra Light Node 302 address.
    pub sendlib_uln302_addr: String,
    /// The ReceiveLib Ultra Light Node 302 address.
    pub receivelib_uln302_addr: String,
    /// The SendLib Ultra Light Node 301 address.
    pub sendlib_uln301_addr: String,
    /// The ReceiveLib Ultra Light Node 301 address.
    pub receivelib_uln301_addr: String,
    /// The Ethereum network ID.
    pub network_id: u64,
    /// Own DVN address. Used to check when the DVN is assigned to a task.
    pub dvn_addr: String,
}

impl DVNConfig {
    /// Load environment variables.
    pub fn load_from_env() -> Result<Self> {
        let settings = Config::builder()
            .add_source(config::File::with_name("workers/config_dvn"))
            .build()?;
        Ok(settings.try_deserialize::<Self>()?)
    }
}

/// Useful events for the DVN workflow.
pub enum DVNEvent {
    PacketSent,
    FeePaid,
}

impl AsRef<str> for DVNEvent {
    fn as_ref(&self) -> &str {
        match self {
            DVNEvent::PacketSent => "PacketSent(bytes,bytes,address)",
            DVNEvent::FeePaid => "DVNFeePaid(address[],address[],uint256[])",
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn load_config_from_env() {
        let conf = DVNConfig::load_from_env().unwrap();
    }
}
