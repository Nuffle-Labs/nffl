use alloy::{primitives::B256, rpc::types::Log};
use sha2::{Digest, Sha256};

pub fn hash_log(log: &Log) -> B256 {
    let mut hasher = Sha256::new();
    
    // Hash the log fields
    hasher.update(log.address().as_slice());
    for topic in log.topics() {
        hasher.update(topic.as_slice());
    }
    hasher.update(log.data().data.clone());
    
    // Hash additional block and tx info
    if let Some(block_hash) = log.block_hash {
        hasher.update(block_hash.as_slice());
    }
    if let Some(transaction_hash) = log.transaction_hash {
        hasher.update(transaction_hash.as_slice());
    }
    if let Some(log_index) = log.log_index {
        hasher.update(&log_index.to_be_bytes());
    }

    B256::from_slice(&hasher.finalize())
}
