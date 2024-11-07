//! Utilities related to connection with a blockchain.

use crate::{
    chain::{HttpProvider, LayerZeroEvent},
    config::WorkerConfig,
};
use alloy::{
    eips::BlockNumberOrTag,
    json_abi::JsonAbi,
    providers::{Provider, ProviderBuilder, RootProvider, WsConnect},
    pubsub::{PubSubFrontend, SubscriptionStream},
    rpc::types::{Filter, Log},
};
use eyre::{eyre, OptionExt, Result};
use std::path::PathBuf;

/// Create the subscriptions for the DVN workflow.
pub async fn build_dvn_subscriptions(
    config: &WorkerConfig,
) -> Result<(
    RootProvider<PubSubFrontend>,
    SubscriptionStream<Log>,
    SubscriptionStream<Log>,
)> {
    // Create the provider
    let ws = WsConnect::new(config.source_ws_rpc_url.clone());
    let provider = ProviderBuilder::new().on_ws(ws).await?;

    // layerzero endpoint filter
    let packet_filter = Filter::new()
        .address(config.source_endpoint)
        .event(LayerZeroEvent::PacketSent.as_ref())
        .from_block(BlockNumberOrTag::Latest);

    // messagelib endpoint filter
    let fee_paid_filter = Filter::new()
        .address(config.source_sendlib)
        .event(LayerZeroEvent::DVNFeePaid.as_ref())
        .from_block(BlockNumberOrTag::Latest);

    // Subscribe to logs
    let endpoint_sub = provider.subscribe_logs(&packet_filter).await?;
    let sendlib_sub = provider.subscribe_logs(&fee_paid_filter).await?;

    // Create some streams from the subscriptions
    let endpoint_stream = endpoint_sub.into_stream();
    let sendlib_stream = sendlib_sub.into_stream();

    Ok((provider, endpoint_stream, sendlib_stream))
}

pub async fn build_executor_subscriptions(
    config: &WorkerConfig,
) -> Result<(
    RootProvider<PubSubFrontend>,
    SubscriptionStream<Log>,
    SubscriptionStream<Log>,
    SubscriptionStream<Log>,
)> {
    // Create the provider
    let ws = WsConnect::new(&config.source_ws_rpc_url);
    let provider = ProviderBuilder::new().on_ws(ws).await?;

    // PacketSent
    let packet_sent_filter = Filter::new()
        .address(config.source_endpoint)
        .event(LayerZeroEvent::PacketSent.as_ref())
        .from_block(BlockNumberOrTag::Latest);

    let executor_fee_paid = Filter::new()
        .address(config.source_sendlib)
        .event(LayerZeroEvent::ExecutorFeePaid.as_ref())
        .from_block(BlockNumberOrTag::Latest);

    let packet_verified_filter = Filter::new()
        .address(config.source_endpoint)
        .event(LayerZeroEvent::PacketVerified.as_ref())
        .from_block(BlockNumberOrTag::Latest);

    let ps_stream = provider.subscribe_logs(&packet_sent_filter).await?.into_stream();
    let ef_stream = provider.subscribe_logs(&executor_fee_paid).await?.into_stream();
    let pv_stream = provider.subscribe_logs(&packet_verified_filter).await?.into_stream();

    Ok((provider, ps_stream, ef_stream, pv_stream))
}

/// Load the MessageLib ABI. The path must be relative to the project root.
pub fn get_abi_from_path(path: &str) -> Result<JsonAbi> {
    let path_buf = PathBuf::from(path);
    let artifact_path = project_root::get_project_root()?.join(path_buf);
    // Get the SendLib ABI
    let artifact =
        std::fs::read(artifact_path).map_err(|e| eyre!("Cannot load config for offchain worker. Error: {:?}", e))?;
    let json: serde_json::Value = serde_json::from_slice(&artifact)?;
    // SAFETY: Assume `unwrap` is safe since the key has been harcoded
    let abi_value = json.get("abi").ok_or_eyre("ABI not found in artifact")?;
    let abi = serde_json::from_str(&abi_value.to_string())?;
    Ok(abi)
}

/// Construct an HTTP provider given the config.
pub fn get_http_provider(rpc_url: &str) -> Result<HttpProvider> {
    let http_provider = ProviderBuilder::new().on_http(rpc_url.to_string().parse()?);
    Ok(http_provider)
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::io::Write;
    use tempfile::NamedTempFile;

    #[test]
    fn test_expect_to_find_all_abis() {
        get_abi_from_path("offchain/abi/ReceiveLibUln302.json").unwrap();
        get_abi_from_path("offchain/abi/SendLibUln302.json").unwrap();
        get_abi_from_path("offchain/abi/L0V2Endpoint.json").unwrap();
    }

    #[test]
    fn test_get_abi_from_path() {
        // Create a file inside of `env::temp_dir()`.
        let mut temp_file = NamedTempFile::new_in(".").unwrap();

        // Some mocked ABI info
        let data = r#"{
             "abi": [
              {
                  "type": "function",
                  "name": "transfer",
                  "inputs": [
                      {
                        "type": "address",
                        "name": "_to",
                        "internalType": "address"
                      },
                      {
                        "type": "uint256",
                        "name": "_amount",
                        "internalType": "uint256"
                      }
                  ],
                  "outputs": [],
                  "stateMutability": "nonpayable"
              }
            ]
         }"#;
        writeln!(temp_file, "{}", data).unwrap();

        get_abi_from_path(temp_file.path().to_str().unwrap()).unwrap();
    }
}
