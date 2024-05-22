use prometheus::core::GenericGauge;
use prometheus::{
    core::{AtomicF64, GenericCounter},
    Counter, Gauge, Histogram, HistogramOpts, Opts, Registry,
};

use crate::errors::Result;

const INDEXER_NAMESPACE: &str = "sffl_indexer";
const CANDIDATES_SUBSYSTEM: &str = "candidates_validator";
const LISTENER_SUBSYSTEM: &str = "block_listener";
const PUBLISHER_SUBSYSTEM: &str = "rabbit_publisher";

#[derive(Clone)]
pub struct CandidatesListener {
    pub num_successful: GenericCounter<AtomicF64>,
    pub num_failed: GenericCounter<AtomicF64>,
}

#[derive(Clone)]
pub struct BlockEventListener {
    pub num_candidates: GenericCounter<AtomicF64>,
    pub current_queued_candidates: GenericGauge<AtomicF64>,
}

#[derive(Clone)]
pub struct PublisherListener {
    pub num_published_blocks: GenericCounter<AtomicF64>,
    pub num_failed_publishes: GenericCounter<AtomicF64>,
    pub publish_duration_histogram: Histogram,
}

pub trait Metricable {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()>;
}

pub(crate) fn make_candidates_validator_metrics(registry: Registry) -> Result<CandidatesListener> {
    let opts = Opts::new("num_of_successful_candidates", "Number of successful candidates")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(CANDIDATES_SUBSYSTEM);
    let num_successful = Counter::with_opts(opts)?;

    registry.register(Box::new(num_successful.clone()))?;

    let opts = Opts::new("num_of_failed_candidates", "Number of rejected candidates")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(CANDIDATES_SUBSYSTEM);
    let num_failed = Counter::with_opts(opts)?;

    registry.register(Box::new(num_failed.clone()))?;

    Ok(CandidatesListener {
        num_successful,
        num_failed,
    })
}

pub(crate) fn make_block_listener_metrics(registry: Registry) -> Result<BlockEventListener> {
    let opts = Opts::new("num_of_candidates", "Number of candidates indexed")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(LISTENER_SUBSYSTEM);
    let num_candidates = Counter::with_opts(opts)?;
    registry.register(Box::new(num_candidates.clone()))?;

    let opts = Opts::new("current_queued_candidates", "Current number of queued messages")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(LISTENER_SUBSYSTEM);
    let current_queued_candidates = Gauge::with_opts(opts)?;
    registry.register(Box::new(current_queued_candidates.clone()))?;

    Ok(BlockEventListener {
        num_candidates,
        current_queued_candidates,
    })
}

pub(crate) fn make_publisher_metrics(registry: Registry) -> Result<PublisherListener> {
    let opts = Opts::new("num_of_published_blocks", "Number of published blocks to MQ")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(PUBLISHER_SUBSYSTEM);
    let num_published_blocks = Counter::with_opts(opts)?;

    registry.register(Box::new(num_published_blocks.clone()))?;

    let opts = Opts::new("num_failed_published", "Number of failed publishes to MQ")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(PUBLISHER_SUBSYSTEM);
    let num_failed_publishes = Counter::with_opts(opts)?;

    registry.register(Box::new(num_failed_publishes.clone()))?;

    let publish_duration_opts = HistogramOpts::new("publish_duration_seconds", "Time spent in publishing messages")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(PUBLISHER_SUBSYSTEM)
        .buckets(vec![5., 10., 25., 50., 100., 250., 500., 1000., 2500., 5000., 10000.]); // ms
    let publish_duration_histogram = Histogram::with_opts(publish_duration_opts)?;

    registry.register(Box::new(publish_duration_histogram.clone()))?;

    Ok(PublisherListener {
        num_published_blocks,
        num_failed_publishes,
        publish_duration_histogram,
    })
}
