use crate::errors::Result;

use deadpool_lapin::{Manager, Pool};
use lapin::ConnectionProperties;

pub(crate) type Connection = deadpool::managed::Object<Manager>;

pub(crate) fn create_connection_pool(addr: String) -> Result<Pool> {
    let options = ConnectionProperties::default()
        .with_executor(tokio_executor_trait::Tokio::current())
        // TODO: reactor is only available for unix.
        .with_reactor(tokio_reactor_trait::Tokio);

    let manager = Manager::new(addr, options);
    let pool: Pool = Pool::builder(manager).max_size(10).build()?;

    Ok(pool)
}
