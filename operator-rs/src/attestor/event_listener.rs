use prometheus::{Registry, Counter, CounterVec, Gauge, GaugeVec, opts};
use anyhow::Result;

pub trait EventListener: Send + Sync {
    fn on_missed_mq_block(&self, rollup_id: u32);
    fn on_block_mismatch(&self, rollup_id: u32);
    fn on_block_received(&self, rollup_id: u32);
    fn observe_last_block_received(&self, rollup_id: u32, block_number: u64);
    fn observe_last_block_received_timestamp(&self, rollup_id: u32, timestamp: u64);
    fn observe_initialization_initial_block_number(&self, rollup_id: u32, block_number: u64);
}

pub struct SelectiveEventListener {
    on_missed_mq_block_cb: Option<Box<dyn Fn(u32) + Send + Sync>>,
    on_block_mismatch_cb: Option<Box<dyn Fn(u32) + Send + Sync>>,
    on_block_received_cb: Option<Box<dyn Fn(u32) + Send + Sync>>,
    observe_last_block_received_cb: Option<Box<dyn Fn(u32, u64) + Send + Sync>>,
    observe_last_block_received_timestamp_cb: Option<Box<dyn Fn(u32, u64) + Send + Sync>>,
    observe_initialization_initial_block_number_cb: Option<Box<dyn Fn(u32, u64) + Send + Sync>>,
}

impl Default for SelectiveEventListener {
    fn default() -> Self {
        Self {
            on_missed_mq_block_cb: None,
            on_block_mismatch_cb: None,
            on_block_received_cb: None,
            observe_last_block_received_cb: None,
            observe_last_block_received_timestamp_cb: None,
            observe_initialization_initial_block_number_cb: None,
        }
    }
}

impl EventListener for SelectiveEventListener {
    fn on_missed_mq_block(&self, rollup_id: u32) {
        if let Some(cb) = &self.on_missed_mq_block_cb {
            cb(rollup_id);
        }
    }

    fn on_block_mismatch(&self, rollup_id: u32) {
        if let Some(cb) = &self.on_block_mismatch_cb {
            cb(rollup_id);
        }
    }

    fn on_block_received(&self, rollup_id: u32) {
        if let Some(cb) = &self.on_block_received_cb {
            cb(rollup_id);
        }
    }

    fn observe_last_block_received(&self, rollup_id: u32, block_number: u64) {
        if let Some(cb) = &self.observe_last_block_received_cb {
            cb(rollup_id, block_number);
        }
    }

    fn observe_last_block_received_timestamp(&self, rollup_id: u32, timestamp: u64) {
        if let Some(cb) = &self.observe_last_block_received_timestamp_cb {
            cb(rollup_id, timestamp);
        }
    }

    fn observe_initialization_initial_block_number(&self, rollup_id: u32, block_number: u64) {
        if let Some(cb) = &self.observe_initialization_initial_block_number_cb {
            cb(rollup_id, block_number);
        }
    }
}

const OPERATOR_NAMESPACE: &str = "sffl_operator";
const ATTESTOR_SUBSYSTEM: &str = "attestor";

pub fn make_attestor_metrics(registry: &Registry) -> Result<impl EventListener> {
    let num_missed_mq_blocks = CounterVec::new(
        opts!("num_of_missed_mq_blocks", "The number of late blocks from MQ")
            .namespace(OPERATOR_NAMESPACE)
            .subsystem(ATTESTOR_SUBSYSTEM),
        &["rollup_id"],
    )?;
    registry.register(Box::new(num_missed_mq_blocks.clone()))?;

    let num_blocks_mismatched = CounterVec::new(
        opts!("num_of_mismatched_blocks", "The number of blocks from MQ mismatched with RPC ones")
            .namespace(OPERATOR_NAMESPACE)
            .subsystem(ATTESTOR_SUBSYSTEM),
        &["rollup_id"],
    )?;
    registry.register(Box::new(num_blocks_mismatched.clone()))?;

    let num_blocks_received = CounterVec::new(
        opts!("num_of_received_blocks", "The number of blocks received from RPC")
            .namespace(OPERATOR_NAMESPACE)
            .subsystem(ATTESTOR_SUBSYSTEM),
        &["rollup_id"],
    )?;
    registry.register(Box::new(num_blocks_received.clone()))?;

    let last_block_received = GaugeVec::new(
        opts!("last_block_received", "Last block received per rollup ID")
            .namespace(OPERATOR_NAMESPACE),
        &["rollup_id"],
    )?;
    registry.register(Box::new(last_block_received.clone()))?;

    let last_block_received_timestamp = GaugeVec::new(
        opts!("last_block_received_timestamp", "Timestamp of last block received per rollup ID")
            .namespace(OPERATOR_NAMESPACE),
        &["rollup_id"],
    )?;
    registry.register(Box::new(last_block_received_timestamp.clone()))?;

    let initialization_initial_block_number = GaugeVec::new(
        opts!("initialization_initial_block_number", "Initialization initial block number per rollup ID")
            .namespace(OPERATOR_NAMESPACE),
        &["rollup_id"],
    )?;
    registry.register(Box::new(initialization_initial_block_number.clone()))?;

    Ok(SelectiveEventListener {
        on_missed_mq_block_cb: Some(Box::new(move |rollup_id| {
            num_missed_mq_blocks.with_label_values(&[&rollup_id.to_string()]).inc();
        })),
        on_block_mismatch_cb: Some(Box::new(move |rollup_id| {
            num_blocks_mismatched.with_label_values(&[&rollup_id.to_string()]).inc();
        })),
        on_block_received_cb: Some(Box::new(move |rollup_id| {
            num_blocks_received.with_label_values(&[&rollup_id.to_string()]).inc();
        })),
        observe_last_block_received_cb: Some(Box::new(move |rollup_id, block_number| {
            last_block_received.with_label_values(&[&rollup_id.to_string()]).set(block_number as f64);
        })),
        observe_last_block_received_timestamp_cb: Some(Box::new(move |rollup_id, timestamp| {
            last_block_received_timestamp.with_label_values(&[&rollup_id.to_string()]).set(timestamp as f64);
        })),
        observe_initialization_initial_block_number_cb: Some(Box::new(move |rollup_id, block_number| {
            initialization_initial_block_number.with_label_values(&[&rollup_id.to_string()]).set(block_number as f64);
        })),
    })
}