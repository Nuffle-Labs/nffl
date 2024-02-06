use alloy_primitives::{hex, keccak256, Address, FixedBytes};
use alloy_sol_types::{eip712_domain, sol, Eip712Domain, SolStruct, SolValue};
use k256::{ecdsa::SigningKey, elliptic_curve::sec1::ToEncodedPoint, AffinePoint};
use near_workspaces::{types::NearToken, Account, AccountId, Contract};
use serde::{Deserialize, Serialize};
use serde_json::json;
use std::{env, fs};

sol! {
    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    struct StateRootUpdateMessage {
        uint32 rollupId;
        uint64 blockHeight;
        uint64 timestamp;
        bytes32 stateRoot;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    struct Operator {
        G1Point pubkey;
        uint128 weight;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    struct OperatorSetUpdateMessage {
        uint64 id;
        uint64 timestamp;
        Operator[] operators;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    struct CheckpointTaskResponseMessage {
        uint32 referenceTaskIndex;
        bytes32 stateRootUpdatesRoot;
        bytes32 operatorSetUpdatesRoot;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    struct EthNearAccountLink {
        address ethAddress;
        string nearAccountId;
    }
}

const DOMAIN: Eip712Domain = eip712_domain!(
    name: "SFFL",
    version: "0",
);

fn secret_key_to_address(secret_key: &SigningKey) -> Address {
    let public_key = secret_key.verifying_key();
    let affine: &AffinePoint = public_key.as_ref();
    let encoded = affine.to_encoded_point(false);

    Address::from_raw_public_key(&encoded.as_bytes()[1..])
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let wasm_arg: &str = &(env::args().nth(1).unwrap());
    let wasm_filepath = fs::canonicalize(env::current_dir()?.join(wasm_arg))?;

    let worker = near_workspaces::sandbox().await?;
    let wasm = std::fs::read(wasm_filepath)?;
    let contract = worker.dev_deploy(&wasm).await?;

    let account = worker.dev_create_account().await?;
    let alice = account
        .create_subaccount("alice")
        .initial_balance(NearToken::from_near(30))
        .transact()
        .await?
        .into_result()?;

    let _ = alice.call(contract.id(), "new").transact().await?.into_result()?;

    let _ = alice
        .call(contract.id(), "storage_deposit")
        .args_json(json!({ "account_id": None::<AccountId>, "registration_only": None::<bool> }))
        .deposit(NearToken::from_near(1))
        .transact()
        .await?
        .into_result()?;

    let test_private_key = hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318");

    test_operator_setup(&alice, &contract, &test_private_key).await?;
    test_message_submission(&alice, &contract).await?;

    Ok(())
}

async fn test_operator_setup(
    account: &Account,
    contract: &Contract,
    private_key: &[u8; 32],
) -> Result<(), Box<dyn std::error::Error>> {
    let signing_key: SigningKey = SigningKey::from_bytes(private_key.into())?;
    let eth_address = secret_key_to_address(&signing_key);

    let msg = EthNearAccountLink {
        ethAddress: eth_address,
        nearAccountId: account.id().to_string(),
    };

    let signing_hash = msg.eip712_signing_hash(&DOMAIN);
    let (signature, recid) = signing_key.sign_prehash_recoverable(signing_hash.as_slice()).unwrap();

    let _ = account
        .call(contract.id(), "init_operator")
        .args_json(json!({
            "msg": msg,
            "signature": FixedBytes::<64>::from_slice(signature.to_bytes().as_slice()),
            "v": u8::from(recid),
        }))
        .transact()
        .await?;

    let stored_eth_address = account
        .call(contract.id(), "get_eth_address")
        .args_json(json!({"account_id": account.id() }))
        .view()
        .await?
        .json::<Option<Address>>()?
        .expect("Stored Ethereum address is None");

    assert_eq!(
        eth_address, stored_eth_address,
        "The Ethereum address should match the initialized value"
    );
    println!("Passed ✅ operator setup");

    Ok(())
}

async fn test_message_submission(account: &Account, contract: &Contract) -> Result<(), Box<dyn std::error::Error>> {
    let msg = StateRootUpdateMessage {
        rollupId: 1,
        blockHeight: 2,
        timestamp: 3,
        stateRoot: FixedBytes::<32>::ZERO,
    };

    let mock_signature = FixedBytes::<64>::left_padding_from(&hex!("def1"));

    let _ = account
        .call(contract.id(), "post_state_root_update_signature")
        .args_json(json!({ "msg": msg, "signature": mock_signature.clone() }))
        .transact()
        .await?;

    let eth_address = account
        .call(contract.id(), "get_eth_address")
        .args_json(json!({"account_id": account.id() }))
        .view()
        .await?
        .json::<Option<Address>>()?
        .expect("Stored Ethereum address is None");

    let stored_messages = account
        .call(contract.id(), "get_state_root_updates")
        .args_json(json!({"rollup_id": msg.rollupId, "block_height": msg.blockHeight }))
        .view()
        .await?
        .json::<Option<Vec<StateRootUpdateMessage>>>()?
        .expect("Stored state root update messasges is None");

    let stored_signature = account
        .call(contract.id(), "get_message_signature")
        .args_json(json!({"msg_hash": keccak256(&msg.abi_encode()), "eth_address": &eth_address }))
        .view()
        .await?
        .json::<Option<FixedBytes<64>>>()?
        .expect("Stored signature is None");

    assert_eq!(
        stored_messages,
        vec![msg],
        "The stored messages vector should be exactly the message that was included"
    );
    assert_eq!(
        stored_signature, mock_signature,
        "The stored signature should be the one used in the message submission"
    );

    println!("Passed ✅ message submission");

    Ok(())
}
