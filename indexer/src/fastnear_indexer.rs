use std::collections::HashMap;
use near_indexer::near_primitives::{types::AccountId, views::{ActionView, ExecutionStatusView, ReceiptEnumView}};
use reqwest::Client;
use tokio::sync::{mpsc::{Sender, Receiver}, mpsc};
use tracing::{info, error};

use crate::{errors::Error, rmq_publisher::{get_routing_key, PublishData, PublishOptions, PublishPayload, PublisherContext, RmqPublisherHandle}, types::{BlockWithTxHashes, CandidateData, IndexerExecutionOutcomeWithReceiptAndTxHash, PartialCandidateData, PartialCandidateDataWithBlockTxHash}};

const FASTNEAR_ENDPOINT: &str = "https://testnet.neardata.xyz/v0/last_block/final";

#[derive(Debug)]
pub struct FastNearIndexer {
    client: Client,
    addresses_to_rollup_ids: HashMap<AccountId, u32>,
}

impl FastNearIndexer {
    pub(crate) fn new(addresses_to_rollup_ids: HashMap<AccountId, u32>) -> Self {
        Self {
            client: Client::new(),
            addresses_to_rollup_ids,
        }
    }

    pub fn run(&self) -> Receiver<PublishData> {
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
        while let Some(block) = block_receiver.recv().await {
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
        for shard in block.shards {
            for receipt_execution_outcome in shard.receipt_execution_outcomes {
                let receiver_id = &receipt_execution_outcome.receipt.receiver_id;
                
                if let Some(rollup_id) = addresses_to_rollup_ids.get(receiver_id) {
                    if !Self::is_successful_execution(&receipt_execution_outcome) {
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
                        Self::send(&candidate_data, publish_sender).await?;
                    }
                }
            }
        }

        Ok(())
    }

    pub fn stream_latest_blocks(&self) -> mpsc::Receiver<BlockWithTxHashes> {
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
        let response = client.get(FASTNEAR_ENDPOINT)
            .send()
            .await
            .map_err(|e| Error::NetworkError(e.to_string()))?;

        if !response.status().is_success() {
            return Err(Error::ApiError(format!("API request failed with status: {}", response.status())));
        }

        response.json::<BlockWithTxHashes>()
            .await
            .map_err(|e| Error::DeserializeJsonError(e.to_string()))
    }

    // Update the send method to use Sender<PublishData> directly
    async fn send(candidate_data: &PartialCandidateDataWithBlockTxHash, sender: &Sender<PublishData>) -> Result<(), Error> {
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

    // Make this method static as it doesn't use &self
    fn is_successful_execution(receipt_execution_outcome: &IndexerExecutionOutcomeWithReceiptAndTxHash) -> bool {
        matches!(
            receipt_execution_outcome.execution_outcome.outcome.status,
            ExecutionStatusView::SuccessValue(ref value) if value.is_empty()
        )
    }

    fn receipt_filter_map(receipt_enum_view: ReceiptEnumView, rollup_id: u32) -> Option<PartialCandidateData> {
        let payloads = match receipt_enum_view {
            ReceiptEnumView::Action { actions, .. } => {
                actions.into_iter()
                    .filter_map(Self::extract_args)
                    .collect::<Vec<Vec<u8>>>()
            }
            _ => return None,
        };

        if payloads.is_empty() {
            return None;
        }

        Some(PartialCandidateData {
            rollup_id,
            payloads,
        })
    }

    fn extract_args(action: ActionView) -> Option<Vec<u8>> {
        match action {
            ActionView::FunctionCall { method_name, args, .. } if method_name == "submit" => Some(args.into()),
            _ => None,
        }
    }
}
