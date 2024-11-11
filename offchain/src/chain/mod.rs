//! Utilities for interacting with the blockchain.

use alloy::{
    contract::ContractInstance,
    network::Ethereum,
    providers::RootProvider,
    transports::http::{Client, Http},
};

pub mod connections;
pub mod contracts;

/// Alias for a contract instance in the Ethereum network.
pub type ContractInst = ContractInstance<Http<Client>, RootProvider<Http<Client>>, Ethereum>;

/// Alias for an HTTP provider.
pub type HttpProvider = RootProvider<Http<Client>>;

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
            LayerZeroEvent::PacketVerified => "PacketVerified((uint32,bytes32,uint64),address,bytes32)",
        }
    }
}
