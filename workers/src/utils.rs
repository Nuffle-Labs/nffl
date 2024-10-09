//! Utilities for interacting with the DVN.

use crate::config::{DVNConfig, DVNEvent};
use alloy::{
    contract::{ContractInstance, Interface},
    dyn_abi::DynSolValue,
    eips::BlockNumberOrTag,
    network::Ethereum,
    primitives::U256,
    providers::{Provider, ProviderBuilder, RootProvider, WsConnect},
    pubsub::{PubSubFrontend, SubscriptionStream},
    rpc::types::{Filter, Log},
    transports::http::{Client, Http},
};
use alloy_json_abi::JsonAbi;
use eyre::Result;
use sha3::{Digest, Keccak256};
use tracing::debug;

pub type ContractInst = ContractInstance<Http<Client>, RootProvider<Http<Client>>, Ethereum>;
pub type HttpProvider = RootProvider<Http<Client>>;

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
pub fn get_sendlib_abi() -> Result<JsonAbi> {
    // Get the SendLib ABI
    let artifact = std::fs::read("./abi/SendLibrary.json")?;
    let json: serde_json::Value = serde_json::from_slice(&artifact)?;
    // SAFETY: Assume `unwrap` is safe since the key is always present
    let abi_value = json.get("abi").unwrap();
    let abi = serde_json::from_str(&abi_value.to_string())?;
    Ok(abi)
}

/// Construct an HTTP provider given the config.
pub fn get_http_provider(config: &DVNConfig) -> Result<HttpProvider> {
    let http_provider = ProviderBuilder::new().on_http(config.http_rpc()?);
    Ok(http_provider)
}

/// Create a contract instance from the ABI to interact with on-chain instance.
pub fn create_contract_instance(config: &DVNConfig, http_provider: HttpProvider, abi: JsonAbi) -> Result<ContractInst> {
    let contract: ContractInstance<Http<Client>, _, Ethereum> = ContractInstance::new(
        config.sendlib_uln302_addr()?,
        http_provider.clone(),
        Interface::new(abi),
    );
    Ok(contract)
}

/// Get the number of required confirmations by the ULN.
pub async fn get_confirmations(config: &DVNConfig, contract: &ContractInst) -> Result<U256> {
    // FIXME: there an error returned by the server:
    //     Error: server returned an error response: error code 3: execution reverted, data: "0xce2c3751"
    // which decodes (https://www.4byte.directory/signatures/?bytes4_signature=0xce2c3751) to:
    //     LZ_ULN_AtLeastOneDVN()

    debug!("Getting confirmations required by the ULN.");
    debug!("Contract address: {:?}", config.sendlib_uln302_addr()?);
    debug!("Contract address: {:?}", config.eid());

    // Call the `getUlnConfig` function on the contract
    let uln_config = contract
        .function(
            "getUlnConfig",
            &[
                DynSolValue::Address(config.sendlib_uln302_addr()?),
                DynSolValue::Uint(config.eid(), 32),
            ],
        )?
        .call()
        .await?;

    let num_confirmations = if let DynSolValue::Tuple(tupled_uint) = uln_config[0].clone() {
        if let Some(value) = tupled_uint[0].as_uint() {
            value.0
        } else {
            U256::from(0)
        }
    } else {
        U256::from(0)
    };

    debug!(
        "{:?} confirmations required by MessageLib at: {:?}",
        num_confirmations,
        contract.address()
    );

    Ok(num_confirmations)
}

/// Idempotent check to see if there's work to do for the DVN.
pub async fn get_verified(config: &DVNConfig, contract: &ContractInst, required_confirmations: U256) -> Result<bool> {
    // Call the `verified` function on the contract
    let uln302 = contract
        .function(
            "_verified",
            &[
                DynSolValue::Address(config.sendlib_uln301_addr()?),
                // HeaderHash
                // PayloadHash
                DynSolValue::Uint(required_confirmations, 32),
            ],
        )?
        .call()
        .await?;

    let uln302_state = if uln302[0].as_bool().unwrap() {
        debug!("Packet already verified, DVN workflow can stop.");
        true
    } else {
        debug!("Packet hasn't been verified. Call `verify`.");
        false
    };

    // Call the `_verified` function on the contract
    let uln301 = contract
        .function(
            "_verified",
            &[
                DynSolValue::Address(config.sendlib_uln301_addr()?),
                // HeaderHash
                // PayloadHash
                DynSolValue::Uint(required_confirmations, 32),
            ],
        )?
        .call()
        .await?;

    let uln301_state = if uln301[0].as_bool().unwrap() {
        debug!("Packet already verified, DVN workflow can stop.");
        true
    } else {
        debug!("Packet hasn't been verified. Call `verify`.");
        false
    };

    Ok(uln302_state && uln301_state)
}

pub async fn verify(
    config: &DVNConfig,
    contract: &ContractInst,
    packet_header: &[u8],
    payload: &[u8],
    confirmations: U256,
) -> Result<bool> {
    // Create the hash of the payload
    let payload_hash = keccak256(payload);

    // Call the `verified` function on the contract
    let _ = contract
        .function(
            "verify",
            &[
                // PacketHeader
                //packet_header,
                // PayloadHash
                //payload_hash,
                // Confirmations
                //DynSolValue::Uint(confirmations, 32),
            ],
        )?
        .call()
        .await?;

    Ok(false)
}

/// Helper for hashing some data with `keccak256`.
fn keccak256(data: &[u8]) -> [u8; 32] {
    let mut hasher = Keccak256::default();
    hasher.update(data);
    let result = hasher.finalize();
    result.into()
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::config;

    #[tokio::test]
    async fn test_get_confirmations() -> Result<()> {
        // Set up
        let config = config::DVNConfig::load_from_env()?;
        let http_provider = get_http_provider(&config)?;
        let sendlib_abi = get_sendlib_abi()?;
        let sendlib_contract = create_contract_instance(&config, http_provider, sendlib_abi)?;

        // Query contract value
        let required_confirmations = get_confirmations(&config, &sendlib_contract).await?;

        // Check the value is what we expect
        assert_eq!(required_confirmations, U256::from(20));

        Ok(())
    }
}
