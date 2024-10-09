//! Configuration for the DVN offchain workflow.

use alloy::{
    primitives::{Address, U256},
    transports::http::reqwest::Url,
};
use eyre::Result;

#[derive(Default)]
pub struct DVNConfig {
    /// The Websocket RPC URL to connect to the Ethereum network.
    ws_rpc_url: String,
    /// The HTTP RPC URL to connect to the Ethereum network.
    http_rpc_url: String,
    /// The LayerZero endpoint address.
    l0_endpoint_addr: String,
    /// The SendLib Ultra Light Node 302 address.
    sendlib_uln302_addr: String,
    /// The ReceiveLib Ultra Light Node 302 address.
    receivelib_uln302_addr: String,
    /// The SendLib Ultra Light Node 301 address.
    sendlib_uln301_addr: String,
    /// The ReceiveLib Ultra Light Node 301 address.
    receivelib_uln301_addr: String,
    /// The Ethereum network ID.
    network_id: u64,
    /// Own DVN address. Used to check when the DVN is assigned to a task.
    dvn_addr: String,
}

impl DVNConfig {
    /// Get the chain's RPC URL.
    pub fn ws_rpc(&self) -> &str {
        &self.ws_rpc_url
    }

    /// Get the chain's RPC URL.
    pub fn http_rpc(&self) -> Result<Url> {
        Ok(self.http_rpc_url.parse::<Url>()?)
    }

    /// Get the LayerZero endpoint address.
    pub fn l0_addr(&self) -> Result<Address> {
        Ok(self.l0_endpoint_addr.parse::<Address>()?)
    }

    /// Get the SendLib ULN302 address.
    pub fn sendlib_uln302_addr(&self) -> Result<Address> {
        Ok(self.sendlib_uln302_addr.parse::<Address>()?)
    }

    /// Get the ReceiveLib ULN302 address.
    pub fn receivelib_uln302_addr(&self) -> Result<Address> {
        Ok(self.receivelib_uln302_addr.parse::<Address>()?)
    }

    /// Get the SendLib ULN301 address.
    pub fn sendlib_uln301_addr(&self) -> Result<Address> {
        Ok(self.sendlib_uln301_addr.parse::<Address>()?)
    }

    /// Get the ReceiveLib ULN301 address.
    pub fn receivelib_uln301_addr(&self) -> Result<Address> {
        Ok(self.receivelib_uln301_addr.parse::<Address>()?)
    }

    /// Get the EID as U256.
    pub fn eid(&self) -> U256 {
        U256::from(self.network_id)
    }

    /// Get the DVN address.
    pub fn dvn_addr(&self) -> Result<Address> {
        Ok(self.dvn_addr.parse::<Address>()?)
    }

    /// Load environment variables.
    pub fn load_from_env() -> Result<Self> {
        dotenv::dotenv()?;

        Ok(Self {
            ws_rpc_url: std::env::var("WS_RPC_URL").unwrap_or_else(|_| Default::default()),
            http_rpc_url: std::env::var("HTTP_RPC_URL").unwrap_or_else(|_| Default::default()),
            l0_endpoint_addr: std::env::var("L0_ENDPOINT_ADDR").unwrap_or_else(|_| Default::default()),
            sendlib_uln302_addr: std::env::var("SENDLIB_ULN302_ADDR").unwrap_or_else(|_| Default::default()),
            receivelib_uln302_addr: std::env::var("RECEIVELIB_ULN302_ADDR").unwrap_or_else(|_| Default::default()),
            sendlib_uln301_addr: std::env::var("SENDLIB_ULN301_ADDR").unwrap_or_else(|_| Default::default()),
            receivelib_uln301_addr: std::env::var("RECEIVELIB_ULN301_ADDR").unwrap_or_else(|_| Default::default()),
            network_id: std::env::var("NETWORK_EID")
                .unwrap_or_else(|_| "0".to_string())
                .parse::<u64>()?,
            dvn_addr: std::env::var("DVN_ADDR").unwrap_or_else(|_| Default::default()),
        })
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
