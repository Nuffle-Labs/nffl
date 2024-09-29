//! FastNearIndexer module for efficient indexing of NEAR blockchain data.
//!
//! This module provides functionality to stream and process the latest blocks
//! from the NEAR blockchain, focusing on specific transactions and receipts
//! related to configured rollup addresses.

use std::collections::HashMap;
use near_indexer::near_primitives::{types::AccountId, views::{ActionView, ExecutionStatusView, ReceiptEnumView}};
use reqwest::Client;
use tokio::sync::{mpsc::{Sender, Receiver}, mpsc};
use tracing::{info, error, debug, trace};

use crate::{errors::Error, rmq_publisher::{get_routing_key, PublishData, PublishOptions, PublishPayload, PublisherContext}, types::{BlockWithTxHashes, IndexerExecutionOutcomeWithReceiptAndTxHash, PartialCandidateData, PartialCandidateDataWithBlockTxHash}};

/// The endpoint URL for fetching the latest finalized block from NEAR testnet.
const FASTNEAR_ENDPOINT: &str = "https://testnet.neardata.xyz/v0/last_block/final";

/// Represents the FastNearIndexer, which processes NEAR blockchain data.
#[derive(Debug)]
pub struct FastNearIndexer {
    /// HTTP client for making requests to the NEAR API.
    client: Client,
    /// Mapping of account IDs to their corresponding rollup IDs.
    addresses_to_rollup_ids: HashMap<AccountId, u32>,
}

impl FastNearIndexer {
    /// Creates a new instance of FastNearIndexer.
    ///
    /// # Arguments
    ///
    /// * `addresses_to_rollup_ids` - A HashMap mapping AccountIds to their corresponding rollup IDs.
    ///
    /// # Returns
    ///
    /// A new `FastNearIndexer` instance.
    pub(crate) fn new(addresses_to_rollup_ids: HashMap<AccountId, u32>) -> Self {
        debug!(target: "fastnear_indexer", "Creating new FastNearIndexer");
        Self {
            client: Client::new(),
            addresses_to_rollup_ids,
        }
    }

    /// Starts the indexing process and returns a receiver for publish data.
    ///
    /// # Returns
    ///
    /// A `Receiver<PublishData>` for consuming the processed blockchain data.
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

    /// Processes incoming blocks and publishes relevant data.
    ///
    /// # Arguments
    ///
    /// * `block_receiver` - A receiver for incoming `BlockWithTxHashes`.
    /// * `publish_sender` - A sender for outgoing `PublishData`.
    /// * `addresses_to_rollup_ids` - A HashMap mapping AccountIds to their corresponding rollup IDs.
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

    /// Parses a block and publishes relevant data.
    ///
    /// # Arguments
    ///
    /// * `block` - The `BlockWithTxHashes` to be parsed.
    /// * `publish_sender` - A sender for outgoing `PublishData`.
    /// * `addresses_to_rollup_ids` - A HashMap mapping AccountIds to their corresponding rollup IDs.
    ///
    /// # Returns
    ///
    /// A `Result` indicating success or containing an `Error`.
    async fn parse_and_publish_block(
        block: BlockWithTxHashes,
        publish_sender: &Sender<PublishData>,
        addresses_to_rollup_ids: &HashMap<AccountId, u32>,
    ) -> Result<(), Error> {
        debug!(target: "fastnear_indexer", "Parsing block: {:?}", block.block.header.height);
        for shard in block.shards {
            for receipt_execution_outcome in shard.receipt_execution_outcomes {
                let receiver_id = &receipt_execution_outcome.receipt.receiver_id;
                
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

    /// Creates a stream of the latest blocks from the NEAR blockchain.
    ///
    /// # Returns
    ///
    /// A `Receiver<BlockWithTxHashes>` for consuming the latest blocks.
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

    /// Fetches the latest block from the NEAR API.
    ///
    /// # Arguments
    ///
    /// * `client` - The HTTP client to use for the API request.
    ///
    /// # Returns
    ///
    /// A `Result` containing the `BlockWithTxHashes` or an `Error`.
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

    /// Sends candidate data to the publish channel.
    ///
    /// # Arguments
    ///
    /// * `candidate_data` - The `PartialCandidateDataWithBlockTxHash` to be sent.
    /// * `sender` - The `Sender<PublishData>` to send the data through.
    ///
    /// # Returns
    ///
    /// A `Result` indicating success or containing an `Error`.
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

    /// Checks if the execution of a receipt was successful.
    ///
    /// # Arguments
    ///
    /// * `receipt_execution_outcome` - The `IndexerExecutionOutcomeWithReceiptAndTxHash` to check.
    ///
    /// # Returns
    ///
    /// A boolean indicating whether the execution was successful.
    fn is_successful_execution(receipt_execution_outcome: &IndexerExecutionOutcomeWithReceiptAndTxHash) -> bool {
        let is_successful = matches!(
            receipt_execution_outcome.execution_outcome.outcome.status,
            ExecutionStatusView::SuccessValue(ref value) if value.is_empty()
        );
        trace!(target: "fastnear_indexer", "Execution successful: {}", is_successful);
        is_successful
    }

    /// Filters and maps a receipt to partial candidate data.
    ///
    /// # Arguments
    ///
    /// * `receipt_enum_view` - The `ReceiptEnumView` to be filtered and mapped.
    /// * `rollup_id` - The rollup ID associated with this receipt.
    ///
    /// # Returns
    ///
    /// An `Option<PartialCandidateData>` containing the filtered and mapped data, if any.
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

    /// Extracts arguments from an action view.
    ///
    /// # Arguments
    ///
    /// * `action` - The `ActionView` to extract arguments from.
    ///
    /// # Returns
    ///
    /// An `Option<Vec<u8>>` containing the extracted arguments, if any.
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
