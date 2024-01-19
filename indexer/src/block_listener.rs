use futures::future::join_all;
use near_indexer::near_primitives::views::ReceiptView;
use near_indexer::near_primitives::{
    types::{AccountId, TransactionOrReceiptId},
    views::{ActionView, ReceiptEnumView},
};
use tokio::sync::mpsc;

use crate::errors::Result;

#[derive(Clone, Debug)]
pub(crate) struct CandidateData {
    pub transaction_or_receipt_id: TransactionOrReceiptId,
    pub payloads: Vec<Vec<u8>>,
}

pub(crate) struct BlockListener {
    stream: mpsc::Receiver<near_indexer::StreamerMessage>,
    receipt_sender: mpsc::Sender<CandidateData>,
    da_contract_id: AccountId,
}

impl BlockListener {
    pub(crate) fn new(
        stream: mpsc::Receiver<near_indexer::StreamerMessage>,
        receipt_sender: mpsc::Sender<CandidateData>,
        da_contract_id: AccountId,
    ) -> Self {
        Self {
            stream,
            receipt_sender,
            da_contract_id,
        }
    }

    fn receipt_filer_map(da_contract_id: &AccountId, receipt: ReceiptView) -> Option<CandidateData> {
        if &receipt.receiver_id != da_contract_id {
            return None;
        }

        let actions = if let ReceiptEnumView::Action { actions, .. } = receipt.receipt {
            actions
        } else {
            return None;
        };

        let payloads = actions
            .into_iter()
            .filter_map(|el| match el {
                ActionView::FunctionCall { method_name, args, .. } if method_name == "submit" => Some(args.into()),
                _ => None,
            })
            .collect::<Vec<Vec<u8>>>();

        if payloads.is_empty() {
            return None;
        }

        Some(CandidateData {
            transaction_or_receipt_id: TransactionOrReceiptId::Receipt {
                receipt_id: receipt.receipt_id,
                receiver_id: receipt.receiver_id,
            },
            payloads,
        })
    }

