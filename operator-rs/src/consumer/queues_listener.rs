use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use alloy_rpc_types::Block;
use tokio::sync::broadcast;
use anyhow::{Result, anyhow};
use lapin::{message::Delivery, Consumer};
use tracing::{info, error};
use borsh::BorshDeserialize;
use alloy_rlp::{Rlp, Decodable};
use futures_util::StreamExt;
use serde_json::Value;
use crate::types::BlockData;

use super::EventListener;
use crate::types::{Blob, Namespace, PublishPayload, SubmitRequest};

pub struct QueuesListener {
    received_blocks_tx: Arc<broadcast::Sender<BlockData>>,
    queue_consumers: Arc<Mutex<HashMap<u32, Consumer>>>,
    event_listener: Arc<dyn EventListener>,
}

impl QueuesListener {
    pub fn new(
        received_blocks_tx: Arc<broadcast::Sender<BlockData>>,
        event_listener: Arc<dyn EventListener>,
    ) -> Self {
        Self {
            received_blocks_tx,
            queue_consumers: Arc::new(Mutex::new(HashMap::new())),
            event_listener,
        }
    }

    pub async fn add(&self, rollup_id: u32, consumer: Consumer) -> Result<()> {
        let mut consumers = self.queue_consumers.lock().unwrap();
        consumers.insert(rollup_id, consumer);
        Ok(())
    }

    pub fn remove(&self, rollup_id: u32) {
        let mut consumers = self.queue_consumers.lock().unwrap();
        consumers.remove(&rollup_id);
    }

    pub async fn listen(&self, rollup_id: u32) -> Result<()> {
        let mut consumer = {
            let consumers = self.queue_consumers.lock().unwrap();
            consumers.get(&rollup_id).cloned().ok_or_else(|| anyhow!("Consumer not found"))?
        };
        
        loop {
            match consumer.next().await {
                Some(Ok(delivery)) => {
                    if let Err(e) = self.process_delivery(rollup_id, delivery).await {
                        error!("Failed to process delivery: {:?}", e);
                    }
                }
                Some(Err(e)) => error!("Failed to get delivery: {:?}", e),
                None => {
                    info!("Consumer channel closed");
                    return Ok(());
                }
            }
        }
    }

    async fn process_delivery(&self, rollup_id: u32, delivery: Delivery) -> Result<()> {
        info!("New delivery, rollup_id: {}", rollup_id);
        self.event_listener.on_arrival();

        let publish_payload = PublishPayload::try_from_slice(&delivery.data)
            .map_err(|e| {
                self.event_listener.on_format_error();
                anyhow!("Error deserializing payload: {:?}", e)
            })?;

        let submit_request = SubmitRequest::try_from_slice(&publish_payload.data)
            .map_err(|e| {
                self.event_listener.on_format_error();
                anyhow!("Invalid blob: {:?}", e)
            })?;

        for blob in submit_request.blobs {
            // Decode the blob data into Blocks
            let mut rlp = Rlp::new(&blob.data)?;
            let mut blocks = Vec::new();
            if let Some(block_vec) = rlp.get_next::<Vec<Vec<u8>>>()? {
                for block_bytes in block_vec {
                    let json_block: Value = serde_json::from_slice(&block_bytes)
                        .map_err(|e| anyhow!("Failed to parse JSON block: {:?}", e))?;
                    let block: Block = serde_json::from_value(json_block)
                        .map_err(|e| anyhow!("Failed to convert JSON to Block: {:?}", e))?;
                    blocks.push(block);
                }
            };

            for block in blocks {
                let block_data = BlockData {
                    rollup_id,
                    transaction_id: publish_payload.transaction_id,
                    commitment: blob.commitment,
                    block
                };

                info!(
                    "MQ Block received: rollup_id={}, block_height={}, transaction_id={}, commitment={}",
                    rollup_id,
                    block_data.block.header.number,
                    hex::encode(block_data.transaction_id),
                    hex::encode(block_data.commitment)
                );

                if let Err(e) = self.received_blocks_tx.send(block_data) {
                    error!("Failed to send block data: {:?}", e);
                }
            }
        }

        delivery.ack(Default::default()).await
            .map_err(|e| anyhow!("Failed to ack delivery: {:?}", e))?;

        Ok(())
    }
}
