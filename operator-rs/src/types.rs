use std::collections::HashMap;
use borsh::BorshDeserialize;
use eigensdk::{crypto_bls::BlsSignature, types::operator::OperatorId};
use serde::{Deserialize, Serialize};
use alloy_rpc_types::Block;
use eigensdk::logging::logger::Logger;
use tracing::{debug, error, info, warn};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct NFFLNodeConfig {
    pub production: bool,
    pub operator_address: String,
    pub operator_state_retriever_address: String,
    pub avs_registry_coordinator_address: String,
    pub token_strategy_addr: String,
    pub eth_rpc_url: String,
    pub eth_ws_url: String,
    pub ecdsa_private_key_store_path: String,
    // Replace bls_private_key_store_path with bls_keypair
    pub bls_keypair: String,
    pub aggregator_server_ip_port_address: String,
    pub register_operator_on_startup: bool,
    pub eigen_metrics_ip_port_address: String,
    pub enable_metrics: bool,
    pub node_api_ip_port_address: String,
    pub enable_node_api: bool,
    pub near_da_indexer_rmq_ip_port_address: String,
    pub near_da_indexer_rollup_ids: Vec<u32>,
    pub rollup_ids_to_rpc_urls: HashMap<u32, String>,
    pub task_response_wait_ms: u32,
}

#[derive(Clone, Debug)]
pub struct SignedStateRootUpdateMessage {
    pub message: StateRootUpdateMessage,
    pub bls_signature: BlsSignature,
    pub operator_id: OperatorId,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct StateRootUpdateMessage {
    pub rollup_id: u32,
    pub block_height: u64,
    pub timestamp: u64,
    pub state_root: [u8; 32],
    pub near_da_transaction_id: [u8; 32],
    pub near_da_commitment: [u8; 32],
}


#[derive(Clone, Debug)]
pub struct BlockData {
    pub rollup_id: u32,
    pub commitment: [u8; 32],
    pub transaction_id: [u8; 32],
    pub block: Block,
}


#[derive(Clone, Debug)]
pub struct SubmitRequest {
    pub blobs: Vec<Blob>,
}

#[derive(Clone, Debug)]
pub struct Blob {
    pub namespace: Namespace,
    pub share_version: u32,
    pub commitment: [u8; 32],
    pub data: Vec<u8>,
}

#[derive(Clone, Debug)]
pub struct Namespace {
    pub version: u8,
    pub id: u32,
}
#[derive(Clone, Debug)]
pub struct PublishPayload {
    pub transaction_id: [u8; 32],
    pub data: Vec<u8>,
}


impl BorshDeserialize for PublishPayload {
    fn deserialize_reader<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        let mut transaction_id = [0u8; 32];
        reader.read_exact(&mut transaction_id)?;
        let mut data = Vec::new();
        reader.read_to_end(&mut data)?;
        Ok(PublishPayload { transaction_id, data })
    }
}

impl BorshDeserialize for SubmitRequest {
    fn deserialize_reader<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        let blobs_len = u32::deserialize_reader(reader)? as usize;
        let mut blobs = Vec::with_capacity(blobs_len);
        for _ in 0..blobs_len {
            blobs.push(Blob::deserialize_reader(reader)?);
        }
        Ok(SubmitRequest { blobs })
    }
}




impl BorshDeserialize for Blob {
    fn deserialize_reader<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        Ok(Blob {
            namespace: Namespace::deserialize_reader(reader)?,
            share_version: u32::deserialize_reader(reader)?,
            commitment: {
                let mut commitment = [0u8; 32];
                reader.read_exact(&mut commitment)?;
                commitment
            },
            data: {
                let data_len = u32::deserialize_reader(reader)? as usize;
                let mut data = vec![0u8; data_len];
                reader.read_exact(&mut data)?;
                data
            },
        })
    }
}

impl BorshDeserialize for Namespace {
    fn deserialize_reader<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {
        Ok(Namespace {
            version: u8::deserialize_reader(reader)?,
            id: u32::deserialize_reader(reader)?,
        })
    }
}

// Add this struct and implementation at the end of the file
#[derive(Debug)]
pub struct DefaultLogger;

impl Logger for DefaultLogger {
    fn info(&self, target: &str, message: &str) {
        info!(target = target, "{}", message);
    }

    fn warn(&self, target: &str, message: &str) {
        warn!(target = target, "{}", message);
    }

    fn error(&self, target: &str, message: &str) {
        error!(target = target, "{}", message);
    }

    fn debug(&self, target: &str, message: &str) {
        debug!(target = target, "{}", message);
    }

    fn fatal(&self, target: &str, message: &str) {
        error!(target = target, "FATAL: {}", message);
    }

    fn log(&self, target: &str, message: &str) {
        info!(target: "default_logger", "{}: {}", target, message);
    }
}

// Function to create a new logger instance
pub fn create_default_logger() -> Box<dyn Logger + Send + Sync> {
    Box::new(DefaultLogger)
}
