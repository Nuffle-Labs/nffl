//! Utilities for interacting with onchain contracts.

use crate::chain::{ContractInst, HttpProvider};
use crate::data::packet_v1_codec::{guid, header, message, nonce, receiver, sender, src_eid};
use alloy::primitives::B256;
use alloy::{
    contract::{ContractInstance, Interface},
    dyn_abi::DynSolValue,
    json_abi::JsonAbi,
    network::Ethereum,
    primitives::{keccak256, Address, U256},
    transports::http::{Client, Http},
};
use eyre::{eyre, OptionExt, Result};
use tracing::{debug, error};

/// Create a contract instance from the ABI to interact with on-chain instance.
pub fn create_contract_instance(addr: Address, http_provider: HttpProvider, abi: JsonAbi) -> Result<ContractInst> {
    let contract: ContractInstance<Http<Client>, _, Ethereum> =
        ContractInstance::new(addr, http_provider, Interface::new(abi));
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

    match receive_library
        .first()
        .ok_or_eyre("ReceiveLibrary not found in contract")?
    {
        DynSolValue::Address(address) if address.len() == 20 => Ok(*address),
        _ => {
            error!("Failed to get a valid address");
            Err(eyre!("Failed to get a valid address"))
        }
    }
}

/// Get the number of required confirmations by the ULN.
///
/// The value returned is a solidity `UlnConfig[]` with, at least, one value.
///
/// See: https://github.com/LayerZero-Labs/LayerZero-v2/blob/main/packages/layerzero-v2/evm/messagelib/contracts/uln/UlnBase.sol
/// The struct `UlnConfig` is defined as follows:
///
/// ```solidity
/// struct UlnConfig {
///     uint64 confirmations;
///     uint8 requiredDVNCount;
///     uint8 optionalDVNCount;
///     uint8 optionalDVNThreshold;
///     address[] requiredDVNs;
///     address[] optionalDVNs;
/// }
/// ```
///
/// and we require only the first value.
pub async fn query_confirmations(contract: &ContractInst, eid: U256) -> Result<U256> {
    // Call the `getUlnConfig` function on the contract
    let uln_config = contract
        .function(
            "getUlnConfig",
            &[DynSolValue::Address(*contract.address()), DynSolValue::Uint(eid, 32)],
        )?
        .call()
        .await?;

    match &uln_config.first().ok_or_eyre("ULN config not found in contract")? {
        DynSolValue::Tuple(tupled_int) => {
            let value = tupled_int[0]
                .as_uint()
                .ok_or_eyre("Cannot parse response from MessageLib")?;
            Ok(value.0)
        }
        _ => {
            error!("Failed to get confirmations");
            Err(eyre!("Failed to get confirmations"))
        }
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
    debug!("Calling _verified on contract's ReceiveLib");

    let contract_state = contract
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

    let packet_state = match contract_state
        .first()
        .ok_or(eyre!("Empty response when querying `_verified`"))?
    {
        DynSolValue::Bool(b) => Ok(b),
        _ => {
            error!("Failed to parse response from ReceiveLib for `_verified`");
            Err(eyre!("Failed to parse response from ReceiveLib for `_verified`"))
        }
    }?;

    Ok(*packet_state)
}

pub async fn verify(contract: &ContractInst, packet_header: &[u8], payload: &[u8], confirmations: U256) -> Result<()> {
    //// Create the hash of the payload
    let payload_hash = keccak256(payload);

    // Call the `verified` function on the contract
    let _ = contract
        .function(
            "verify",
            &[
                DynSolValue::Bytes(packet_header.to_vec()), // PacketHeader
                DynSolValue::FixedBytes(payload_hash, 32),  // PayloadHash
                DynSolValue::Uint(confirmations, 64),       // Confirmations
            ],
        )?
        .call()
        .await?;

    Ok(())
}

/// If the state is `Executable`, your `Executor` should decode the packet's options
/// using the options.ts package and call the Endpoint's `lzReceive` function with
/// the packet information:
/// `endpoint.lzReceive(_origin, _receiver, _guid, _message, _extraData)`
pub async fn lz_receive(contract: &ContractInst, packet: &[u8]) -> Result<()> {
    let guid = guid(packet);
    let call_builder = contract.function(
        "_lzReceive",
        &[
            prepare_header(header(packet)),
            DynSolValue::Address(Address::from_slice(receiver(packet).as_slice())),
            DynSolValue::FixedBytes(B256::from_slice(guid.as_slice()), 32),
            DynSolValue::Bytes(message(packet).to_vec()),
            DynSolValue::Bytes(vec![]),
        ],
    )?;

    match call_builder.call().await {
        Ok(_) => debug!("Successfully called lzReceive for packet {:?}", guid),
        Err(e) => error!("Failed to call lzReceive for packet {:?}: {:?}", guid, e),
    }
    Ok(())
}

/// Converts `Origin` data structure from the received `PacketVerified`
/// to the `DynSolValue`, understandable by `alloy-rs`.
pub(crate) fn prepare_header(packet: &[u8]) -> DynSolValue {
    DynSolValue::CustomStruct {
        name: String::from("Origin"),
        prop_names: vec![String::from("srcEid"), String::from("sender"), String::from("nonce")],
        tuple: vec![
            DynSolValue::Uint(U256::from(src_eid(packet)), 32),
            DynSolValue::Bytes(sender(packet).to_vec()),
            DynSolValue::Uint(U256::from(nonce(packet)), 64),
        ],
    }
}
