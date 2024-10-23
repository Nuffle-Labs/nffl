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
