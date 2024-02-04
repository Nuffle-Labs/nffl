use std::str::FromStr;

use near_sdk::borsh::{self, BorshDeserialize, BorshSerialize};
use near_sdk::json_types::U128;
use near_sdk::serde::{Deserialize, Serialize};
use near_sdk::store::{LookupMap, Vector};
use near_sdk::{env, near_bindgen, AccountId, PanicOnDefault};

use alloy_primitives::{Address, FixedBytes};
use alloy_sol_types::{eip712_domain, sol, Eip712Domain, SolStruct, SolValue};
use near_sdk_contract_tools::{
    hook::Hook, standard::nep145::hooks::PredecessorStorageAccountingHook, standard::nep145::*,
    Nep145,
};

sol! {
    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    #[serde(crate = "near_sdk::serde")]
    struct StateRootUpdateMessage {
        uint32 rollupId;
        uint64 blockHeight;
        uint64 nearBlockHeight;
        bytes32 stateRoot;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    #[serde(crate = "near_sdk::serde")]
    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    #[serde(crate = "near_sdk::serde")]
    struct Operator {
        G1Point pubkey;
        uint128 weight;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    #[serde(crate = "near_sdk::serde")]
    struct OperatorSetUpdateMessage {
        uint64 id;
        uint64 nearBlockHeight;
        Operator[] operators;
    }

    #[derive(Serialize, Deserialize, PartialEq, Debug)]
    #[serde(crate = "near_sdk::serde")]
    struct CheckpointTaskResponseMessage {
        uint32 referenceTaskIndex;
        bytes32 stateRootUpdatesRoot;
        bytes32 operatorSetUpdatesRoot;
    }

    #[derive(Serialize, Deserialize)]
    #[serde(crate = "near_sdk::serde")]
    struct EthNearAccountLink {
        address ethAddress;
        string nearAccountId;
    }
}

const DOMAIN: Eip712Domain = eip712_domain!(
    name: "SFFL",
    version: "0",
);

macro_rules! sol_type_borsh {
    ($name: ident) => {
        impl BorshSerialize for $name {
            #[inline]
            fn serialize<W: std::io::Write>(&self, writer: &mut W) -> Result<(), std::io::Error> {
                BorshSerialize::serialize(&$name::abi_encode(self), writer)?;
                Ok(())
            }
        }

        impl BorshDeserialize for $name {
            #[inline]
            fn deserialize(buf: &mut &[u8]) -> Result<Self, std::io::Error> {
                let encoded: Vec<u8> = BorshDeserialize::deserialize(buf)?;
                let value = $name::abi_decode(encoded.as_slice(), true)
                    .map_err(|err| std::io::Error::new(std::io::ErrorKind::InvalidInput, err))?;
                Ok(value)
            }
        }
    };
}

sol_type_borsh!(StateRootUpdateMessage);
sol_type_borsh!(OperatorSetUpdateMessage);
sol_type_borsh!(CheckpointTaskResponseMessage);
sol_type_borsh!(EthNearAccountLink);

#[derive(BorshDeserialize, BorshSerialize, PanicOnDefault, Nep145)]
#[nep145]
#[near_bindgen]
pub struct SFFLAgreementRegistry {
    operator_eth_address: LookupMap<AccountId, [u8; 20]>,
    state_root_updates: LookupMap<(u32, u64), Vector<StateRootUpdateMessage>>,
    operator_set_updates: LookupMap<u64, Vector<OperatorSetUpdateMessage>>,
    checkpoint_task_responses: LookupMap<u32, Vector<CheckpointTaskResponseMessage>>,
    message_signatures: LookupMap<[u8; 32], LookupMap<[u8; 20], [u8; 64]>>,
}

#[near_bindgen]
impl SFFLAgreementRegistry {
    #[init]
    pub fn new() -> Self {
        let mut _self = Self {
            operator_eth_address: LookupMap::new(b"operator_eth_address".to_vec()),
            state_root_updates: LookupMap::new(b"state_root_updates".to_vec()),
            operator_set_updates: LookupMap::new(b"operator_set_updates".to_vec()),
            checkpoint_task_responses: LookupMap::new(b"checkpoint_task_responses".to_vec()),
            message_signatures: LookupMap::new(b"message_signatures".to_vec()),
        };

        Nep145Controller::set_storage_balance_bounds(
            &mut _self,
            &StorageBalanceBounds {
                min: U128(0),
                max: None,
            },
        );

        _self
    }

    #[payable]
    pub fn init_operator(&mut self, msg: &EthNearAccountLink, signature: &FixedBytes<64>, v: u8) {
        PredecessorStorageAccountingHook::hook(self, &(), |contract| {
            let signing_hash = msg.eip712_signing_hash(&DOMAIN);
            let address = contract.ecrecover(&signing_hash, signature, v, true);

            if address != msg.ethAddress {
                std::panic!(
                    "Wrong message address: expected {}, found {}",
                    msg.ethAddress,
                    address
                );
            }

            let account_id = AccountId::from_str(&msg.nearAccountId)
                .unwrap_or_else(|_| std::panic!("Invalid account ID"));

            contract.operator_eth_address.insert(account_id, address.0 .0);
        })
    }

    #[payable]
    pub fn post_state_root_update_signature(
        &mut self,
        msg: StateRootUpdateMessage,
        signature: &FixedBytes<64>,
    ) {
        PredecessorStorageAccountingHook::hook(self, &(), |contract| {
            let eth_address = contract.caller_eth_address();
            let msg_hash = env::keccak256_array(&msg.abi_encode());

            if !contract.push_bls_signature(&msg_hash, &eth_address, &signature.0) {
                let vec = contract.get_and_init_state_root_updates(msg.rollupId, msg.blockHeight);
                vec.push(msg);
            }
        })
    }

    #[private]
    fn get_and_init_state_root_updates(
        &mut self,
        rollup_id: u32,
        block_height: u64,
    ) -> &mut Vector<StateRootUpdateMessage> {
        self.state_root_updates
            .entry((rollup_id, block_height))
            .or_insert_with(|| {
                let prefix: Vec<u8> = [
                    b"state_root_updates_vec".as_slice(),
                    rollup_id.to_be_bytes().as_slice(),
                    block_height.to_be_bytes().as_slice(),
                ]
                .concat();

                Vector::new(prefix)
            })
    }

    pub fn post_operator_set_update_signature(
        &mut self,
        msg: OperatorSetUpdateMessage,
        signature: &FixedBytes<64>,
    ) {
        PredecessorStorageAccountingHook::hook(self, &(), |contract| {
            let eth_address = contract.caller_eth_address();
            let msg_hash = env::keccak256_array(&msg.abi_encode());

            if !contract.push_bls_signature(&msg_hash, &eth_address, &signature.0) {
                let vec = contract.get_and_init_operator_set_updates(msg.id);
                vec.push(msg);
            }
        })
    }

    #[private]
    fn get_and_init_operator_set_updates(
        &mut self,
        id: u64,
    ) -> &mut Vector<OperatorSetUpdateMessage> {
        self.operator_set_updates.entry(id).or_insert_with(|| {
            let prefix: Vec<u8> = [
                b"operator_set_updates_vec".as_slice(),
                id.to_be_bytes().as_slice(),
            ]
            .concat();

            Vector::new(prefix)
        })
    }

    pub fn post_checkpoint_signature(
        &mut self,
        msg: CheckpointTaskResponseMessage,
        signature: &FixedBytes<64>,
    ) {
        PredecessorStorageAccountingHook::hook(self, &(), |contract| {
            let eth_address = contract.caller_eth_address();
            let msg_hash = env::keccak256_array(&msg.abi_encode());

            if !contract.push_bls_signature(&msg_hash, &eth_address, &signature.0) {
                let vec = contract.get_and_init_checkpoint_task_responses(msg.referenceTaskIndex);
                vec.push(msg);
            }
        })
    }

    #[private]
    fn get_and_init_checkpoint_task_responses(
        &mut self,
        reference_task_index: u32,
    ) -> &mut Vector<CheckpointTaskResponseMessage> {
        self.checkpoint_task_responses
            .entry(reference_task_index)
            .or_insert_with(|| {
                let prefix: Vec<u8> = [
                    b"checkpoint_task_responses_vec".as_slice(),
                    reference_task_index.to_be_bytes().as_slice(),
                ]
                .concat();

                Vector::new(prefix)
            })
    }

    pub fn get_eth_address(&self, account_id: &AccountId) -> Option<Address> {
        self.operator_eth_address
            .get(&account_id)
            .map(Address::from)
    }

    pub fn get_state_root_updates(
        &self,
        rollup_id: u32,
        block_height: u64,
    ) -> Option<Vec<&StateRootUpdateMessage>> {
        self.state_root_updates
            .get(&(rollup_id, block_height))
            .and_then(|vector| Some(vector.iter().collect()))
    }

    pub fn get_operator_set_updates(&self, id: u64) -> Option<Vec<&OperatorSetUpdateMessage>> {
        self.operator_set_updates
            .get(&id)
            .and_then(|vector| Some(vector.iter().collect()))
    }

    pub fn get_checkpoint_task_responses(
        &self,
        task_id: u32,
    ) -> Option<Vec<&CheckpointTaskResponseMessage>> {
        self.checkpoint_task_responses
            .get(&task_id)
            .and_then(|vector| Some(vector.iter().collect()))
    }

    pub fn get_message_signature(
        &self,
        msg_hash: &FixedBytes<32>,
        eth_address: &Address,
    ) -> Option<FixedBytes<64>> {
        self.message_signatures
            .get(&msg_hash.0)
            .and_then(|addr_to_sig| addr_to_sig.get(&eth_address.0 .0))
            .map(FixedBytes::from)
    }

    #[private]
    fn caller_eth_address(&self) -> [u8; 20] {
        *self
            .operator_eth_address
            .get(&env::predecessor_account_id())
            .unwrap_or_else(|| std::panic!("No known ethereum address"))
    }

    #[private]
    fn push_bls_signature(
        &mut self,
        msg_hash: &[u8; 32],
        eth_address: &[u8; 20],
        signature: &[u8; 64],
    ) -> bool {
        let mut had_key = true;

        let map = self.message_signatures.entry(*msg_hash).or_insert_with(|| {
            had_key = false;

            let prefix: Vec<u8> =
                [b"message_signatures_map".as_slice(), msg_hash.as_slice()].concat();

            LookupMap::new(prefix)
        });

        map.insert(*eth_address, signature.clone());

        had_key
    }

    #[private]
    fn ecrecover(
        &self,
        msg_hash: &FixedBytes<32>,
        signature: &FixedBytes<64>,
        v: u8,
        malleability_flag: bool,
    ) -> Address {
        match env::ecrecover(
            msg_hash.as_slice(),
            signature.as_slice(),
            v,
            malleability_flag,
        ) {
            Some(pubkey) => {
                let hash = env::keccak256_array(&pubkey);

                Address::from_slice(&hash[12..32])
            }
            None => std::panic!("Invalid signature"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use alloy_primitives::{hex, U256};
    use k256::{ecdsa::SigningKey, elliptic_curve::sec1::ToEncodedPoint, AffinePoint};
    use near_sdk::test_utils::VMContextBuilder;
    use near_sdk::{testing_env, VMContext, ONE_NEAR};

    fn secret_key_to_address(secret_key: &SigningKey) -> Address {
        let public_key = secret_key.verifying_key();
        let affine: &AffinePoint = public_key.as_ref();
        let encoded = affine.to_encoded_point(false);

        Address::from_raw_public_key(&encoded.as_bytes()[1..])
    }

    fn get_context(account_id: AccountId) -> VMContext {
        VMContextBuilder::new()
            .predecessor_account_id(account_id)
            .attached_deposit(ONE_NEAR)
            .build()
    }

    fn _storage_deposit(contract: &mut SFFLAgreementRegistry) {
        contract.storage_deposit(None, None);
    }

    fn _init_operator(
        contract: &mut SFFLAgreementRegistry,
        account_id: AccountId,
        private_key: &[u8; 32],
        msg: Option<&EthNearAccountLink>,
    ) -> Address {
        let signing_key = SigningKey::from_bytes(private_key.into()).unwrap();
        let eth_address = secret_key_to_address(&signing_key);

        let init_msg = EthNearAccountLink {
            ethAddress: eth_address,
            nearAccountId: account_id.to_string(),
        };

        let msg = msg.unwrap_or(&init_msg);

        let signing_hash = msg.eip712_signing_hash(&DOMAIN);
        let (signature, recid) = signing_key
            .sign_prehash_recoverable(signing_hash.as_slice())
            .unwrap();

        _storage_deposit(contract);
        contract.init_operator(
            &msg.clone(),
            &FixedBytes::<64>::from_slice(signature.to_bytes().as_slice()),
            u8::from(recid),
        );

        msg.ethAddress
    }

    #[test]
    fn init_operator() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        assert_eq!(
            contract.get_eth_address(&"alice.near".parse().unwrap()),
            None
        );

        let eth_address = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        assert_eq!(
            contract.get_eth_address(&"alice.near".parse().unwrap()),
            Some(eth_address)
        );
    }

    #[test]
    #[should_panic(
        expected = "Wrong message address: expected 0x0000000000000000000000000000000000000000, found 0x2c7536E3605D9C16a7a3D7b1898e529396a65c23"
    )]
    fn init_operator_wrong_address() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let msg = EthNearAccountLink {
            ethAddress: Address::ZERO,
            nearAccountId: "alice.near".to_string(),
        };

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            Some(&msg),
        );
    }

    #[test]
    #[should_panic(expected = "Invalid account ID")]
    fn init_operator_invalid_account_id() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let signing_key = SigningKey::from_bytes(
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318").into(),
        )
        .unwrap();
        let eth_address = secret_key_to_address(&signing_key);

        let msg = EthNearAccountLink {
            ethAddress: eth_address,
            nearAccountId: String::from_str("").unwrap(),
        };

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            Some(&msg),
        );
    }

    #[test]
    #[should_panic(expected = "No known ethereum address")]
    fn post_state_root_update_signature_no_eth_address() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let msg = StateRootUpdateMessage {
            rollupId: 1,
            blockHeight: 2,
            nearBlockHeight: 3,
            stateRoot: FixedBytes::<32>::ZERO,
        };

        _storage_deposit(&mut contract);
        contract.post_state_root_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);
    }

    #[test]
    fn post_state_root_update_signature_unexisting_message() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let eth_address = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = StateRootUpdateMessage {
            rollupId: 1,
            blockHeight: 2,
            nearBlockHeight: 3,
            stateRoot: FixedBytes::<32>::ZERO,
        };

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_state_root_updates(msg.rollupId, msg.blockHeight),
            None
        );

        contract.post_state_root_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            Some(FixedBytes::<64>::ZERO)
        );
        assert_eq!(
            contract.get_state_root_updates(msg.rollupId, msg.blockHeight),
            Some(vec![&msg])
        );
    }

    #[test]
    fn post_state_root_update_signature_existing_message() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = StateRootUpdateMessage {
            rollupId: 1,
            blockHeight: 2,
            nearBlockHeight: 3,
            stateRoot: FixedBytes::<32>::ZERO,
        };

        contract.post_state_root_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        let context = get_context("bob.near".parse().unwrap());
        testing_env!(context);

        let eth_address = _init_operator(
            &mut contract,
            "bob.near".parse().unwrap(),
            &hex!("24341428553285e10e74a5f26f4638ac53afb28c032aff1a04900e6eb115a404"),
            None,
        );

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_state_root_updates(msg.rollupId, msg.blockHeight),
            Some(vec![&msg])
        );

        contract.post_state_root_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            Some(FixedBytes::<64>::ZERO)
        );
        assert_eq!(
            contract.get_state_root_updates(msg.rollupId, msg.blockHeight),
            Some(vec![&msg])
        );
    }

    #[test]
    fn post_state_root_update_signature_existing_header() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = StateRootUpdateMessage {
            rollupId: 1,
            blockHeight: 2,
            nearBlockHeight: 3,
            stateRoot: FixedBytes::<32>::ZERO,
        };

        contract.post_state_root_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        let context = get_context("bob.near".parse().unwrap());
        testing_env!(context);

        let eth_address = _init_operator(
            &mut contract,
            "bob.near".parse().unwrap(),
            &hex!("24341428553285e10e74a5f26f4638ac53afb28c032aff1a04900e6eb115a404"),
            None,
        );

        let msg2 = StateRootUpdateMessage {
            rollupId: 1,
            blockHeight: 2,
            nearBlockHeight: 4,
            stateRoot: FixedBytes::<32>::left_padding_from(&hex!("f00d")),
        };

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg2.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_state_root_updates(msg.rollupId, msg.blockHeight),
            Some(vec![&msg])
        );

        let mock_sig = FixedBytes::<64>::left_padding_from(&hex!("abcd"));

        contract.post_state_root_update_signature(msg2.clone(), &mock_sig);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg2.abi_encode())),
                &eth_address
            ),
            Some(mock_sig)
        );
        assert_eq!(
            contract.get_state_root_updates(msg.rollupId, msg.blockHeight),
            Some(vec![&msg.clone(), &msg2])
        );
    }

    #[test]
    #[should_panic(expected = "No known ethereum address")]
    fn post_operator_set_update_signature_no_eth_address() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let msg = OperatorSetUpdateMessage {
            id: 1,
            nearBlockHeight: 2,
            operators: vec![],
        };

        _storage_deposit(&mut contract);
        contract.post_operator_set_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);
    }

    #[test]
    fn post_operator_set_update_signature_unexisting_message() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let eth_address = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = OperatorSetUpdateMessage {
            id: 1,
            nearBlockHeight: 2,
            operators: vec![],
        };

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(contract.get_operator_set_updates(msg.id), None);

        contract.post_operator_set_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            Some(FixedBytes::<64>::ZERO)
        );
        assert_eq!(contract.get_operator_set_updates(msg.id), Some(vec![&msg]));
    }

    #[test]
    fn post_operator_set_update_signature_existing_message() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = OperatorSetUpdateMessage {
            id: 1,
            nearBlockHeight: 2,
            operators: vec![],
        };

        contract.post_operator_set_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        let context = get_context("bob.near".parse().unwrap());
        testing_env!(context);

        let eth_address = _init_operator(
            &mut contract,
            "bob.near".parse().unwrap(),
            &hex!("24341428553285e10e74a5f26f4638ac53afb28c032aff1a04900e6eb115a404"),
            None,
        );

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(contract.get_operator_set_updates(msg.id), Some(vec![&msg]));

        contract.post_operator_set_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            Some(FixedBytes::<64>::ZERO)
        );
        assert_eq!(contract.get_operator_set_updates(msg.id), Some(vec![&msg]));
    }

    #[test]
    fn post_operator_set_update_signature_existing_header() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = OperatorSetUpdateMessage {
            id: 1,
            nearBlockHeight: 2,
            operators: vec![],
        };

        contract.post_operator_set_update_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        let context = get_context("bob.near".parse().unwrap());
        testing_env!(context);

        let eth_address = _init_operator(
            &mut contract,
            "bob.near".parse().unwrap(),
            &hex!("24341428553285e10e74a5f26f4638ac53afb28c032aff1a04900e6eb115a404"),
            None,
        );

        let msg2 = OperatorSetUpdateMessage {
            id: 1,
            nearBlockHeight: 3,
            operators: vec![Operator {
                pubkey: G1Point {
                    X: U256::ZERO,
                    Y: U256::ZERO,
                },
                weight: 0,
            }],
        };

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg2.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(contract.get_operator_set_updates(msg.id), Some(vec![&msg]));

        let mock_sig = FixedBytes::<64>::left_padding_from(&hex!("abcd"));

        contract.post_operator_set_update_signature(msg2.clone(), &mock_sig);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg2.abi_encode())),
                &eth_address
            ),
            Some(mock_sig)
        );
        assert_eq!(
            contract.get_operator_set_updates(msg.id),
            Some(vec![&msg.clone(), &msg2])
        );
    }

    #[test]
    #[should_panic(expected = "No known ethereum address")]
    fn post_checkpoint_signature_no_eth_address() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let msg = CheckpointTaskResponseMessage {
            referenceTaskIndex: 1,
            stateRootUpdatesRoot: FixedBytes::<32>::ZERO,
            operatorSetUpdatesRoot: FixedBytes::<32>::ZERO,
        };

        _storage_deposit(&mut contract);
        contract.post_checkpoint_signature(msg.clone(), &FixedBytes::<64>::ZERO);
    }

    #[test]
    fn post_checkpoint_signature_unexisting_message() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let eth_address = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = CheckpointTaskResponseMessage {
            referenceTaskIndex: 1,
            stateRootUpdatesRoot: FixedBytes::<32>::ZERO,
            operatorSetUpdatesRoot: FixedBytes::<32>::ZERO,
        };

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_checkpoint_task_responses(msg.referenceTaskIndex),
            None
        );

        contract.post_checkpoint_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            Some(FixedBytes::<64>::ZERO)
        );
        assert_eq!(
            contract.get_checkpoint_task_responses(msg.referenceTaskIndex),
            Some(vec![&msg])
        );
    }

    #[test]
    fn post_checkpoint_signature_existing_message() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = CheckpointTaskResponseMessage {
            referenceTaskIndex: 1,
            stateRootUpdatesRoot: FixedBytes::<32>::ZERO,
            operatorSetUpdatesRoot: FixedBytes::<32>::ZERO,
        };

        contract.post_checkpoint_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        let context = get_context("bob.near".parse().unwrap());
        testing_env!(context);

        let eth_address = _init_operator(
            &mut contract,
            "bob.near".parse().unwrap(),
            &hex!("24341428553285e10e74a5f26f4638ac53afb28c032aff1a04900e6eb115a404"),
            None,
        );

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_checkpoint_task_responses(msg.referenceTaskIndex),
            Some(vec![&msg])
        );

        contract.post_checkpoint_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            Some(FixedBytes::<64>::ZERO)
        );
        assert_eq!(
            contract.get_checkpoint_task_responses(msg.referenceTaskIndex),
            Some(vec![&msg])
        );
    }

    #[test]
    fn post_checkpoint_signature_existing_header() {
        let context = get_context("alice.near".parse().unwrap());
        testing_env!(context);

        let mut contract = SFFLAgreementRegistry::new();

        let _ = _init_operator(
            &mut contract,
            "alice.near".parse().unwrap(),
            &hex!("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"),
            None,
        );

        let msg = CheckpointTaskResponseMessage {
            referenceTaskIndex: 1,
            stateRootUpdatesRoot: FixedBytes::<32>::ZERO,
            operatorSetUpdatesRoot: FixedBytes::<32>::ZERO,
        };

        contract.post_checkpoint_signature(msg.clone(), &FixedBytes::<64>::ZERO);

        let context = get_context("bob.near".parse().unwrap());
        testing_env!(context);

        let eth_address = _init_operator(
            &mut contract,
            "bob.near".parse().unwrap(),
            &hex!("24341428553285e10e74a5f26f4638ac53afb28c032aff1a04900e6eb115a404"),
            None,
        );

        let msg2 = CheckpointTaskResponseMessage {
            referenceTaskIndex: 1,
            stateRootUpdatesRoot: FixedBytes::<32>::left_padding_from(&hex!("def1")),
            operatorSetUpdatesRoot: FixedBytes::<32>::left_padding_from(&hex!("f00d")),
        };

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg2.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_checkpoint_task_responses(msg.referenceTaskIndex),
            Some(vec![&msg])
        );

        let mock_sig = FixedBytes::<64>::left_padding_from(&hex!("abcd"));

        contract.post_checkpoint_signature(msg2.clone(), &mock_sig);

        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg.abi_encode())),
                &eth_address
            ),
            None
        );
        assert_eq!(
            contract.get_message_signature(
                &FixedBytes::<32>::from(env::keccak256_array(&msg2.abi_encode())),
                &eth_address
            ),
            Some(mock_sig)
        );
        assert_eq!(
            contract.get_checkpoint_task_responses(msg.referenceTaskIndex),
            Some(vec![&msg.clone(), &msg2])
        );
    }
}
