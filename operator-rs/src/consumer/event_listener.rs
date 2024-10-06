use prometheus::{Registry, IntCounter};
use anyhow::Result;
use std::sync::Arc;

pub trait EventListener: Send + Sync {
    fn on_arrival(&self);
    fn on_format_error(&self);
}

const OPERATOR_NAMESPACE: &str = "sffl_operator";
const CONSUMER_SUBSYSTEM: &str = "consumer";

#[derive(Default)]
pub struct SelectiveListener {
    on_arrival_cb: Option<Arc<dyn Fn() + Send + Sync>>,
    on_format_error_cb: Option<Arc<dyn Fn() + Send + Sync>>,
}

impl EventListener for SelectiveListener {
    fn on_arrival(&self) {
        if let Some(cb) = &self.on_arrival_cb {
            cb();
        }
    }

    fn on_format_error(&self) {
        if let Some(cb) = &self.on_format_error_cb {
            cb();
        }
    }
}

impl Clone for SelectiveListener {
    fn clone(&self) -> Self {
        Self {
            on_arrival_cb: self.on_arrival_cb.clone(),
            on_format_error_cb: self.on_format_error_cb.clone(),
        }
    }
}

pub fn make_consumer_metrics(registry: &Registry) -> Result<Arc<dyn EventListener>> {
    let num_blocks_arrived = IntCounter::new(
        "num_of_mq_arrivals",
        "The number of consumed blocks from MQ",
    )?;
    registry.register(Box::new(num_blocks_arrived.clone()))?;

    let num_format_errors = IntCounter::new(
        "num_of_mismatched_blocks",
        "The number of blocks from MQ with invalid format",
    )?;
    registry.register(Box::new(num_format_errors.clone()))?;

    Ok(Arc::new(SelectiveListener {
        on_arrival_cb: Some(Arc::new(move || num_blocks_arrived.inc())),
        on_format_error_cb: Some(Arc::new(move || num_format_errors.inc())),
    }))
}