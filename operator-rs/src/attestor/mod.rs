mod notifier;
mod event_listener;

use std::collections::HashMap;
use std::sync::Arc;
use alloy_rlp::Header;
use eigensdk::logging::logger::Logger;
use tokio::sync::mpsc;
use tokio::time::{Duration, sleep};
use anyhow::{Result, anyhow};
use tracing::{info, warn, error};
use prometheus::Registry;

use serde::{Serialize, Deserialize};
use crate::types::NodeConfig;
use crate::types::messages::SignedStateRootUpdateMessage;
use self::notifier::Notifier;
use self::event_listener::{EventListener, SelectiveEventListener};
use eigensdk::crypto_bls::BlsKeyPair;
// Constants
const MQ_WAIT_TIMEOUT: Duration = Duration::from_secs(30);
const MQ_REBROADCAST_TIMEOUT: Duration = Duration::from_secs(15);

// Mock types (to be replaced with actual implementations)
type SafeClient = Arc<dyn SafeClientTrait>;
trait SafeClientTrait: Send + Sync {
    fn subscribe_new_head(&self) -> mpsc::Receiver<Header>;
    fn block_number(&self) -> u64;
    fn close(&self);
}

// TODO: Replace wth actual types from eigensdk-rs
type OperatorId = [u8; 32];

pub struct Attestor {
    signed_root_tx: mpsc::Sender<SignedStateRootUpdateMessage>,
    rollup_ids_to_urls: HashMap<u32, String>,
    clients: HashMap<u32, SafeClient>,
    rpc_calls_collectors: HashMap<u32, ()>, // Replace with actual RPC calls collector
    notifier: Notifier,
    consumer: Consumer,
    config: NodeConfig,
    bls_keypair: BlsKeyPair,
    operator_id: OperatorId,
    logger: Box<dyn eigensdk::logging::logger::Logger>,
    listener: Box<dyn EventListener>,
    registry: Registry,
}

impl Attestor {
    pub fn new(config: &NodeConfig, bls_keypair: BlsKeyPair, operator_id: OperatorId, registry: Registry, logger: Box<dyn Logger>) -> Result<Self> {
        let consumer = Consumer::new(ConsumerConfig {
            rollup_ids: config.near_da_indexer_rollup_ids.clone(),
            id: hex::encode(&operator_id),
        });

        let mut clients = HashMap::new();
        let mut rpc_calls_collectors = HashMap::new();

        for (rollup_id, url) in &config.rollup_ids_to_rpc_urls {
            let client = create_safe_client(url)?;
            clients.insert(*rollup_id, client);

            if config.enable_metrics {
                // Create and add RPC calls collector (mock for now)
                rpc_calls_collectors.insert(*rollup_id, ());
            }
        }

        let (signed_root_tx, _) = mpsc::channel(100);

        Ok(Self {
            signed_root_tx,
            rollup_ids_to_urls: config.rollup_ids_to_rpc_urls.clone(),
            clients,
            rpc_calls_collectors,
            notifier: Notifier::new(),
            consumer,
            config: config.clone(),
            bls_keypair,
            operator_id,
            logger,
            listener: Box::new(SelectiveEventListener::default()),
            registry,
        })
    }

    pub fn enable_metrics(&mut self, registry: &Registry) -> Result<()> {
        let listener = event_listener::make_attestor_metrics(registry)?;
        self.listener = Box::new(listener);
        self.consumer.enable_metrics(registry)?;
        Ok(())
    }

    pub async fn start(&self) -> Result<()> {
        self.consumer.start(&self.config.near_da_indexer_rmq_ip_port_address).await?;

        let mut subscriptions = HashMap::new();
        let mut headers_rxs = HashMap::new();

        for (rollup_id, client) in &self.clients {
            let headers_rx = client.subscribe_new_head();
            let block_number = client.block_number();

            self.listener.observe_initialization_initial_block_number(*rollup_id, block_number);

            subscriptions.insert(*rollup_id, ());
            headers_rxs.insert(*rollup_id, headers_rx);
        }

        let mq_task = tokio::spawn(self.process_mq_blocks());

        let mut rollup_tasks = Vec::new();
        for (rollup_id, headers_rx) in headers_rxs {
            let task = tokio::spawn(self.process_rollup_headers(rollup_id, headers_rx));
            rollup_tasks.push(task);
        }

        // You might want to store these JoinHandles somewhere if you need to cancel them later
        Ok(())
    }

    async fn process_mq_blocks(&self) {
        let mut mq_block_rx = self.consumer.get_block_stream();

        while let Some(mq_block) = mq_block_rx.recv().await {
            if ctx.is_done() {
                return;
            }

            info!(self.logger, "Notifying"; "rollupId" => mq_block.rollup_id, "height" => mq_block.block.header().number);
            if let Err(e) = self.notifier.notify(mq_block.rollup_id, mq_block.clone()) {
                error!(self.logger, "Notifier error"; "err" => ?e);
            }

            let notifier = self.notifier.clone();
            let logger = self.logger.clone();
            tokio::spawn(async move {
                tokio::select! {
                    _ = sleep(MQ_REBROADCAST_TIMEOUT) => {
                        info!(logger, "Renotifying"; "rollupId" => mq_block.rollup_id, "height" => mq_block.block.header().number);
                        if let Err(e) = notifier.notify(mq_block.rollup_id, mq_block) {
                            error!(logger, "Error while renotifying"; "err" => ?e);
                        }
                    }
                    _ = ctx.done() => {}
                }
            });
        }
    }

