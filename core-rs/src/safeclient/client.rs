use alloy::eips::BlockNumberOrTag;
use alloy::primitives::{B256, U256};
use alloy::providers::{Provider, ProviderBuilder, RootProvider, WsConnect};
use alloy::pubsub::PubSubFrontend;
use alloy::rpc::types::{Block, Filter, Log, Transaction, TransactionReceipt};
use anyhow::Result;
use async_trait::async_trait;
use std::sync::Arc;
use std::time::Duration;
use tokio::sync::broadcast;
use futures_util::StreamExt;
use log::{error, warn};


#[async_trait]
pub trait SafeClient: Send + Sync {
    async fn block_number(&self) -> Result<U256>;
    async fn get_block(&self, block_number: BlockNumberOrTag) -> Result<Option<Block>>;
    async fn get_transaction(&self, tx_hash: B256) -> Result<Option<Transaction>>;
    async fn get_transaction_receipt(&self, tx_hash: B256) -> Result<Option<TransactionReceipt>>;
    async fn get_logs(&self, filter: Filter) -> Result<Vec<Log>>;
    async fn subscribe_logs(&self, filter: Filter) -> Result<broadcast::Receiver<Log>>;
    async fn subscribe_new_heads(&self) -> Result<broadcast::Receiver<Block>>;
    fn close(&self);
}

pub struct SafeEthClient {
    provider: Arc<RootProvider<PubSubFrontend>>,
    log_resub_interval: Duration,
    header_timeout: Duration,
    block_chunk_size: u64,
    block_max_range: u64,
    close_sender: broadcast::Sender<()>,
}

pub struct SafeEthClientOptions {
    pub log_resub_interval: Duration,
    pub header_timeout: Duration,
    pub block_chunk_size: u64,
    pub block_max_range: u64,
}

impl Default for SafeEthClientOptions {
    fn default() -> Self {
        Self {
            log_resub_interval: Duration::from_secs(300),
            header_timeout: Duration::from_secs(30),
            block_chunk_size: 100,
            block_max_range: 100,
        }
    }
}

impl SafeEthClient {
    pub async fn new(ws_url: &str, options: SafeEthClientOptions) -> Result<Self> {
        let ws = WsConnect::new(ws_url);
        let provider = ProviderBuilder::new().on_ws(ws).await?;
        let (close_sender, _) = broadcast::channel(1);

        Ok(Self {
            provider: Arc::new(provider),
            log_resub_interval: options.log_resub_interval,
            header_timeout: options.header_timeout,
            block_chunk_size: options.block_chunk_size,
            block_max_range: options.block_max_range,
            close_sender,
        })
    }
}

#[async_trait]
impl SafeClient for SafeEthClient {
    async fn block_number(&self) -> Result<U256> {
        Ok(U256::from(self.provider.get_block_number().await?))
    }

    async fn get_block(&self, block_number: BlockNumberOrTag) -> Result<Option<Block>> {
        Ok(self.provider.get_block(alloy::eips::BlockId::Number(block_number), alloy::rpc::types::BlockTransactionsKind::Hashes).await?)
    }

    async fn get_transaction(&self, tx_hash: B256) -> Result<Option<Transaction>> {
        Ok(self.provider.get_transaction_by_hash(tx_hash).await?)
    }

    async fn get_transaction_receipt(&self, tx_hash: B256) -> Result<Option<TransactionReceipt>> {
        Ok(self.provider.get_transaction_receipt(tx_hash).await?)
    }

    async fn get_logs(&self, filter: Filter) -> Result<Vec<Log>> {
        Ok(self.provider.get_logs(&filter).await?)
    }

    async fn subscribe_logs(&self, filter: Filter) -> Result<broadcast::Receiver<Log>> {
        let (tx, rx) = broadcast::channel(100);
        let log_resub_interval = self.log_resub_interval;
        let subscription = self.provider.subscribe_logs(&filter).await?;
        let mut stream = subscription.into_stream();

        tokio::spawn(async move {
            loop {
                tokio::select! {
                    Some(log) = stream.next() => {
                        if tx.send(log).is_err() {
                            error!("Error sending log: channel closed");
                            break;
                        }
                    }
                    _ = tokio::time::sleep(log_resub_interval) => {
                        warn!("Timeout waiting for new log");
                        break;
                    }
                    else => {
                        error!("Log stream ended unexpectedly");
                        break;
                    }
                }
            }
        });

        Ok(rx)
    }

    async fn subscribe_new_heads(&self) -> Result<broadcast::Receiver<Block>> {
        let (tx, rx) = broadcast::channel(100);
        let header_timeout = self.header_timeout;
        let subscription = self.provider.subscribe_blocks().await?;
        let mut stream = subscription.into_stream();
    
        tokio::spawn(async move {
            loop {
                tokio::select! {
                    Some(block) = stream.next() => {
                        if tx.send(block).is_err() {
                            error!("Error sending block: channel closed");
                            break;
                        }
                    }
                    _ = tokio::time::sleep(header_timeout) => {
                        warn!("Timeout waiting for new block");
                        break;
                    }
                    else => {
                        error!("Block stream ended unexpectedly");
                        break;
                    }
                }
            }
        });
    
        Ok(rx)
    }

    fn close(&self) {
        let _ = self.close_sender.send(());
    }
}
