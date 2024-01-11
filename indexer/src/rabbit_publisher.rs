use crate::errors::{Error, Result};

use deadpool_lapin::{Manager, Pool};
use lapin::options::BasicPublishOptions;
use lapin::{BasicProperties, ConnectionProperties};
use tokio::sync::mpsc;

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
struct PublishData {
    pub publish_options: PublishOptions,
    pub payload: Vec<u8>,
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

    pub async fn publish(&mut self, data: &Vec<u8>) -> Result<()> {
        self.publish_custom(PublishOptions::default(), data).await
    }

    pub async fn publish_custom(&mut self, publish_options: PublishOptions, data: &Vec<u8>) -> Result<()> {
        self.sender
            .send(PublishData {
                publish_options,
                payload: data.clone(),
            })
            .await
            .map_err(|_| Error::SendError)
    }

    async fn publisher(connection_pool: Pool, mut receiver: mpsc::Receiver<PublishData>) {
        let mut connection = match connection_pool.get().await {
            Ok(connection) => connection,
            Err(err) => {
                Self::handle_error(err, None);
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

        while let Some(publish_data) = receiver.recv().await {
            match logic(connection_pool.clone(), connection, publish_data.clone()).await {
                Ok(new_connection) => connection = new_connection,
                Err(err) => {
                    Self::handle_error(err, Some(publish_data));
                    break;
                }
            }
        }

        receiver.close();

        // Drain receiver
        while let Some(_) = receiver.recv().await {}
    }

    fn handle_error(error: impl Into<Error>, _publish_data: Option<PublishData>) {
        // TODO: handle error here
        let error = error.into();
        eprintln!("{}", error.to_string());
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
