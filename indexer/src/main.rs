use clap::Parser;
use configs::{Opts, SubCommand};
use near_indexer::near_primitives::types::AccountId;
use std::ops::Deref;
use tokio::sync::mpsc;

use crate::{
    configs::RunConfigArgs,
    errors::{Error, Result},
    producer::process_receipt_candidates,
    producer::{listen_blocks, CandidateData},
    rabbit_publisher::RabbitBuilder,
};

mod configs;
mod errors;
mod producer;
mod rabbit_publisher;

fn run(home_dir: std::path::PathBuf, config: RunConfigArgs) -> Result<()> {
    let da_contract_id: AccountId = config.da_contract_id.parse()?;
    let rabbit_builder = RabbitBuilder::new(config.rmq_address);

    let indexer_config = near_indexer::IndexerConfig {
        home_dir,
        sync_mode: near_indexer::SyncModeEnum::FromInterruption,
        await_for_node_synced: near_indexer::AwaitForNodeSyncedEnum::WaitForFullSync,
        validate_genesis: true,
    };

    let system = actix::System::new();
    system.block_on(async move {
        let indexer = near_indexer::Indexer::new(indexer_config).expect("Indexer::new()");
        let stream = indexer.streamer();
        let (view_client, _) = indexer.client_actors();

        // TODO: define buffer: usize const
        let (sender, receiver) = mpsc::channel::<CandidateData>(100);

        let rabbit_publisher = rabbit_builder.build()?;
        // TODO handle error
        actix::spawn(process_receipt_candidates(view_client, receiver, rabbit_publisher));
        if let Err(err) = listen_blocks(stream, sender, da_contract_id).await {
            eprintln!("{}", err.to_string());
        }

        actix::System::current().stop();

        Ok::<(), Error>(())
    })?;

    // TODO: unnecessary? since only used when directly spawninng on runtime?
    // TODO: Needed since we have actix::spawn?
    Ok(system.run()?)
}

fn main() -> Result<()> {
    // We use it to automatically search the for root certificates to perform HTTPS calls
    // (sending telemetry and downloading genesis)
    openssl_probe::init_ssl_cert_env_vars();
    let env_filter = near_o11y::tracing_subscriber::EnvFilter::new(
        "nearcore=info,indexer_example=info,tokio_reactor=info,near=info,\
         stats=info,telemetry=info,indexer=info,near-performance-metrics=info",
    );
    let _subscriber = near_o11y::default_subscriber(env_filter, &Default::default()).global();
    let opts: Opts = Opts::parse();

    let home_dir = opts.home_dir.unwrap_or(near_indexer::get_default_home());
    match opts.subcmd {
        SubCommand::Init(config) => near_indexer::indexer_init_configs(&home_dir, config.into())?,
        SubCommand::Run(config) => run(home_dir, config)?,
    }

    Ok(())
}
