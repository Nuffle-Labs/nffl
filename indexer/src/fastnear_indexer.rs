use reqwest::Client;
use tokio::sync::mpsc;
use near_indexer::StreamerMessage;
// const VERSION: &str = "v0";
// const NETWORK: &str = "testnet";
// const HEIGHT: u64 = 100000000;
// const LATEST: bool = true;
const FASTNEAR_ENDPOINT: &str = "https://testnet.neardata.xyz/v0/block/latest";

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

    pub fn stream_latest_blocks(&self) -> mpsc::Receiver<StreamerMessage> {
        let (sender, receiver) = mpsc::channel(100);
        let client = self.client.clone();

        tokio::spawn(async move {
            loop {
                match client.get(FASTNEAR_ENDPOINT).send().await {
                    Ok(response) => {
                        if let Ok(block) = response.json::<StreamerMessage>().await {
                            if sender.send(block).await.is_err() {
                                break;
                            }
                        }
                    }
                    Err(e) => eprintln!("Error fetching latest block: {:?}", e),
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
