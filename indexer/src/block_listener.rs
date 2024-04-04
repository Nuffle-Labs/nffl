use futures::future::join_all;
use near_indexer::near_primitives::{types::AccountId, views::ActionView};
use std::collections::HashMap;
use tokio::sync::mpsc;

use crate::errors::Result;

#[derive(Clone, Debug)]
pub(crate) struct CandidateData {
    pub rollup_id: u32,
    pub transaction: near_indexer::IndexerTransactionWithOutcome,
    pub payloads: Vec<Vec<u8>>,
}

pub(crate) struct BlockListener {
    stream: mpsc::Receiver<near_indexer::StreamerMessage>,
    receipt_sender: mpsc::Sender<CandidateData>,
    addresses_to_rollup_ids: HashMap<AccountId, u32>,
}

#[derive(Clone, Debug, serde::Serialize, serde::Deserialize)]
pub(crate) struct TransactionWithRollupId {
    pub(crate) rollup_id: u32,
    pub(crate) transaction: near_indexer::IndexerTransactionWithOutcome,
}

impl BlockListener {
    pub(crate) fn new(
        stream: mpsc::Receiver<near_indexer::StreamerMessage>,
        receipt_sender: mpsc::Sender<CandidateData>,
        addresses_to_rollup_ids: HashMap<AccountId, u32>,
    ) -> Self {
        Self {
            stream,
            receipt_sender,
            addresses_to_rollup_ids,
        }
    }

    fn transaction_filter_map(transaction: TransactionWithRollupId) -> Option<CandidateData> {
        let actions = &transaction.transaction.transaction.actions;
        let payloads: Vec<Vec<u8>> = actions.clone().into_iter().filter_map(Self::extract_args).collect();

        if payloads.is_empty() {
            return None;
        }

        Some(CandidateData {
            rollup_id: transaction.rollup_id,
            transaction: transaction.transaction,
            payloads,
        })
    }

    fn extract_args(action: ActionView) -> Option<Vec<u8>> {
        match action {
            ActionView::FunctionCall { method_name, args, .. } if method_name == "submit" => Some(args.into()),
            _ => None,
        }
    }

    async fn process_stream(self) -> Result<()> {
        let Self {
            mut stream,
            receipt_sender,
            addresses_to_rollup_ids,
        } = self;

        while let Some(streamer_message) = stream.recv().await {
            // TODO: check receipt_receiver is closed?
            let candidates_data: Vec<CandidateData> = streamer_message
                .shards
                .into_iter()
                .flat_map(|shard| shard.chunk)
                .flat_map(|chunk| {
                    chunk.transactions.into_iter().filter_map(|transaction| {
                        addresses_to_rollup_ids
                            .get(&transaction.transaction.receiver_id)
                            .map(|rollup_id| TransactionWithRollupId {
                                rollup_id: *rollup_id,
                                transaction,
                            })
                    })
                })
                .filter_map(Self::transaction_filter_map)
                .collect();

            if candidates_data.is_empty() {
                continue;
            }

            let results = join_all(candidates_data.into_iter().map(|receipt| receipt_sender.send(receipt))).await;
            results.into_iter().collect::<Result<_, _>>()?;
        }

        Ok(())
    }

