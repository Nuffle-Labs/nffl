mod config;
mod metrics;
mod relayer;

use anyhow::Result;
use relayer::Relayer;
use tracing_subscriber::FmtSubscriber;

#[tokio::main]
async fn main() -> Result<()> {
    let subscriber = FmtSubscriber::new();
    tracing::subscriber::set_global_default(subscriber)?;

    let config = config::load_config()?;
    let mut relayer = Relayer::new(config).await?;
    relayer.start().await?;

    Ok(())
}