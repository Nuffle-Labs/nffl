use crate::config::RelayerConfig;
use crate::metrics::RelayerMetrics;
use alloy::pubsub::PubSubFrontend;
use anyhow::Result;
use alloy_rpc_types::Block;
use alloy_transport_ws::WsConnect;
use prometheus::Registry;
use std::{path::PathBuf, sync::Arc};
use std::time::Duration;
use tokio::time;
use tracing::{error, info};
use near_da_rpc::*;
use alloy::providers::{Provider, ProviderBuilder, RootProvider};
use futures_util::StreamExt;

const NAMESPACE_ID: u8 = 1;
const SUBMIT_BLOCK_INTERVAL: Duration = Duration::from_millis(2500);
const SUBMIT_BLOCK_RETRY_TIMEOUT: Duration = Duration::from_secs(2);
const SUBMIT_BLOCK_RETRIES: usize = 3;

pub struct Relayer {
    provider: Arc<RootProvider<PubSubFrontend>>,
    near_da_client: near_da_rpc::near::Client,
    metrics: Arc<RelayerMetrics>,
}

impl Relayer {
    pub async fn new(config: &RelayerConfig) -> Result<Self> {
        let ws = WsConnect::new(&config.rpc_url);
        let provider = ProviderBuilder::new().on_ws(ws).await?;

        let config = near_da_rpc::near::config::Config {
            key: near_da_rpc::near::config::KeyType::File(PathBuf::from(config.key_path.clone())),  
            contract: config.da_account_id.clone(),
            network: near_da_rpc::near::config::Network::Custom(config.network.clone()),
            namespace: Some(near_da_primitives::Namespace::new(NAMESPACE_ID, 1)),
            mode: near_da_primitives::Mode::Standard,
        };
        let near_da_client = near_da_rpc::near::Client::new(&config);
        let registry = Registry::new();
        let metrics = RelayerMetrics::new(&registry)?;

        Ok(Self {
            provider: Arc::new(provider),
            near_da_client,
            metrics,
        })
    }

    pub async fn start(&mut self) -> Result<()> {
        let mut interval = time::interval(SUBMIT_BLOCK_INTERVAL);

        let mut blocks = Vec::new();
         // Subscribe to blocks.
        let subscription = self.provider.subscribe_blocks().await?;
        let mut stream = subscription.into_stream();

        loop {
            tokio::select! {
                Some(block) = stream.next() => {
                    info!("Received rollup block header: {:?}", block.header.number);
                    self.metrics.num_blocks_received.inc();
                    blocks.push(block);
                }
                _ = interval.tick() => {
                    if !blocks.is_empty() {
                        if let Err(e) = self.handle_blocks(&blocks).await {
                            error!("Error handling blocks: {:?}", e);
                        }
                        blocks.clear();
                    }
                }
            }
        }
    }

    async fn handle_blocks(&self, blocks: &[Block]) -> Result<()> {
        info!("Submitting blocks to NEAR: {:?}", blocks.iter().map(|b| b.header.number).collect::<Vec<_>>());
        let serialized_blocks = serde_json::to_vec(blocks)?;
        let encoded_blocks = alloy_rlp::encode(&serialized_blocks);
        self.submit_encoded_blocks(encoded_blocks).await?;

        Ok(())
    }

    async fn submit_encoded_blocks(&self, encoded_blocks: Vec<u8>) -> Result<()> {
        let start_time = std::time::Instant::now();
        for i in 0..SUBMIT_BLOCK_RETRIES {
            match self.near_da_client.submit(Blob::from(encoded_blocks.to_vec())).await {
                Ok(out) => {
                    self.metrics.submission_duration_ms.observe(start_time.elapsed().as_millis() as f64);
                    self.metrics.retries_histogram.observe(i as f64);
                    info!("Blocks submitted successfully: {:?}", out);
                    return Ok(());
                }
                Err(e) => {
                    error!("Error submitting blocks to NEAR, retrying: {:?}", e);
                    if e.to_string().contains("InvalidNonce") {
                        self.metrics.num_of_invalid_nonces.inc();
                    } else if e.to_string().contains("Expired") {
                        self.metrics.num_of_expired_txs.inc();
                    } else if e.to_string().contains("Timeout") {
                        self.metrics.num_of_timeout_txs.inc();
                    }
                    time::sleep(SUBMIT_BLOCK_RETRY_TIMEOUT).await;
                }
            }
        }

        self.metrics.num_da_submissions_failed.inc();
        Err(anyhow::anyhow!("Failed to submit blocks to NEAR after retries"))
    }
}