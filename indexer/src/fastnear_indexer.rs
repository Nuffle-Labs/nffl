use reqwest::Client;
use tokio::sync::mpsc;
use tracing::{info, error};

use crate::types::BlockWithTxHashes;

const FASTNEAR_ENDPOINT: &str = "https://testnet.neardata.xyz/v0/last_block/final";


#[derive(Debug)]
pub struct FastNearIndexer {
    client: Client,
}


impl FastNearIndexer {
    pub fn new() -> Self {
        FastNearIndexer {
            client: Client::new(),
        }
    }

    pub fn stream_latest_blocks(&self) -> mpsc::Receiver<BlockWithTxHashes> {
        let (sender, receiver) = mpsc::channel(100);
        let client = self.client.clone();

        tokio::spawn(async move {
            loop {
                match client.get(FASTNEAR_ENDPOINT).send().await.and_then(|resp| resp.error_for_status()) {
                    Ok(response) => {
                        if let Ok(block) = response.json::<BlockWithTxHashes>().await {
                            if sender.send(block.clone()).await.is_err() {
                                error!(target: "fastnear_indexer", "Failed to send block to channel");
                                break;
                            }
                            info!(target: "fastnear_indexer", "Successfully fetched and sent latest block with id: {}", block.block.header.height);
                        } else {
                            error!(target: "fastnear_indexer", "Failed to deserialize response into StreamerMessage");
                        }
                    }
                    Err(e) => error!(target: "fastnear_indexer", "Error fetching latest block: {:?}", e),
                }
                tokio::time::sleep(std::time::Duration::from_secs(1)).await;
            }
        });

        receiver
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn test_stream_latest_blocks() {
        let indexer = FastNearIndexer::new();
        let stream = indexer.stream_latest_blocks();
        assert!(stream.is_empty());
    }
}
