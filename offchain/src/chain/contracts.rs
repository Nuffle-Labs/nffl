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
/// The function is defined as:
/// ```solidity
/// function getUlnConfig(address _oapp, uint32 _remoteEid) public view returns (UlnConfig memory rtnConfig);
/// ```
///
/// The value returned is a solidity `UlnConfig[]` with, at least, one value.
/// See: https://github.com/LayerZero-Labs/LayerZero-v2/blob/main/packages/layerzero-v2/evm/messagelib/contracts/uln/UlnBase.sol
///
/// The struct `UlnConfig` is defined as follows:
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
                .ok_or_eyre("Cannot parse response as `uint` from MessageLib")?;
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
) -> bool {
    // Call the `_verified` function on the 302 contract, to check if the DVN has already verified
    // the packet.
    debug!("Calling _verified on contract's ReceiveLib");

    let call_builder = contract
        .function(
            "_verified",
            &[
                DynSolValue::Address(dvn_address),             // DVN address
                DynSolValue::Bytes(header_hash.to_vec()),      // HeaderHash
                DynSolValue::Bytes(payload_hash.to_vec()),     // PayloadHash
                DynSolValue::Uint(required_confirmations, 32), // confirmations
            ],
        );

    let Ok(call_builder) = call_builder else {
        error!("Failed to construct `_verified` caller");
        return false
    };

    let Ok(state) = call_builder.call().await else {
        error!("Failed to call `_verified` on contract");
        return false
    };

    match state.first() {
        Some(DynSolValue::Bool(b)) => *b,
        _ => {
            error!("Failed to parse as bool the `_verified` response from ReceiveLib");
            false
        }
    }
}

pub async fn verify(contract: &ContractInst, packet_header: &[u8], payload: &[u8], confirmations: U256) {
    //// Create the hash of the payload
    let payload_hash = keccak256(payload);

    // Call the `verified` function on the contract
    let call_builder = contract
        .function(
            "verify",
            &[
                DynSolValue::Bytes(packet_header.to_vec()), // PacketHeader
                DynSolValue::FixedBytes(payload_hash, 32),  // PayloadHash
                DynSolValue::Uint(confirmations, 64),       // Confirmations
            ],
        );

    if let Ok(call_builder) = call_builder {
        match call_builder.call().await {
            Err(e) => error!("Failed to call `verify`. Error: {:?}", e),
            _ => {},
        }
    } else {
        error!("Failed to construct `verify` caller");
    };
}

/// If the state is `Executable`, your `Executor` should decode the packet's options
/// using the options.ts package and call the Endpoint's `lzReceive` function with
/// the packet information:
/// `endpoint.lzReceive(_origin, _receiver, _guid, _message, _extraData)`
pub async fn lz_receive(contract: &ContractInst, packet: &[u8]) -> Result<()> {
    let guid = guid(packet);
    let call_builder_result = contract.function(
        "lzReceive",
        &[
            prepare_header(header(packet)),
            DynSolValue::Address(Address::from_slice(&receiver(packet)[0..20])),
            DynSolValue::FixedBytes(B256::from_slice(guid.as_slice()), 32),
            DynSolValue::Bytes(message(packet).to_vec()),
            DynSolValue::Bytes(vec![]),
        ],
    );

    if call_builder_result.is_err() {
        error!("Failed to call lzReceive, because it doesn't exist in the contract/ABI.");
        return Ok(());
    }

    call_builder_result.unwrap().call().await.map_err(|e| {
        error!("Failed to call lzReceive for packet {:?}: {:?}", guid, e);
        eyre!("lzReceive call failed: {}", e)
    })?;
    debug!("Successfully called lzReceive for packet {:?}", guid);
    Ok(())
}

/// Converts `Origin` data structure from the received `PacketVerified`
/// to the `DynSolValue`, understandable by `alloy-rs`.
pub(crate) fn prepare_header(packet: &[u8]) -> DynSolValue {
    const ORIGIN_STRUCT_NAME: &str = "Origin";
    const ORIGIN_PROPS: [&str; 3] = ["srcEid", "sender", "nonce"];

    DynSolValue::CustomStruct {
        name: String::from(ORIGIN_STRUCT_NAME),
        prop_names: ORIGIN_PROPS.iter().map(|&s| String::from(s)).collect(),
        tuple: vec![
            DynSolValue::Uint(U256::from(src_eid(packet)), 32),
            DynSolValue::FixedBytes(B256::from_slice(sender(packet).as_ref()), 32),
            DynSolValue::Uint(U256::from(nonce(packet)), 64),
        ],
    }
}
