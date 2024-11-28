use near_indexer::near_primitives::{types::AccountId, views::{ActionView, ExecutionStatusView, ReceiptEnumView}};
use reqwest::Client;
use std::collections::HashMap;
use tokio::sync::{mpsc, mpsc::{Receiver, Sender}};
use tracing::{debug, error, info, trace};

use crate::errors::Result;
use crate::metrics::{make_block_listener_metrics, BlockEventListener, Metricable};
use crate::{errors::Error, rmq_publisher::{get_routing_key, PublishData, PublishOptions, PublishPayload, PublisherContext}, types::{BlockWithTxHashes, IndexerExecutionOutcomeWithReceiptAndTxHash, PartialCandidateData, PartialCandidateDataWithBlockTxHash}};

const FASTNEAR_ENDPOINT: &str = "https://testnet.neardata.xyz/v0/last_block/final";

pub struct FastNearIndexer {
    client: Client,
    addresses_to_rollup_ids: HashMap<AccountId, u32>,
    listener: Option<BlockEventListener>,
}

impl FastNearIndexer {
    pub(crate) fn new(addresses_to_rollup_ids: HashMap<AccountId, u32>) -> Self {
        debug!(target: "fastnear_indexer", "Creating new FastNearIndexer");
        Self { 
            client: Client::new(),
            addresses_to_rollup_ids,
            listener: None,
        }
    }

    pub fn run(&self) -> Receiver<PublishData> {
        info!(target: "fastnear_indexer", "Starting FastNearIndexer");
        let block_receiver = self.stream_latest_blocks();
        let (publish_sender, publish_receiver) = mpsc::channel(100);

        let addresses_to_rollup_ids = self.addresses_to_rollup_ids.clone();

        tokio::spawn(async move {
            Self::process_blocks(block_receiver, publish_sender, addresses_to_rollup_ids).await;
        });

        publish_receiver
    }

    async fn process_blocks(
        mut block_receiver: Receiver<BlockWithTxHashes>,
        publish_sender: Sender<PublishData>,
        addresses_to_rollup_ids: HashMap<AccountId, u32>,
    ) {
        debug!(target: "fastnear_indexer", "Starting block processing");
        while let Some(block) = block_receiver.recv().await {
            trace!(target: "fastnear_indexer", "Received block: {:?}", block.block.header.height);
            if let Err(e) = Self::parse_and_publish_block(block, &publish_sender, &addresses_to_rollup_ids).await {
                error!(target: "fastnear_indexer", "Error parsing and publishing block: {:?}", e);
            }
        }
    }

    async fn parse_and_publish_block(
        block: BlockWithTxHashes,
        publish_sender: &Sender<PublishData>,
        addresses_to_rollup_ids: &HashMap<AccountId, u32>,
    ) -> Result<(), Error> {
        debug!(target: "fastnear_indexer", "Parsing block: {:?}", block.block.header.height);
        for shard in block.shards {
            for receipt_execution_outcome in shard.receipt_execution_outcomes {
                let receiver_id = &receipt_execution_outcome.receipt.receiver_id;
                debug!(target: "fastnear_indexer", "Processing receipt for receiver_id: {}", receiver_id);
                if let Some(rollup_id) = addresses_to_rollup_ids.get(receiver_id) {
                    trace!(target: "fastnear_indexer", "Processing receipt for rollup_id: {}", rollup_id);
                    if !Self::is_successful_execution(&receipt_execution_outcome) {
                        trace!(target: "fastnear_indexer", "Skipping unsuccessful execution for rollup_id: {}", rollup_id);
                        continue;
                    }

                    let partial_candidate_data = Self::receipt_filter_map(
                        receipt_execution_outcome.receipt.receipt,
                        *rollup_id
                    );

                    if let (Some(partial_data), Some(tx_hash)) = (partial_candidate_data, receipt_execution_outcome.tx_hash) {
                        let candidate_data = PartialCandidateDataWithBlockTxHash {
                            rollup_id: *rollup_id,
                            payloads: partial_data.payloads,
                            tx_hash,
                            block_hash: block.block.header.hash,
                        };
                        debug!(target: "fastnear_indexer", "Sending candidate data for rollup_id: {}", rollup_id);
                        Self::send(&candidate_data, publish_sender).await?;
                    }
                }
            }
        }

        Ok(())
    }

    pub fn stream_latest_blocks(&self) -> mpsc::Receiver<BlockWithTxHashes> {
        info!(target: "fastnear_indexer", "Starting block stream");
        let (block_sender, block_receiver) = mpsc::channel(100);
        let client = self.client.clone();

        tokio::spawn(async move {
            loop {
                match Self::fetch_latest_block(&client).await {
                    Ok(block) => {
                        if block_sender.send(block.clone()).await.is_err() {
                            error!(target: "fastnear_indexer", "Failed to send block to channel");
                            break;
                        }
                        info!(target: "fastnear_indexer", "Successfully fetched and sent latest block with id: {}", block.block.header.height);
                    }
                    Err(e) => error!(target: "fastnear_indexer", "Error fetching latest block: {:?}", e),
                }
                tokio::time::sleep(std::time::Duration::from_secs(1)).await;
            }
        });

        block_receiver
    }

    async fn fetch_latest_block(client: &Client) -> Result<BlockWithTxHashes, Error> {
        debug!(target: "fastnear_indexer", "Fetching latest block");
        let response = client.get(FASTNEAR_ENDPOINT)
            .send()
            .await
            .and_then(|r| r.error_for_status())
            .map_err(|e| Error::NetworkError(e.to_string()))?;

        response.json::<BlockWithTxHashes>()
            .await
            .map_err(|e| Error::DeserializeJsonError(e.to_string()))
    }

