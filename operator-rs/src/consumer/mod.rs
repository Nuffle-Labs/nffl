pub mod config;
mod event_listener;
mod queues_listener;

use std::sync::Arc;
use tokio::sync::broadcast;
use tokio::time::{Duration, sleep};
use anyhow::{Result, anyhow};
use tracing::{info, warn, error};
use lapin::{Connection, ConnectionProperties};
use prometheus::Registry;
use crate::types::BlockData;

use self::config::ConsumerConfig;
use self::event_listener::EventListener;
use self::queues_listener::QueuesListener;

const RECONNECT_DELAY: Duration = Duration::from_secs(5);
const RECHANNEL_DELAY: Duration = Duration::from_secs(2);
const EXCHANGE_NAME: &str = "rollup_exchange";

pub struct Consumer {
    received_blocks_tx: Arc<broadcast::Sender<BlockData>>,
    queues_listener: Option<Arc<QueuesListener>>,
    config: ConsumerConfig,
    event_listener: Arc<dyn EventListener>,
    connection: Option<Connection>,
}

impl Consumer {
    pub fn new(config: ConsumerConfig) -> Self {
        let (received_blocks_tx, _) = broadcast::channel(100);
        let received_blocks_tx = Arc::new(received_blocks_tx);
        Self {
            received_blocks_tx,
            queues_listener: None,
            config,
            event_listener: Arc::new(event_listener::SelectiveListener::default()),
            connection: None,
        }
    }

    pub fn enable_metrics(&mut self, registry: &Registry) -> Result<()> {
        self.event_listener = event_listener::make_consumer_metrics(registry)?;
        Ok(())
    }

    pub async fn start(&mut self, addr: &str) -> Result<()> {
        loop {
            match self.connect(addr).await {
                Ok(conn) => {
                    self.connection = Some(conn);
                    if let Err(e) = self.setup_channel().await {
                        error!("Failed to setup channel: {:?}", e);
                        sleep(RECHANNEL_DELAY).await;
                        continue;
                    }
                    info!("Connected and channel set up");
                }
                Err(e) => {
                    error!("Failed to connect: {:?}", e);
                    sleep(RECONNECT_DELAY).await;
                    continue;
                }
            }

            if let Some(conn) = &self.connection {
                match conn.status().connected() {
                    true => {
                        // Connection is still active, wait for a while before checking again
                        sleep(Duration::from_secs(30)).await;
                    }
                    false => {
                        warn!("Connection lost, attempting to reconnect");
                        self.connection = None;
                        sleep(RECONNECT_DELAY).await;
                    }
                }
            } else {
                warn!("No active connection, attempting to reconnect");
                sleep(RECONNECT_DELAY).await;
            }
        }
    }

    async fn connect(&self, addr: &str) -> Result<Connection> {
        Connection::connect(
            addr,
            ConnectionProperties::default(),
        ).await.map_err(|e| anyhow::anyhow!("Failed to connect: {:?}", e))
    }

    async fn setup_channel(&mut self) -> Result<()> {
        let conn = self.connection.as_ref().ok_or_else(|| anyhow!("No active connection"))?;
        let channel = conn.create_channel().await?;

        let mut queues_listener = QueuesListener::new(Arc::clone(&self.received_blocks_tx), Arc::clone(&self.event_listener));

        for &rollup_id in &self.config.rollup_ids {
            let queue_name = config::get_queue_name(rollup_id, &self.config.id);
            let queue = channel.queue_declare(&queue_name, Default::default(), Default::default()).await?;
            
            if queue.message_count() > 0 {
                info!("Queue '{}' declared with {} messages", queue_name, queue.message_count());
            }

            channel.queue_bind(
                &queue_name,
                EXCHANGE_NAME,
                &config::get_routing_key(rollup_id),
                Default::default(),
                Default::default(),
            ).await?;

            let consumer = channel.basic_consume(
                &queue_name,
                &config::get_consumer_tag(rollup_id),
                Default::default(),
                Default::default(),
            ).await?;

            queues_listener.add(rollup_id, consumer).await?;
        }

        self.queues_listener = Some(Arc::new(queues_listener));
        Ok(())
    }

    pub async fn close(&mut self) -> Result<()> {
        if let Some(conn) = self.connection.take() {
            conn.close(0, "Consumer closed").await?;
        }
        Ok(())
    }

    pub fn get_block_stream(&self) -> broadcast::Receiver<BlockData> {
        self.received_blocks_tx.subscribe()
    }
}