    pub(crate) async fn start(self) -> Result<()> {
        let sender = self.receipt_sender.clone();
        tokio::select! {
            result = self.process_stream() => result,
            _ = sender.closed() => {
                Ok(())
            },
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::block_listener::{BlockListener, CandidateData, TransactionWithRollupId};
    use near_crypto::{KeyType, PublicKey, Signature};
    use near_indexer::near_primitives::hash::CryptoHash;
    use near_indexer::near_primitives::types::AccountId;
    use near_indexer::near_primitives::views::{
        ActionView, ExecutionMetadataView, ExecutionOutcomeView, ExecutionOutcomeWithIdView, ExecutionStatusView,
        SignedTransactionView,
    };
    use near_indexer::{IndexerExecutionOutcomeWithOptionalReceipt, IndexerTransactionWithOutcome, StreamerMessage};
    use std::collections::HashMap;
    use std::path::PathBuf;
    use std::str::FromStr;
    use std::time::Duration;
    use tokio::sync::mpsc;
    use tokio::sync::mpsc::error::TryRecvError;

    impl PartialEq for CandidateData {
        fn eq(&self, other: &Self) -> bool {
            let outcome_eq = (self.transaction.outcome.execution_outcome
                == other.transaction.outcome.execution_outcome)
                && (self.transaction.outcome.receipt == other.transaction.outcome.receipt);
            let transaction_eq = (self.transaction.transaction == other.transaction.transaction) && outcome_eq;
            self.payloads == other.payloads && transaction_eq
        }
    }

    fn get_default_execution_outcome() -> IndexerExecutionOutcomeWithOptionalReceipt {
        IndexerExecutionOutcomeWithOptionalReceipt {
            execution_outcome: ExecutionOutcomeWithIdView {
                proof: vec![],
                block_hash: CryptoHash::default(),
                id: CryptoHash::default(),
                outcome: ExecutionOutcomeView {
                    logs: vec![],
                    receipt_ids: vec![],
                    gas_burnt: 0,
                    tokens_burnt: 0,
                    executor_id: AccountId::from_str("test_signer").unwrap(),
                    status: ExecutionStatusView::SuccessValue(vec![]),
                    metadata: ExecutionMetadataView::default(),
                },
            },
            receipt: None,
        }
    }

    #[test]
    fn test_candidate_data_extraction() {
        let rollup_id = 0;
        let da_contract_id = AccountId::from_str("a.test").unwrap();

        let test_predecessor_id = AccountId::from_str("test_predecessor").unwrap();

        let default_outcome = get_default_execution_outcome();

        let valid_transaction = IndexerTransactionWithOutcome {
            transaction: SignedTransactionView {
                signer_id: test_predecessor_id.clone(),
                public_key: PublicKey::empty(KeyType::ED25519),
                nonce: 0,
                receiver_id: da_contract_id.clone(),
                actions: vec![ActionView::FunctionCall {
                    method_name: "submit".into(),
                    args: vec![1, 2, 3].into(),
                    gas: 100,
                    deposit: 100,
                }],
                signature: Signature::default(),
                hash: CryptoHash::default(),
            },
            outcome: default_outcome.clone(),
        };
        let valid_transaction = TransactionWithRollupId {
            rollup_id,
            transaction: valid_transaction,
        };

        let invalid_transaction = IndexerTransactionWithOutcome {
            transaction: SignedTransactionView {
                signer_id: test_predecessor_id.clone(),
                public_key: PublicKey::empty(KeyType::ED25519),
                nonce: 0,
                receiver_id: da_contract_id,
                actions: vec![ActionView::CreateAccount],
                signature: Signature::default(),
                hash: CryptoHash::default(),
            },
            outcome: default_outcome,
        };

        let invalid_transaction = TransactionWithRollupId {
            rollup_id,
            transaction: invalid_transaction,
        };

        // Test valid receipt
        assert_eq!(
            BlockListener::transaction_filter_map(valid_transaction.clone()),
            Some(CandidateData {
                rollup_id,
                transaction: valid_transaction.transaction,
                payloads: vec![vec![1, 2, 3]],
            })
        );

        // Test invalid action receipt
        assert_eq!(BlockListener::transaction_filter_map(invalid_transaction), None);
    }

    #[test]
    fn test_multiple_submit_actions() {
        let rollup_id = 10;
        let da_contract_id = AccountId::from_str("a.test").unwrap();
        let actions = vec![
            ActionView::FunctionCall {
                method_name: "submit".into(),
                args: vec![1, 2, 3].into(),
                gas: 100,
                deposit: 100,
            },
            ActionView::FunctionCall {
                method_name: "submit".into(),
                args: vec![4, 4, 4].into(),
                gas: 100,
                deposit: 100,
            },
            ActionView::FunctionCall {
                method_name: "random".into(),
                args: vec![1, 2, 3].into(),
                gas: 100,
                deposit: 100,
            },
            ActionView::DeleteAccount {
                beneficiary_id: da_contract_id.clone(),
            },
        ];

        let valid_transaction = IndexerTransactionWithOutcome {
            transaction: SignedTransactionView {
                signer_id: AccountId::from_str("test_signer").unwrap(),
                public_key: PublicKey::empty(KeyType::ED25519),
                nonce: 0,
                receiver_id: da_contract_id.clone(),
                actions,
                signature: Signature::default(),
                hash: CryptoHash::hash_bytes(b"test_tx_id"),
            },
            outcome: get_default_execution_outcome(),
        };
        let valid_transaction = TransactionWithRollupId {
            rollup_id,
            transaction: valid_transaction,
        };

        // Test valid receipt
        assert_eq!(
            BlockListener::transaction_filter_map(valid_transaction.clone()),
            Some(CandidateData {
                rollup_id,
                transaction: valid_transaction.transaction,
                payloads: vec![vec![1, 2, 3], vec![4, 4, 4]],
            })
        );
    }

    struct StreamerMessages {
        pub empty: Vec<StreamerMessage>,
        pub candidates: Vec<StreamerMessage>,
    }

    struct StreamerMessagesLoader;
    impl StreamerMessagesLoader {
        fn load() -> StreamerMessages {
            let test_data_dir = concat!(env!("CARGO_MANIFEST_DIR"), "/test_data");
            let empty_data_path = [test_data_dir, "/empty"].concat();
            let candidates_data_path = [test_data_dir, "/candidates"].concat();

            StreamerMessages {
                empty: Self::load_messages(&empty_data_path),
                candidates: Self::load_messages(&candidates_data_path),
            }
        }

        fn load_messages(dir_path: &str) -> Vec<StreamerMessage> {
            let files = std::fs::read_dir(dir_path)
                .map(|entry| {
                    entry
                        .map(|dir_entry| {
                            let dir_entry = dir_entry.unwrap();
                            dir_entry.path()
                        })
                        .collect::<Vec<PathBuf>>()
                })
                .unwrap();

            files
                .into_iter()
                .map(|file| {
                    let file = std::fs::File::open(file).unwrap();
                    serde_json::from_reader(file).unwrap()
                })
                .collect()
        }
    }

    #[tokio::test]
    async fn test_empty_listener() {
        let rollup_id = 1;
        let da_contract_id: AccountId = "da.test.near".parse().unwrap();

        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, mut receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, HashMap::from([(da_contract_id, rollup_id)]));
        let _ = tokio::spawn(listener.start());

        for el in streamer_messages.empty {
            sender.send(el).await.unwrap();
        }

        tokio::time::sleep(Duration::from_millis(10)).await;

        assert_eq!(receipt_receiver.try_recv(), Err(TryRecvError::Empty));
    }

    #[tokio::test]
    async fn test_candidates_listener() {
        let rollup_id = 1;
        let da_contract_id = "da.test.near".parse().unwrap();
        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, mut receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, HashMap::from([(da_contract_id, rollup_id)]));
        let handle = tokio::spawn(listener.start());

        let expected = streamer_messages.candidates.len();
        for el in streamer_messages.candidates {
            sender.send(el).await.unwrap();
        }

        drop(sender);
        handle.await.unwrap().unwrap();

        let mut counter = 0;
        while let Some(_) = receipt_receiver.recv().await {
            counter += 1;
        }

        assert_eq!(expected, counter);
    }

    #[tokio::test]
    async fn test_shutdown() {
        let rollup_id = 5;
        let da_contract_id = "da.test.near".parse().unwrap();
        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, HashMap::from([(da_contract_id, rollup_id)]));
        let handle = tokio::spawn(listener.start());
        for (i, el) in streamer_messages.empty.into_iter().enumerate() {
            sender.send(el).await.unwrap();

            // some random number
            if i == 5 {
                break;
            }
        }

        drop(receipt_receiver);
        // Sender::closed is trigerred
        assert!(handle.await.unwrap().is_ok());
    }
}
