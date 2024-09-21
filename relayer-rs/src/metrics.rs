use prometheus::{Counter, Encoder, Histogram, HistogramOpts, IntCounter, Registry};
use std::sync::Arc;
use tokio::time::Duration;
use warp::Filter;
use anyhow::Result;

const RELAYER_NAMESPACE: &str = "sffl_relayer";

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
    pub fn new(registry: &Registry) -> Result<Arc<Self>> {
        let num_blocks_received = IntCounter::new(
            format!("{}_num_blocks_received", RELAYER_NAMESPACE),
            "The number of blocks received from rollup",
        )?;
        registry.register(Box::new(num_blocks_received.clone()))?;

        let num_da_submissions_failed = IntCounter::new(
            format!("{}_num_da_submissions_failed", RELAYER_NAMESPACE),
            "The number of failed da submissions",
        )?;
        registry.register(Box::new(num_da_submissions_failed.clone()))?;

        let latency_buckets = vec![
            1.0, 25.0, 50.0, 75.0, 100.0, 250.0, 500.0, 1000.0, 2000.0, 3000.0, 4000.0, 5000.0,
            10000.0, f64::INFINITY,
        ];

        let submission_duration_ms = Histogram::with_opts(
            HistogramOpts::new(
                format!("{}_submission_duration_ms", RELAYER_NAMESPACE),
                "Duration of successful DA submissions",
            )
            .buckets(latency_buckets),
        )?;
        registry.register(Box::new(submission_duration_ms.clone()))?;

        let retries_histogram = Histogram::with_opts(
            HistogramOpts::new(
                format!("{}_retries_histogram", RELAYER_NAMESPACE),
                "Histogram of retry counts",
            )
            .buckets(prometheus::linear_buckets(0.0, 1.0, 3).unwrap()),
        )?;
        registry.register(Box::new(retries_histogram.clone()))?;

        let num_of_invalid_nonces = IntCounter::new(
            format!("{}_num_of_invalid_nonces", RELAYER_NAMESPACE),
            "Number of InvalidNonce error",
        )?;
        registry.register(Box::new(num_of_invalid_nonces.clone()))?;

        let num_of_expired_txs = IntCounter::new(
            format!("{}_num_of_expired_txs", RELAYER_NAMESPACE),
            "Number of Expired transactions",
        )?;
        registry.register(Box::new(num_of_expired_txs.clone()))?;

        let num_of_timeout_txs = IntCounter::new(
            format!("{}_num_of_timeout_txs", RELAYER_NAMESPACE),
            "Number of Timeout transactions",
        )?;
        registry.register(Box::new(num_of_timeout_txs.clone()))?;

        Ok(Arc::new(Self {
            num_blocks_received,
            num_da_submissions_failed,
            submission_duration_ms,
            retries_histogram,
            num_of_invalid_nonces,
            num_of_expired_txs,
            num_of_timeout_txs,
        }))
    }
}

pub async fn start_metrics_server(metrics_addr: String, registry: Registry) {
    let metrics_route = warp::path!("metrics").and(warp::get()).map(move || {
        let encoder = prometheus::TextEncoder::new();
        let mut buffer = Vec::new();
        encoder.encode(&registry.gather(), &mut buffer).unwrap();
        String::from_utf8(buffer).unwrap()
    });

    let socket_addr: std::net::SocketAddr = metrics_addr.parse().expect("Failed to parse metrics address");
    warp::serve(metrics_route).run(socket_addr).await;
}

pub trait EventListener {
    fn on_block_received(&self);
    fn on_da_submission_failed(&self);
    fn on_da_submitted(&self, duration: Duration);
    fn on_retries_required(&self, retries: i32);
    fn on_invalid_nonce(&self);
    fn on_expired_tx(&self);
    fn on_timeout_tx(&self);
}

impl EventListener for Arc<RelayerMetrics> {
    fn on_block_received(&self) {
        self.num_blocks_received.inc();
    }

    fn on_da_submission_failed(&self) {
        self.num_da_submissions_failed.inc();
    }

    fn on_da_submitted(&self, duration: Duration) {
        self.submission_duration_ms.observe(duration.as_millis() as f64);
    }

    fn on_retries_required(&self, retries: i32) {
        self.retries_histogram.observe(retries as f64);
    }

    fn on_invalid_nonce(&self) {
        self.num_of_invalid_nonces.inc();
    }

    fn on_expired_tx(&self) {
        self.num_of_expired_txs.inc();
    }

    fn on_timeout_tx(&self) {
        self.num_of_timeout_txs.inc();
    }
}