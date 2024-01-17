use deadpool_lapin::{Manager, Pool};
use lapin::options::BasicPublishOptions;
use lapin::{BasicProperties, ConnectionProperties};
use near_indexer::near_primitives::hash::CryptoHash;
use near_indexer::near_primitives::types::TransactionOrReceiptId;
use tokio::sync::mpsc;
use tracing::error;

use crate::errors::{Error, Result};

const DEFAULT_EXCHANGE: &str = "";
const DEFAULT_ROUTING_KEY: &str = "da-mq";

pub(crate) type Connection = deadpool::managed::Object<Manager>;

#[derive(Clone, Debug)]
pub struct PublishOptions {
    pub exchange: String,
    pub routing_key: String,
    pub basic_publish_options: BasicPublishOptions,
    pub basic_properties: BasicProperties,
}

impl Default for PublishOptions {
    fn default() -> Self {
        Self {
            exchange: DEFAULT_EXCHANGE.into(),
            routing_key: DEFAULT_ROUTING_KEY.into(),
            basic_publish_options: BasicPublishOptions::default(),
            basic_properties: BasicProperties::default(),
        }
    }
}

#[derive(Clone, Debug)]
pub struct PublisherContext {
    pub block_hash: CryptoHash,
    pub id: TransactionOrReceiptId,
}

#[derive(Clone, Debug)]
pub struct PublishData {
    pub publish_options: PublishOptions,
    pub payload: Vec<u8>,
    pub cx: PublisherContext,
}

pub struct RabbitBuilder {
    addr: String,
}

impl RabbitBuilder {
    pub fn new(addr: String) -> Self {
        Self { addr }
    }

    /// Shall be called within actix runtime
    pub fn build(self) -> Result<RabbitPublisher> {
        RabbitPublisher::new(&self.addr)
    }
}

#[derive(Clone)]
pub struct RabbitPublisher {
    sender: mpsc::Sender<PublishData>,
}

// TODO: try to put error in inner state?
impl RabbitPublisher {
    pub fn new(addr: &str) -> Result<Self> {
        let pool = create_connection_pool(addr.into())?;

        let (sender, receiver) = mpsc::channel(100);
        actix::spawn(Self::publisher(pool, receiver));

        Ok(Self { sender })
    }

    pub async fn publish(&mut self, publish_data: PublishData) -> Result<()> {
        Ok(self.sender.send(publish_data).await?)
    }

    async fn publisher(connection_pool: Pool, mut receiver: mpsc::Receiver<PublishData>) {
        const ERROR_CODE: i32 = 1;

        let mut connection = match connection_pool.get().await {
            Ok(connection) => connection,
            Err(err) => {
                Self::handle_error(err, None);
                actix::System::current().stop_with_code(ERROR_CODE);
                return;
            }
        };

        let logic = |connection_pool: Pool, mut connection: Connection, publish_data: PublishData| async move {
            if !connection.status().connected() {
                connection = connection_pool.get().await?;
            }

            let channel = connection.create_channel().await?;

            let PublishOptions {
                exchange,
                routing_key,
                basic_publish_options,
                basic_properties,
            } = publish_data.publish_options.clone();

            channel
                .basic_publish(
                    &exchange,
                    &routing_key,
                    basic_publish_options,
                    &publish_data.payload,
                    basic_properties,
                )
                .await?;

            Ok::<_, Error>(connection)
        };

        let code = loop {
            match receiver.recv().await {
                Some(publish_data) => match logic(connection_pool.clone(), connection, publish_data.clone()).await {
                    Ok(new_connection) => connection = new_connection,
                    Err(err) => {
                        Self::handle_error(err, Some(publish_data));
                        break ERROR_CODE;
                    }
                },
                None => break 0,
            };
        };

        receiver.close();
        // Drain receiver
        while let Some(_) = receiver.recv().await {}

        actix::System::current().stop_with_code(code);
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

        error!(target: "publisher", message = display(msg.as_str()));
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

// indexes, check execution, publish
// shall be able to abort with
// Stop system with log?
