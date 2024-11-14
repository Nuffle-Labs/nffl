use near_indexer::{
    near_primitives::{types::AccountId, views::ActionView},
    StreamerMessage,
};
use prometheus::Registry;
use std::{collections::HashMap, collections::VecDeque, sync, time::Duration};
use tokio::{
    sync::{
        mpsc::{self, error::TrySendError, Receiver},
        oneshot, Mutex,
    },
    task::JoinHandle,
    time,
};
use tracing::info;

use crate::{
    errors::Result,
    metrics::{make_block_listener_metrics, BlockEventListener, Metricable},
    types,
    types::CandidateData,
    INDEXER,
};

#[cfg(not(test))]
const EXPIRATION_TIMEOUT: Duration = Duration::from_secs(30);
#[cfg(test)]
const EXPIRATION_TIMEOUT: Duration = Duration::from_millis(200);

#[derive(Clone)]
struct ExpirableCandidateData {
    timestamp: time::Instant,
    inner: CandidateData,
}

impl From<ExpirableCandidateData> for CandidateData {
    fn from(value: ExpirableCandidateData) -> Self {
        value.inner
    }
}

#[derive(Clone, Debug, serde::Serialize, serde::Deserialize)]
struct TransactionWithRollupId {
    pub(crate) rollup_id: u32,
    pub(crate) transaction: near_indexer::IndexerTransactionWithOutcome,
}

#[derive(Clone)]
pub(crate) struct BlockListener {
    addresses_to_rollup_ids: HashMap<AccountId, u32>,
    listener: Option<BlockEventListener>,
}