    async fn send(candidate_data: &PartialCandidateDataWithBlockTxHash, sender: &Sender<PublishData>) -> Result<(), Error> {
        trace!(target: "fastnear_indexer", "Sending candidate data: {:?}", candidate_data);
        for data in candidate_data.clone().payloads {
            let publish_data = PublishData {
                publish_options: PublishOptions {
                    routing_key: get_routing_key(candidate_data.rollup_id),
                    ..PublishOptions::default()
                },
                cx: PublisherContext {
                    block_hash: candidate_data.block_hash,
                },
                payload: PublishPayload {
                    transaction_id: candidate_data.tx_hash,
                    data,
                },
            };
            sender.send(publish_data).await?
        }

        Ok(())
    }

    fn is_successful_execution(receipt_execution_outcome: &IndexerExecutionOutcomeWithReceiptAndTxHash) -> bool {
        let is_successful = matches!(
            receipt_execution_outcome.execution_outcome.outcome.status,
            ExecutionStatusView::SuccessValue(ref value) if value.is_empty()
        );
        trace!(target: "fastnear_indexer", "Execution successful: {}", is_successful);
        is_successful
    }

    fn receipt_filter_map(receipt_enum_view: ReceiptEnumView, rollup_id: u32) -> Option<PartialCandidateData> {
        trace!(target: "fastnear_indexer", "Filtering receipt for rollup_id: {}", rollup_id);
        let payloads = match receipt_enum_view {
            ReceiptEnumView::Action { actions, .. } => {
                actions.into_iter()
                    .filter_map(Self::extract_args)
                    .collect::<Vec<Vec<u8>>>()
            }
            _ => return None,
        };

        if payloads.is_empty() {
            trace!(target: "fastnear_indexer", "No payloads found for rollup_id: {}", rollup_id);
            return None;
        }

        Some(PartialCandidateData {
            rollup_id,
            payloads,
        })
    }

    fn extract_args(action: ActionView) -> Option<Vec<u8>> {
        match action {
            ActionView::FunctionCall { method_name, args, .. } if method_name == "submit" => {
                trace!(target: "fastnear_indexer", "Extracted args for 'submit' method");
                Some(args.into())
            },
            _ => {
                trace!(target: "fastnear_indexer", "Skipped non-'submit' method");
                None
            },
        }
    }
}

impl Metricable for FastNearIndexer {
    fn enable_metrics(&mut self, registry: prometheus::Registry) -> Result<()> {
        let listener = make_block_listener_metrics(registry)?;
        self.listener = Some(listener);

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use near_crypto::{KeyType, PublicKey};
    use near_indexer::near_primitives::views::{ActionView, ReceiptEnumView};
    use reqwest::Client;
    use std::collections::HashMap;

    #[tokio::test]
    async fn test_run() {
        let addresses_to_rollup_ids = HashMap::new();
        let indexer = FastNearIndexer::new(addresses_to_rollup_ids);
        let receiver = indexer.run();
        // Since the run method spawns asynchronous tasks, we can check if the receiver is valid
        assert!(receiver.capacity() > 0);
    }

    #[tokio::test]
    async fn test_fetch_latest_block() {
        let client = Client::new();
        let result = FastNearIndexer::fetch_latest_block(&client).await;
        assert!(result.is_ok(), "Failed to fetch latest block");
        let block = result.unwrap();
        // Basic assertions about the block
        assert!(block.block.header.height > 0, "Block height should be greater than 0");
    }

    #[tokio::test]
    async fn test_extract_args_with_submit_method() {
        let action = ActionView::FunctionCall {
            method_name: "submit".to_string(),
            args: vec![1, 2, 3].into(),
            gas: 1000,
            deposit: 0,
        };
        let args = FastNearIndexer::extract_args(action);
        assert!(args.is_some());
        assert_eq!(args.unwrap(), vec![1, 2, 3]);
    }

    #[tokio::test]
    async fn test_extract_args_with_other_method() {
        let action = ActionView::FunctionCall {
            method_name: "other_method".to_string(),
            args: vec![1, 2, 3].into(),
            gas: 1000,
            deposit: 0,
        };
        let args = FastNearIndexer::extract_args(action);
        assert!(args.is_none());
    }

    #[tokio::test]
    async fn test_receipt_filter_map_with_matching_rollup_id() {
        let rollup_id = 1;
        let action = ActionView::FunctionCall {
            method_name: "submit".to_string(),
            args: vec![1, 2, 3].into(),
            gas: 1000,
            deposit: 0,
        };
        let receipt_enum_view = ReceiptEnumView::Action {
            signer_id: "signer.near".to_string().try_into().unwrap(),
            signer_public_key: PublicKey::empty(KeyType::ED25519),
            gas_price: 0,
            output_data_receivers: vec![],
            input_data_ids: vec![],
            actions: vec![action],
            is_promise_yield: false,
        };
        let result = FastNearIndexer::receipt_filter_map(receipt_enum_view, rollup_id);
        assert!(result.is_some());
        let partial_candidate = result.unwrap();
        assert_eq!(partial_candidate.rollup_id, rollup_id);
        assert_eq!(partial_candidate.payloads, vec![vec![1, 2, 3]]);
    }

    #[tokio::test]
    async fn test_receipt_filter_map_with_non_matching_rollup_id() {
        let rollup_id = 1;
        let receipt_enum_view = ReceiptEnumView::Action {
            signer_id: "signer.near".to_string().try_into().unwrap(),
            signer_public_key: PublicKey::empty(KeyType::ED25519),
            gas_price: 0,
            output_data_receivers: vec![],
            input_data_ids: vec![],
            actions: vec![],
            is_promise_yield: false,
        };
        let result = FastNearIndexer::receipt_filter_map(receipt_enum_view, rollup_id);
        assert!(result.is_none());
    }
}