    pub(crate) async fn start(self) -> Result<()> {
        let Self {
            mut stream,
            receipt_sender,
            da_contract_id,
        } = self;

        while let Some(streamer_message) = stream.recv().await {
            // TODO: check receipt_receiver is closed?
            let candidates_data: Vec<CandidateData> = streamer_message
                .shards
                .into_iter()
                .flat_map(|shard| shard.chunk)
                .flat_map(|chunk| {
                    chunk
                        .receipts
                        .into_iter()
                        .filter_map(|receipt| Self::receipt_filer_map(&da_contract_id, receipt))
                })
                .collect();

            if candidates_data.is_empty() {
                continue;
            }

            let results = join_all(candidates_data.into_iter().map(|receipt| receipt_sender.send(receipt))).await;
            results.into_iter().collect::<Result<_, _>>()?;
        }

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use std::path::PathBuf;
    use crate::block_listener::{BlockListener, CandidateData};
    use near_crypto::{KeyType, PublicKey};
    use near_indexer::near_primitives::hash::CryptoHash;
    use near_indexer::near_primitives::types::{AccountId, TransactionOrReceiptId};
    use near_indexer::near_primitives::views::{ActionView, ReceiptEnumView, ReceiptView};
    use std::str::FromStr;
    use std::time::Duration;
    use tokio::sync::mpsc;
    use near_indexer::StreamerMessage;
    use tokio::sync::mpsc::error::TryRecvError;

    impl PartialEq for CandidateData {
        fn eq(&self, other: &Self) -> bool {
            let res = match (&self.transaction_or_receipt_id, &other.transaction_or_receipt_id) {
                (
                    TransactionOrReceiptId::Receipt {
                        receiver_id,
                        receipt_id,
                    },
                    TransactionOrReceiptId::Receipt {
                        receipt_id: other_receipt_id,
                        receiver_id: other_receiver_id,
                    },
                ) => return receipt_id == other_receipt_id && receiver_id == other_receiver_id,
                (
                    TransactionOrReceiptId::Transaction {
                        transaction_hash,
                        sender_id,
                    },
                    TransactionOrReceiptId::Transaction {
                        transaction_hash: other_hash,
                        sender_id: other_id,
                    },
                ) => return transaction_hash == other_hash && sender_id == other_id,
                _ => false,
            };

            self.payloads == other.payloads && res
        }
    }

    #[test]
    fn test_receipt_filter_map() {
        let da_contract_id = AccountId::from_str("a.test").unwrap();

        let common_action_receipt = ReceiptEnumView::Action {
            signer_id: AccountId::from_str("test_signer").unwrap(),
            signer_public_key: PublicKey::empty(KeyType::ED25519),
            gas_price: 100,
            output_data_receivers: vec![],
            input_data_ids: vec![CryptoHash::hash_bytes(b"test_input_data_id")],
            actions: vec![ActionView::FunctionCall {
                method_name: "submit".into(),
                args: vec![1, 2, 3].into(),
                gas: 100,
                deposit: 100,
            }],
        };

        let test_predecessor_id = AccountId::from_str("test_predecessor").unwrap();
        let valid_receipt = ReceiptView {
            predecessor_id: test_predecessor_id.clone(),
            receiver_id: da_contract_id.clone(),
            receipt_id: CryptoHash::hash_bytes(b"test_receipt_id"),
            receipt: common_action_receipt.clone(),
        };

        let invalid_receiver_receipt = ReceiptView {
            predecessor_id: test_predecessor_id.clone(),
            receiver_id: AccountId::from_str("other_contract").unwrap(),
            receipt_id: CryptoHash::hash_bytes(b"test_receipt_id"),
            receipt: common_action_receipt,
        };

        let invalid_action_receipt = ReceiptView {
            predecessor_id: test_predecessor_id.clone(),
            receiver_id: da_contract_id.clone(),
            receipt_id: CryptoHash::hash_bytes(b"test_receipt_id"),
            receipt: ReceiptEnumView::Data {
                data_id: CryptoHash::hash_bytes(b"test_data_id"),
                data: Some(vec![1, 2, 3]),
            },
        };

        // Test valid receipt
        assert_eq!(
            BlockListener::receipt_filer_map(&da_contract_id, valid_receipt.clone()),
            Some(CandidateData {
                transaction_or_receipt_id: TransactionOrReceiptId::Receipt {
                    receipt_id: valid_receipt.receipt_id.clone(),
                    receiver_id: valid_receipt.receiver_id.clone(),
                },
                payloads: vec![vec![1, 2, 3]],
            })
        );

        // Test invalid receiver receipt
        assert_eq!(
            BlockListener::receipt_filer_map(&da_contract_id, invalid_receiver_receipt),
            None
        );

        // Test invalid action receipt
        assert_eq!(
            BlockListener::receipt_filer_map(&da_contract_id, invalid_action_receipt),
            None
        );
    }

    #[test]
    fn test_multiple_submit_actions() {
        let da_contract_id = AccountId::from_str("a.test").unwrap();

        let common_action_receipt = ReceiptEnumView::Action {
            signer_id: AccountId::from_str("test_signer").unwrap(),
            signer_public_key: PublicKey::empty(KeyType::ED25519),
            gas_price: 100,
            output_data_receivers: vec![],
            input_data_ids: vec![CryptoHash::hash_bytes(b"test_input_data_id")],
            actions: vec![
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
            ],
        };

        let test_predecessor_id = AccountId::from_str("test_predecessor").unwrap();
        let valid_receipt = ReceiptView {
            predecessor_id: test_predecessor_id.clone(),
            receiver_id: da_contract_id.clone(),
            receipt_id: CryptoHash::hash_bytes(b"test_receipt_id"),
            receipt: common_action_receipt.clone(),
        };

        // Test valid receipt
        assert_eq!(
            BlockListener::receipt_filer_map(&da_contract_id, valid_receipt.clone()),
            Some(CandidateData {
                transaction_or_receipt_id: TransactionOrReceiptId::Receipt {
                    receipt_id: valid_receipt.receipt_id.clone(),
                    receiver_id: valid_receipt.receiver_id.clone(),
                },
                payloads: vec![vec![1, 2, 3], vec![4, 4, 4]],
            })
        );
    }

    struct StreamerMessages {
        pub empty: Vec<StreamerMessage>,
        pub candidates: Vec<StreamerMessage>
    }

    struct StreamerMessagesLoader;
    impl StreamerMessagesLoader {
        fn load() -> StreamerMessages {
            let test_data_dir = concat!(env!("CARGO_MANIFEST_DIR"), "/test_data");
            let empty_data_path = [test_data_dir, "/empty"].concat();
            let candidates_data_path = [test_data_dir, "/candidates"].concat();

            StreamerMessages {
                empty: Self::load_messages(&empty_data_path),
                candidates: Self::load_messages(&candidates_data_path)
            }
        }

        fn load_messages(dir_path: &str) -> Vec<StreamerMessage> {
            let files = std::fs::read_dir(dir_path).map(|entry| {
                entry.map(|dir_entry| {
                    let dir_entry = dir_entry.unwrap();
                     dir_entry.path()
                }).collect::<Vec<PathBuf>>()
            }).unwrap();

            files.into_iter().map(|file| {
                let file = std::fs::File::open(file).unwrap();
                serde_json::from_reader(file).unwrap()
            }).collect()
        }
    }

    #[tokio::test]
    async fn test_empty_listener() {
        let da_contract_id = "da.test.near".parse().unwrap();
        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, mut receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, da_contract_id);
        let _ = tokio::spawn(listener.start());

        for el in streamer_messages.empty {
            sender.send(el).await.unwrap();
        }

        tokio::time::sleep(Duration::from_millis(10)).await;

        assert_eq!(receipt_receiver.try_recv(), Err(TryRecvError::Empty));
    }

    #[tokio::test]
    async fn test_candidates_listener() {
        let da_contract_id = "da.test.near".parse().unwrap();
        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, mut receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, da_contract_id);
        let handle = tokio::spawn(listener.start());

        let expected = streamer_messages.candidates.len();
        for el in streamer_messages.candidates {
            sender.send(el).await.unwrap();
        }

        drop(sender);
        handle.await.unwrap().unwrap();

        let mut counter = 0;
        loop {
            match receipt_receiver.recv().await {
                Some(_) => counter += 1,
                None => break
            }
        }

        assert_eq!(expected, counter);
    }

    #[tokio::test]
    async fn test_shutdown() {
        let da_contract_id = "da.test.near".parse().unwrap();
        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, da_contract_id);
        let handle = tokio::spawn(listener.start());
        for (i, el) in streamer_messages.empty.into_iter().enumerate() {
            sender.send(el).await.unwrap();

            // some random number
            if i == 5 {
                break;
            }
        }

        drop(receipt_receiver);
        let _ = sender.send(streamer_messages.candidates.into_iter().next_back().unwrap()).await;

        // Assert tha handle terminated with error
        assert!(handle.await.unwrap().is_err());
    }
}
