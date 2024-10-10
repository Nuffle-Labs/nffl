//! Utilities for interacting with onchain contracts.

use crate::{
    chain::{ContractInst, HttpProvider},
    config::DVNConfig,
};
use alloy::{
    contract::{ContractInstance, Interface},
    dyn_abi::DynSolValue,
    network::Ethereum,
    primitives::{Address, Bytes, FixedBytes, U256},
    transports::http::{Client, Http},
};
use alloy_json_abi::JsonAbi;
use eyre::{eyre, OptionExt, Result};
use sha3::{Digest, Keccak256};

/// Create a contract instance from the ABI to interact with on-chain instance.
pub fn create_contract_instance(config: &DVNConfig, http_provider: HttpProvider, abi: JsonAbi) -> Result<ContractInst> {
    let contract: ContractInstance<Http<Client>, _, Ethereum> = ContractInstance::new(
        config.sendlib_uln302_addr()?,
        http_provider.clone(),
        Interface::new(abi),
    );
    Ok(contract)
}

/// Get the address of the MessageLib on the destination chain
pub async fn get_messagelib_addr(contract: &ContractInst, eid: U256) -> Result<Address> {
    // Call the `getUlnConfig` function on the contract
    let receive_library = contract
        .function(
            "getReceiveLibrary",
            &[DynSolValue::Address(*contract.address()), DynSolValue::Uint(eid, 32)],
        )?
        .call()
        .await?;

    match receive_library[0] {
        DynSolValue::Address(address) => Ok(address),
        _ => Err(eyre!("Failed to get address")),
    }
}

/// Get the number of required confirmations by the ULN.
pub async fn query_confirmations(contract: &ContractInst, eid: U256) -> Result<U256> {
    // Call the `getUlnConfig` function on the contract
    let uln_config = contract
        .function(
            "getUlnConfig",
            &[DynSolValue::Address(*contract.address()), DynSolValue::Uint(eid, 32)],
        )?
        .call()
        .await?;

    match &uln_config[0] {
        DynSolValue::Tuple(tupled_int) => {
            let value = tupled_int[0]
                .as_uint()
                .ok_or_eyre("Cannot parse response from MessageLib")?;
            Ok(value.0)
        }
        _ => Err(eyre::eyre!("Failed to get confirmations")),
    }
}

/// Idempotent check to see if there's work to do for the DVN.
pub async fn query_already_verified(
    contract: &ContractInst,
    dvn_address: Address,
    header_hash: &[u8],
    payload_hash: &[u8],
    required_confirmations: U256,
) -> Result<bool> {
    // Call the `_verified` function on the 302 contract, to check if the DVN has already verified
    // the packet.
    let receive_uln302 = contract
        .function(
            "_verified",
            &[
                DynSolValue::Address(dvn_address),             // DVN address
                DynSolValue::Bytes(header_hash.to_vec()),      // HeaderHash
                DynSolValue::Bytes(payload_hash.to_vec()),     // PayloadHash
                DynSolValue::Uint(required_confirmations, 32), // confirmations
            ],
        )?
        .call()
        .await?;

    let uln302_state = match receive_uln302[0] {
        DynSolValue::Bool(b) => Ok(b),
        _ => Err(eyre!("Failed to parse response from ULN302 for `_verified`")),
    }?;

    Ok(uln302_state)
}

pub async fn verify(
    contract: &ContractInst,
    packet_header: &Bytes,
    payload: &Bytes,
    confirmations: U256,
) -> Result<bool> {
    //// Create the hash of the payload
    let payload_hash = keccak256(payload);
    //
    // Call the `verified` function on the contract
    let _ = contract
        .function(
            "verify",
            &[
                DynSolValue::Bytes(packet_header.to_vec()),            // PacketHeader
                DynSolValue::FixedBytes(FixedBytes(payload_hash), 32), // PayloadHash
                DynSolValue::Uint(confirmations, 64),                  // Confirmations
            ],
        )?
        .call()
        .await?;

    Ok(false)
}

/// Hash some data with `keccak256`.
fn keccak256(data: &[u8]) -> [u8; 32] {
    let mut hasher = Keccak256::default();
    hasher.update(data);
    let result = hasher.finalize();
    result.into()
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::{
        chain::connections::{get_abi_from_path, get_http_provider},
        config,
    };

    #[tokio::test]
    async fn test_get_confirmations() -> Result<()> {
        // Set up
        let config = config::DVNConfig::load_from_env()?;
        let http_provider = get_http_provider(&config)?;
        let sendlib_abi = get_abi_from_path("./abi/ArbitrumSendLibUln302.json")?;
        let sendlib_contract = create_contract_instance(&config, http_provider, sendlib_abi)?;

        // Query contract value
        let required_confirmations = query_confirmations(&sendlib_contract, U256::from(30110)).await?;

        // Check the value is what we expect
        assert_eq!(required_confirmations, U256::from(20));

        Ok(())
    }

    //#[tokio::test]
    async fn test_get_verified() -> Result<()> {
        // NOTE: this method doesn't seem to exist in the contract
        // Set up
        let config = config::DVNConfig::load_from_env()?;
        let http_provider = get_http_provider(&config)?;
        let receivelib_abi = get_abi_from_path("./abi/ArbitrumReceiveLibUln302.json")?;
        let receivelib_contract = create_contract_instance(&config, http_provider, receivelib_abi)?;

        // Query contract value
        let uln302 = query_already_verified(
            &receivelib_contract,
            "0x7B9E184e07a6EE1aC23eAe0fe8D6Be2f663f05e6".parse::<Address>()?,
            &[1, 2, 3],
            &[1, 2, 3],
            U256::from(20),
        )
        .await?;

        // Check the value is what we expect
        assert!(uln302);

        Ok(())
    }
}
