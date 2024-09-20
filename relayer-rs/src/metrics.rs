use prometheus::{Counter, Histogram, IntCounter, Registry};
use std::sync::Arc;

pub struct RelayerMetrics {
    pub num_blocks_received: IntCounter,
    pub num_da_submissions_failed: IntCounter,
    pub submission_duration_ms: Histogram,
    pub retries_histogram: Histogram,
    pub num_of_invalid_nonces: IntCounter,
    pub num_of_expired_txs: IntCounter,
    pub num_of_timeout_txs: IntCounter,
}

impl RelayerMetrics {
    pub fn new(registry: &Registry) -> anyhow::Result<Arc<Self>> {
        let num_blocks_received = IntCounter::new(
            "sffl_relayer_num_blocks_received",
            "The number of blocks received from rollup",
        )?;
        registry.register(Box::new(num_blocks_received.clone()))?;

        // ... Initialize other metrics similarly ...

        Ok(Arc::new(Self {
            num_blocks_received,
            num_da_submissions_failed: IntCounter::new("sffl_relayer_num_da_submissions_failed", "The number of failed da submissions")?,
            submission_duration_ms: Histogram::new("sffl_relayer_submission_duration_ms", "Duration of successful DA submissions")?,
            retries_histogram: Histogram::new("sffl_relayer_retries_histogram", "Histogram of retry counts")?,
            num_of_invalid_nonces: IntCounter::new("sffl_relayer_num_of_invalid_nonces", "Number of InvalidNonce error")?,
            num_of_expired_txs: IntCounter::new("sffl_relayer_num_of_expired_txs", "Number of Expired transactions")?,
            num_of_timeout_txs: IntCounter::new("sffl_relayer_num_of_timeout_txs", "Number of Timeout transactions")?,
        }))
    }
}