impl BlockListener {
    pub(crate) fn new(addresses_to_rollup_ids: HashMap<AccountId, u32>) -> Self {
        Self {
            addresses_to_rollup_ids,
            listener: None,
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

    async fn ticker(
        mut done: oneshot::Receiver<()>,
        queue_protected: types::ProtectedQueue<ExpirableCandidateData>,
        candidates_sender: mpsc::Sender<CandidateData>,
        listener: Option<BlockEventListener>
    ) {
        #[cfg(not(test))]
        const FLUSH_INTERVAL: Duration = Duration::from_secs(1);
        #[cfg(test)]
        const FLUSH_INTERVAL: Duration = Duration::from_millis(100);

        let mut interval = time::interval(FLUSH_INTERVAL);
        loop {
            tokio::select! {
                _ = interval.tick() => {
                    let mut queue = queue_protected.lock().await;
                    let _ = Self::flush(&mut queue, &candidates_sender);
                    listener.as_ref().map(|l| l.current_queued_candidates.set(queue.len() as f64));

                    interval.reset();
                },
                _ = &mut done => {
                    return
                }
            }
        }
    }

    fn flush(queue: &mut VecDeque<ExpirableCandidateData>, candidates_sender: &mpsc::Sender<CandidateData>) -> bool {
        if queue.is_empty() {
            return true;
        }

        info!(target: INDEXER, "Flushing");

        let now = time::Instant::now();
        while let Some(candidate) = queue.front() {
            if now.duration_since(candidate.timestamp) >= EXPIRATION_TIMEOUT {
                queue.pop_front();
                continue;
            }

            match candidates_sender.try_send(candidate.clone().into()) {
                Ok(_) => {
                    let _ = queue.pop_front();
                }
                // TODO: return TrySendError instead
                Err(_) => return false,
            }
        }

        true
    }

    fn extract_candidates(addresses_to_rollup_ids: &HashMap<AccountId, u32>, streamer_message: StreamerMessage) -> Vec<CandidateData> {
        streamer_message
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
            .collect()
    }

    // TODO: introduce Task struct
    async fn process_stream(
        self,
        mut indexer_stream: Receiver<StreamerMessage>,
        candidates_sender: mpsc::Sender<CandidateData>,
    ) {
        let Self {
            addresses_to_rollup_ids,
            listener,
        } = self;

        let queue_protected = sync::Arc::new(Mutex::new(VecDeque::new()));
        let (done_sender, done_receiver) = oneshot::channel();
        actix::spawn(Self::ticker(
            done_receiver,
            queue_protected.clone(),
            candidates_sender.clone(),
            listener.clone(),
        ));

        while let Some(streamer_message) = indexer_stream.recv().await {
            info!(target: INDEXER, "Received streamer message");

            let candidates_data = Self::extract_candidates(&addresses_to_rollup_ids, streamer_message);
            if candidates_data.is_empty() {
                info!(target: INDEXER, "No candidate data found in the streamer message");
                continue;
            }

            // TODO: attempt flushing even if no new candidates or not?
            // Flushing old messages before new one
            {
                let mut queue = queue_protected.lock().await;
                let flushed = Self::flush(&mut queue, &candidates_sender);
                listener.as_ref().map(|l| l.current_queued_candidates.set(queue.len() as f64));
                if !flushed {
                    info!(target: INDEXER, "Not flushed, so enqueuing candidate data");

                    let timestamp = time::Instant::now();
                    queue.extend(
                        candidates_data
                            .into_iter()
                            .map(|el| ExpirableCandidateData { timestamp, inner: el }),
                    );
                    continue;
                }
            }

            {
                let candidates_len = candidates_data.len();
                info!(target: INDEXER, "Found {} candidate(s)", candidates_len);
                listener
                    .as_ref()
                    .map(|listener| listener.num_candidates.inc_by(candidates_len as f64));
            }

            let mut iter = candidates_data.into_iter();
            while let Some(candidate) = iter.next() {
                match candidates_sender.try_send(candidate) {
                    Ok(_) => {}
                    Err(err) => match err {
                        TrySendError::Full(candidate) => {
                            let mut queue = queue_protected.lock().await;

                            let timestamp = time::Instant::now();
                            queue.push_back(ExpirableCandidateData {
                                timestamp,
                                inner: candidate,
                            });
                            queue.extend(iter.map(|el| ExpirableCandidateData { timestamp, inner: el }));
                            listener.as_ref().map(|l| l.current_queued_candidates.set(queue.len() as f64));

                            break;
                        }
                        TrySendError::Closed(_) => {
                            return;
                        }
                    },
                }
            }
        }

        let _ = done_sender.send(());
    }

    /// Filters indexer stream and returns receiving channel.
    pub(crate) fn run(&self, indexer_stream: Receiver<StreamerMessage>) -> (JoinHandle<()>, Receiver<CandidateData>) {
        let (candidates_sender, candidates_receiver) = mpsc::channel(1000);
        let handle = actix::spawn(Self::process_stream(self.clone(), indexer_stream, candidates_sender));

        (handle, candidates_receiver)
    }

    #[cfg(test)]
    pub(crate) fn test_run(
        &self,
        indexer_stream: Receiver<StreamerMessage>,
    ) -> (JoinHandle<()>, Receiver<CandidateData>) {
        let (candidates_sender, candidates_receiver) = mpsc::channel(1);
        let handle = actix::spawn(Self::process_stream(self.clone(), indexer_stream, candidates_sender));

        (handle, candidates_receiver)
    }
}

impl Metricable for BlockListener {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()> {
        let listener = make_block_listener_metrics(registry)?;
        self.listener = Some(listener);

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use crate::block_listener::{BlockListener, TransactionWithRollupId, EXPIRATION_TIMEOUT};
    use crate::types::CandidateData;
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
                priority_fee: 0,
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
                priority_fee: 0,
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
                priority_fee: 0,
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

    #[actix::test]
    async fn test_empty_listener() {
        let rollup_id = 1;
        let da_contract_id: AccountId = "da.test.near".parse().unwrap();

        let streamer_messages = StreamerMessagesLoader::load();
        let (stream_sender, stream_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(HashMap::from([(da_contract_id, rollup_id)]));
        let (_, mut candidates_receiver) = listener.run(stream_receiver);

        for el in streamer_messages.empty {
            stream_sender.send(el).await.unwrap();
        }

        tokio::time::sleep(Duration::from_millis(10)).await;

        assert_eq!(candidates_receiver.try_recv(), Err(TryRecvError::Empty));
    }

    #[actix::test]
    async fn test_candidates_listener() {
        let rollup_id = 1;
        let da_contract_id = "da.test.near".parse().unwrap();

        let streamer_messages = StreamerMessagesLoader::load();
        let (stream_sender, stream_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(HashMap::from([(da_contract_id, rollup_id)]));
        let (handle, mut candidates_receiver) = listener.run(stream_receiver);

        let expected = streamer_messages.candidates.len();
        for el in streamer_messages.candidates {
            stream_sender.send(el).await.unwrap();
        }

        drop(stream_sender);
        handle.await.unwrap();

        let mut counter = 0;
        while let Some(_) = candidates_receiver.recv().await {
            counter += 1;
        }

        assert_eq!(expected, counter);
    }

    #[ignore]
    #[actix::test]
    async fn test_shutdown() {
        let rollup_id = 5;
        let da_contract_id = "da.test.near".parse().unwrap();

        let streamer_messages = StreamerMessagesLoader::load();
        let (stream_sender, stream_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(HashMap::from([(da_contract_id, rollup_id)]));
        let (handle, candidates_receiver) = listener.run(stream_receiver);

        for (i, el) in streamer_messages.empty.into_iter().enumerate() {
            stream_sender.send(el).await.unwrap();

            // some random number
            if i == 5 {
                break;
            }
        }

        drop(candidates_receiver);
        // Sender::closed is triggered
        assert!(handle.await.is_ok());
    }

    #[actix::test]
    async fn test_expiration() {
        let rollup_id = 1;
        let da_contract_id = "da.test.near".parse().unwrap();

        let streamer_messages = StreamerMessagesLoader::load();
        let (stream_sender, stream_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(HashMap::from([(da_contract_id, rollup_id)]));
        let (_, mut candidates_receiver) = listener.test_run(stream_receiver);

        for el in streamer_messages.candidates {
            stream_sender.send(el).await.unwrap();
        }

        // Let messages expire
        tokio::time::sleep(2 * EXPIRATION_TIMEOUT).await;

        // There shall be first message available
        assert!(candidates_receiver.try_recv().is_ok(), "Receiver shall have one value");
        // The rest shall be expired
        assert_eq!(Err(TryRecvError::Empty), candidates_receiver.try_recv());
    }
}
