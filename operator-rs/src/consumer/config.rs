use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct ConsumerConfig {
    pub rollup_ids: Vec<u32>,
    pub id: String,
}

pub fn get_queue_name(rollup_id: u32, id: &str) -> String {
    format!("rollup{}-{}", rollup_id, id)
}

pub fn get_routing_key(rollup_id: u32) -> String {
    format!("rollup{}", rollup_id)
}

pub fn get_consumer_tag(rollup_id: u32) -> String {
    format!("operator{}", rollup_id)
}