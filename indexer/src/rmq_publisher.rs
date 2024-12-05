use borsh::{BorshDeserialize, BorshSerialize};
use deadpool::managed::PoolError;
use deadpool_lapin::{Manager, Pool};
use lapin::{
    options::{BasicPublishOptions, ExchangeDeclareOptions},
    types::FieldTable,
    BasicProperties, ConnectionProperties, ExchangeKind,
};
use near_indexer::near_primitives::hash::CryptoHash;
use prometheus::Registry;
use std::time::Duration;
use tokio::{sync::mpsc, task::JoinHandle};
use tracing::{error, info};

use crate::{
    errors::{Error, Result},
    metrics::{make_publisher_metrics, Metricable, PublisherListener},
};

const PUBLISHER: &str = "publisher";
const EXCHANGE_NAME: &str = "rollup_exchange";
const DEFAULT_ROUTING_KEY: &str = "da-mq";
const PERSISTENT_DELIVERY_MODE: u8 = 2;
const MESSAGE_TTL_MS: u32 = 120_000;

pub(crate) type Connection = deadpool::managed::Object<Manager>;

#[derive(Clone, Debug)]
pub struct PublishOptions {
    pub exchange: String,
    pub routing_key: String,
    pub basic_publish_options: BasicPublishOptions,
    pub basic_properties: BasicProperties,
}

pub(crate) fn get_routing_key(rollup_id: u32) -> String {
    return format!("rollup{}", rollup_id);
}

impl Default for PublishOptions {
    fn default() -> Self {
        Self {
            exchange: EXCHANGE_NAME.into(),
            routing_key: DEFAULT_ROUTING_KEY.into(),
            basic_publish_options: BasicPublishOptions::default(),
            basic_properties: BasicProperties::default()
                .with_delivery_mode(PERSISTENT_DELIVERY_MODE)
                .with_expiration(MESSAGE_TTL_MS.to_string().into()),
        }
    }
}

#[derive(Clone, Debug)]
pub struct PublisherContext {
    pub block_hash: CryptoHash,
}

#[derive(Clone, Debug, BorshDeserialize, BorshSerialize)]
pub struct PublishPayload {
    pub transaction_id: CryptoHash,
    pub data: Vec<u8>,
}

#[derive(Clone, Debug)]
pub struct PublishData {
    pub publish_options: PublishOptions,
    pub payload: PublishPayload,
    pub cx: PublisherContext,
}

#[derive(Clone)]
pub struct RmqPublisherHandle {
    pub sender: mpsc::Sender<PublishData>,
}

impl RmqPublisherHandle {
    pub async fn publish(&mut self, publish_data: PublishData) -> Result<()> {
        Ok(self.sender.send(publish_data).await?)
    }

    #[allow(dead_code)]
    pub async fn closed(&self) {
        self.sender.closed().await
    }
}

#[derive(Clone)]
pub struct RmqPublisher {
    connection_pool: Pool,
    listener: Option<PublisherListener>,
}

impl RmqPublisher {
    pub fn new(addr: &str) -> Result<Self> {
        let connection_pool = create_connection_pool(addr.into())?;
        info!(target: PUBLISHER, "Connection pool created, RMQ address {}", addr);

        Ok(Self {
            connection_pool,
            listener: None,
        })
    }

    pub fn run(&self, receiver: mpsc::Receiver<PublishData>) -> JoinHandle<()> {
        let task = RmqPublisherTask::new(self.connection_pool.clone(), self.listener.clone(), receiver);

        actix::spawn(task.run())
    }
}

impl Metricable for RmqPublisher {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()> {
        let listener = make_publisher_metrics(registry)?;
        self.listener = Some(listener);

        Ok(())
    }
}

enum RmqPublisherState {
    Shutdown,
    WaitingForConnection,
    Connected { connection: Connection },
}

struct RmqPublisherTask {
    connection_pool: Pool,
    receiver: mpsc::Receiver<PublishData>,
    listener: Option<PublisherListener>,
}

impl RmqPublisherTask {
    pub fn new(
        connection_pool: Pool,
        listener: Option<PublisherListener>,
        receiver: mpsc::Receiver<PublishData>,
    ) -> Self {
        Self {
            connection_pool,
            receiver,
            listener,
        }
    }

