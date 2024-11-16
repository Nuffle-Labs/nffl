//! Main off-chain workflow for Nuff DVN.

use eyre::Result;
use offchain::workers::dvn::Dvn;
use tracing::level_filters::LevelFilter;
use tracing_subscriber::EnvFilter;

#[tokio::main]
async fn main() -> Result<()> {
    // Initialize tracing
    tracing_subscriber::fmt()
        .with_target(false)
        .with_env_filter(
            EnvFilter::builder()
                .with_default_directive(LevelFilter::INFO.into())
                .from_env_lossy(),
        )
        .init();

    let mut dvn = Dvn::new_from_env()?;
    dvn.listen().await
}
