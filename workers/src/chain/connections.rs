//! Utilities related to connection with a blockchain.

use crate::{
    chain::HttpProvider,
    config::{DVNConfig, DVNEvent},
};
use alloy::{
    eips::BlockNumberOrTag,
    providers::{Provider, ProviderBuilder, RootProvider, WsConnect},
    pubsub::{PubSubFrontend, SubscriptionStream},
    rpc::types::{Filter, Log},
};
use alloy_json_abi::JsonAbi;
use eyre::Result;

/// Create the subscriptions for the DVN workflow.
pub async fn build_subscriptions(
    config: &DVNConfig,
) -> Result<(
    RootProvider<PubSubFrontend>,
    SubscriptionStream<Log>,
    SubscriptionStream<Log>,
)> {
    // Create the provider
    let rpc_url = config.ws_rpc();
    let ws = WsConnect::new(rpc_url);
    let provider = ProviderBuilder::new().on_ws(ws).await?;

    // layerzero endpoint filter
    let packet_filter = Filter::new()
        .address(config.l0_addr()?)
        .event(DVNEvent::PacketSent.as_ref())
        .from_block(BlockNumberOrTag::Latest);

    // messagelib endpoint filter
    let fee_paid_filter = Filter::new()
        .address(config.sendlib_uln302_addr()?)
        .event(DVNEvent::FeePaid.as_ref())
        .from_block(BlockNumberOrTag::Latest);

    // Subscribe to logs
    let endpoint_sub = provider.subscribe_logs(&packet_filter).await?;
    let sendlib_sub = provider.subscribe_logs(&fee_paid_filter).await?;

    // Create some streams from the subscriptions
    let endpoint_stream = endpoint_sub.into_stream();
    let sendlib_stream = sendlib_sub.into_stream();

    Ok((provider, endpoint_stream, sendlib_stream))
}

/// Load the MessageLib ABI.
pub fn get_abi_from_path(path: &str) -> Result<JsonAbi> {
    // Get the SendLib ABI
    let artifact = std::fs::read(path)?;
    let json: serde_json::Value = serde_json::from_slice(&artifact)?;
    // SAFETY: Assume `unwrap` is safe since the key has been harcoded
    let abi_value = json.get("abi").unwrap();
    let abi = serde_json::from_str(&abi_value.to_string())?;
    Ok(abi)
}

/// Construct an HTTP provider given the config.
pub fn get_http_provider(config: &DVNConfig) -> Result<HttpProvider> {
    let http_provider = ProviderBuilder::new().on_http(config.http_rpc()?);
    Ok(http_provider)
}