    pub async fn run(mut self) {
        const RECONNECTION_INTERVAL: Duration = Duration::from_secs(2);

        let mut next_step = self.connect().await;
        loop {
            next_step = match next_step {
                RmqPublisherState::WaitingForConnection => {
                    tokio::time::sleep(RECONNECTION_INTERVAL).await;

                    info!(target: PUBLISHER, "Reconnecting to RMQ");
                    self.connect().await
                }
                RmqPublisherState::Connected { connection } => {
                    info!(target: PUBLISHER, "RMQ connected");
                    self.process_stream(connection).await
                }
                RmqPublisherState::Shutdown => {
                    self.receiver.close();
                    return;
                }
            }
        }
    }

    async fn exchange_declare(connection: &Connection) -> Result<(), lapin::Error> {
        let channel = connection.create_channel().await?;
        channel
            .exchange_declare(
                EXCHANGE_NAME,
                ExchangeKind::Topic,
                ExchangeDeclareOptions {
                    passive: false,
                    durable: true,
                    auto_delete: false,
                    internal: false,
                    nowait: false,
                },
                FieldTable::default(),
            )
            .await?;

        Ok(())
    }

    async fn connect(&mut self) -> RmqPublisherState {
        let Self { connection_pool, .. } = self;
        let connection = match connection_pool.get().await {
            Ok(connection) => connection,
            Err(err) => {
                return match err {
                    PoolError::Timeout(_) | PoolError::Backend(_) => RmqPublisherState::WaitingForConnection,
                    PoolError::Closed | PoolError::NoRuntimeSpecified | PoolError::PostCreateHook(_) => {
                        RmqPublisherState::Shutdown
                    }
                }
            }
        };

        match Self::exchange_declare(&connection).await {
            Ok(_) => RmqPublisherState::Connected { connection },
            Err(err) => {
                error!(target: PUBLISHER, "Failed to declare exchange: {}", err);
                RmqPublisherState::WaitingForConnection
            }
        }
    }

    async fn publish(
        connection: &Connection,
        payload: &[u8],
        publish_options: PublishOptions,
    ) -> Result<(), lapin::Error> {
        let PublishOptions {
            exchange,
            routing_key,
            basic_publish_options,
            basic_properties,
        } = publish_options;

        let channel = connection.create_channel().await?;
        channel
            .basic_publish(
                &exchange,
                &routing_key,
                basic_publish_options,
                &payload,
                basic_properties,
            )
            .await?;

        Ok(())
    }

    async fn process_stream(&mut self, connection: Connection) -> RmqPublisherState {
        while let Some(publish_data) = self.receiver.recv().await {
            let mut payload: Vec<u8> = Vec::new();
            if let Err(err) = publish_data.payload.serialize(&mut payload) {
                info!(target: PUBLISHER, "couldn't serialize publish payload {}", err.to_string());
                continue;
            }

            let start_time = std::time::Instant::now();
            match Self::publish(&connection, &payload, publish_data.publish_options.clone()).await {
                Ok(_) => {
                    let duration = start_time.elapsed();
                    self.listener.as_ref().map(|l| {
                        l.num_published_blocks.inc();
                        l.publish_duration_histogram.observe(duration.as_millis() as f64);
                    });
                    info!(target: PUBLISHER, "published tx: {}, routing_key: {}", publish_data.payload.transaction_id, publish_data.publish_options.routing_key);
                }
                Err(err) => {
                    self.listener.as_ref().map(|l| l.num_failed_publishes.inc());
                    Self::handle_error(err, Some(publish_data));

                    return RmqPublisherState::WaitingForConnection;
                }
            }
        }

        RmqPublisherState::Shutdown
    }

    fn handle_error(error: impl Into<Error>, publish_data: Option<PublishData>) {
        let error = error.into();
        let msg = if let Some(data) = publish_data {
            // TODO: add display for cx
            // TODO: handle error here
            format!("Publisher Error: {}, cx: {}", error.to_string(), data.cx.block_hash)
        } else {
            format!("Publisher Error: {}", error.to_string())
        };

        error!(target: PUBLISHER, message = display(msg.as_str()));
    }
}

pub(crate) fn create_connection_pool(addr: String) -> Result<Pool> {
    let options = ConnectionProperties::default()
        .with_executor(tokio_executor_trait::Tokio::current())
        // TODO: reactor is only available for unix.
        .with_reactor(tokio_reactor_trait::Tokio);

    let manager = Manager::new(addr, options);
    let pool: Pool = Pool::builder(manager).max_size(10).build()?;

    Ok(pool)
}
