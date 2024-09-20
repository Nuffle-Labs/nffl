use crate::config::RelayerConfig;
use crate::metrics::RelayerMetrics;
use anyhow::Result;
use ethers::prelude::*;
use prometheus::Registry;
use std::sync::Arc;
use std::time::Duration;
use tokio::time;
use tracing::{error, info};

const NAMESPACE_ID: u64 = 1;
const SUBMIT_BLOCK_INTERVAL: Duration = Duration::from_millis(2500);
const SUBMIT_BLOCK_RETRY_TIMEOUT: Duration = Duration::from_secs(2);
const SUBMIT_BLOCK_RETRIES: usize = 3;

pub struct Relayer {
    rpc_client: Provider<Ws>,
    near_client: Arc<near_sdk::NearClient>,
    metrics: Arc<RelayerMetrics>,
}

impl Relayer {
    pub fn new(config: RelayerConfig) -> Result<Self> {
        let rpc_client = Provider::<Ws>::connect(&config.rpc_url)?;
        let near_client = Arc::new(near_sdk::NearClient::new(
            &config.key_path,
            &config.da_account_id,
            &config.network,
            NAMESPACE_ID,
        )?);

        let registry = Registry::new();
        let metrics = RelayerMetrics::new(&registry)?;

        Ok(Self {
            rpc_client,
            near_client,
            metrics,
        })
    }

    pub async fn start(&mut self) -> Result<()> {
        let mut block_stream = self.rpc_client.subscribe_blocks().await?;
        let mut interval = time::interval(SUBMIT_BLOCK_INTERVAL);

        let mut blocks = Vec::new();

        loop {
            tokio::select! {
                Some(block) = block_stream.next() => {
                    info!("Received rollup block header: {:?}", block.number);
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

    async fn handle_blocks(&self, blocks: &[Block<TxHash>]) -> Result<()> {
        info!("Submitting blocks to NEAR: {:?}", blocks.iter().map(|b| b.number).collect::<Vec<_>>());

        let encoded_blocks = rlp::encode_list(blocks);
        self.submit_encoded_blocks(&encoded_blocks).await?;

        Ok(())
    }

    async fn submit_encoded_blocks(&self, encoded_blocks: &[u8]) -> Result<()> {
        let start_time = std::time::Instant::now();

        for i in 0..SUBMIT_BLOCK_RETRIES {
            match self.near_client.force_submit(encoded_blocks).await {
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
        anyhow::bail!("Failed to submit blocks to NEAR after retries")
    }
}