    async fn process_rollup_headers(&self, rollup_id: u32, mut headers_rx: mpsc::Receiver<Header>) {
        while let Some(header) = headers_rx.recv().await {
            self.process_header(rollup_id, header).await;
        }
    }

    async fn process_header(&self, rollup_id: u32, rollup_header: Header) {
        info!(self.logger, "Processing header"; "rollupId" => rollup_id, "height" => get_header_number(&rollup_header));

        self.listener.observe_last_block_received(rollup_id, get_header_number(&rollup_header));
        self.listener.observe_last_block_received_timestamp(rollup_id, get_header_timestamp(&rollup_header));
        self.listener.on_block_received(rollup_id);

        let predicate = |mq_block: &BlockData| {
            if mq_block.rollup_id != rollup_id {
                warn!(self.logger, "Subscriber rollupId mismatch"; "expected" => rollup_id, "actual" => mq_block.rollup_id);
                return false;
            }

            if get_header_number(&rollup_header) != get_block_number(&mq_block.block) {
                return false;
            }

            if get_block_root(&mq_block.block) != get_header_root(&rollup_header) {
                warn!(self.logger, "StateRoot from MQ doesn't match one from Node");
                self.listener.on_block_mismatch(rollup_id);
                return false;
            }

            true
        };

        let (mq_blocks_rx, id) = self.notifier.subscribe(rollup_id, predicate);

        let mut transaction_id = [0u8; 32];
        let mut da_commitment = [0u8; 32];

        tokio::select! {
            _ = sleep(MQ_WAIT_TIMEOUT) => {
                info!(self.logger, "MQ timeout"; "rollupId" => rollup_id, "height" => get_header_number(&rollup_header));
                self.listener.on_missed_mq_block(rollup_id);
            }
            Some(mq_block) = mq_blocks_rx.recv() => {
                info!(self.logger, "MQ block found"; "height" => get_block_number(&mq_block.block), "rollupId" => mq_block.rollup_id);
                transaction_id = mq_block.transaction_id;
                da_commitment = mq_block.commitment;
            }
            _ = ctx.done() => {
                return;
            }
        }

        self.notifier.unsubscribe(rollup_id, id);

        let message = StateRootUpdateMessage {
            rollup_id,
            block_height: get_header_number(&rollup_header),
            timestamp: get_header_timestamp(&rollup_header),
            state_root: get_header_root(&rollup_header),
            near_da_transaction_id: transaction_id,
            near_da_commitment: da_commitment,
        };

        match sign_state_root_update_message(&self.bls_keypair, &message) {
            Ok(signature) => {
                let signed_message = SignedStateRootUpdateMessage {
                    message,
                    bls_signature: signature,
                    operator_id: self.operator_id,
                };
                if let Err(e) = self.signed_root_tx.send(signed_message).await {
                    warn!(self.logger, "Failed to send signed state root update"; "err" => ?e);
                }
            }
            Err(e) => {
                warn!(self.logger, "StateRoot sign failed"; "err" => ?e);
            }
        }
    }

    pub fn get_signed_root_rx(&self) -> mpsc::Receiver<SignedStateRootUpdateMessage> {
        self.signed_root_tx.subscribe()
    }

    pub fn close(&self) -> Result<()> {
        self.consumer.close()?;
        for client in self.clients.values() {
            client.close();
        }
        Ok(())
    }
}

// Helper functions (to be implemented)
fn create_safe_client(_url: &str) -> Result<SafeClient> {
    unimplemented!()
}

fn sign_state_root_update_message(_keypair: &BlsKeyPair, _message: &StateRootUpdateMessage) -> Result<BlsSignature> {
    unimplemented!()
}

fn get_header_number(_header: &Header) -> u64 {
    unimplemented!()
}

fn get_header_timestamp(_header: &Header) -> u64 {
    unimplemented!()
}

fn get_header_root(_header: &Header) -> [u8; 32] {
    unimplemented!()
}

fn get_block_number(_block: &Block) -> u64 {
    unimplemented!()
}

fn get_block_root(_block: &Block) -> [u8; 32] {
    unimplemented!()
}

// Mock types (to be replaced with actual implementations)
struct Consumer;
struct ConsumerConfig {
    rollup_ids: Vec<u32>,
    id: String,
}
struct Context;

#[derive(Clone, Debug, Serialize, Deserialize)]
struct BlockData {
    rollup_id: u32,
    block: Block,
    transaction_id: [u8; 32],
    commitment: [u8; 32],
}

#[derive(Clone, Debug, Serialize, Deserialize)]
struct StateRootUpdateMessage {
    rollup_id: u32,
    block_height: u64,
    timestamp: u64,
    state_root: [u8; 32],
    near_da_transaction_id: [u8; 32],
    near_da_commitment: [u8; 32],
}
type BlsSignature = ();

impl Consumer {
    fn new(_config: ConsumerConfig) -> Self {
        unimplemented!()
    }
    async fn start(&self, _address: &str) -> Result<()> {
        unimplemented!()
    }
    fn get_block_stream(&self) -> mpsc::Receiver<BlockData> {
        unimplemented!()
    }
    fn enable_metrics(&self, _registry: &Registry) -> Result<()> {
        unimplemented!()
    }
    fn close(&self) -> Result<()> {
        unimplemented!()
    }
}

impl Context {
    fn is_done(&self) -> bool {
        unimplemented!()
    }
    fn done(&self) -> impl std::future::Future<Output = ()> {
        std::future::ready(())
    }
}