use offchain::{config, workers::executor::NFFLExecutor};
use tracing::level_filters::LevelFilter;
use tracing_subscriber::EnvFilter;

/// Executor is expected to work with low work rate, and we have a bonus
/// from this observation - we don't need/want to care about concurrency control,
/// so we choose run single-threaded runtime so far.
#[tokio::main(flavor = "current_thread")]
async fn main() -> eyre::Result<()> {
    // Initialize tracing
    tracing_subscriber::fmt()
        .with_target(false)
        .with_env_filter(
            EnvFilter::builder()
                .with_default_directive(LevelFilter::DEBUG.into())
                .from_env_lossy(),
        )
        .init();

    let mut executor = NFFLExecutor::new(config::WorkerConfig::load_from_env()?);
    executor.listen().await?;

    Ok(())
}
