use clap::Parser;
use configs::{Opts, SubCommand};
use prometheus::Registry;
use tracing::{error, info};

use crate::{
    candidates_validator::CandidatesValidator, configs::RunConfigArgs, errors::Error, errors::Result,
    indexer_wrapper::IndexerWrapper, metrics::Metricable, metrics_server::MetricsServer, rmq_publisher::RmqPublisher,
};
use crate::fastnear_indexer::FastNearIndexer;
use crate::metrics::INDEXER_NAMESPACE;

mod block_listener;
mod candidates_validator;
mod configs;
mod errors;
mod indexer_wrapper;
mod metrics;
mod metrics_server;
mod rmq_publisher;
mod types;
mod fastnear_indexer;

const INDEXER: &str = "indexer";

fn run(home_dir: std::path::PathBuf, config: RunConfigArgs) -> Result<()> {
    let addresses_to_rollup_ids = config.compile_addresses_to_ids_map()?;
    let system = actix::System::new();
    let registry = Registry::new();
    let server_handle = if let Some(metrics_addr) = config.metrics_ip_port_address {
        let metrics_server = MetricsServer::new(metrics_addr, registry.clone());
        Some(system.runtime().spawn(metrics_server.run()))
    } else {
        None
    };

    let block_res = system.block_on(async move {
        let mut rmq_publisher = RmqPublisher::new(&config.rmq_address)?;
        if config.metrics_ip_port_address.is_some() {
            rmq_publisher.enable_metrics(registry.clone())?;
        }

        if cfg!(feature = "use_fastnear") {
            let fastnear_indexer = FastNearIndexer::new(addresses_to_rollup_ids, config.channel_width);
            let validated_stream = fastnear_indexer.run();

            rmq_publisher.run(validated_stream);

            Ok::<_, Error>(())
        } else {
            let indexer_config = near_indexer::IndexerConfig {
                home_dir,
                sync_mode: near_indexer::SyncModeEnum::LatestSynced,
                await_for_node_synced: near_indexer::AwaitForNodeSyncedEnum::WaitForFullSync,
                validate_genesis: true,
            };

            let mut indexer = IndexerWrapper::new(indexer_config, addresses_to_rollup_ids);
            if config.metrics_ip_port_address.is_some() {
                indexer.enable_metrics(registry.clone())?;
            }

            let (view_client, _) = indexer.client_actors();
            let (block_handle, candidates_stream) = indexer.run();
            let mut candidates_validator = CandidatesValidator::new(view_client);
            if config.metrics_ip_port_address.is_some() {
                candidates_validator.enable_metrics(registry.clone())?;
            }

            let validated_stream = candidates_validator.run(candidates_stream);
            
            rmq_publisher.run(validated_stream);

            // TODO: Handle block_handle whether cancelled or panics
            Ok::<_, Error>(block_handle.await?)
        }
    });

    if let Some(handle) = server_handle {
        handle.abort();
    }

    // Run until publishing finished
    system.run()?;

    block_res.map_err(|err| {
        error!(target: INDEXER, "Indexer Error: {}", err);
        err
    })
}

fn read_config<T: serde::de::DeserializeOwned>(
    config_path: Option<std::path::PathBuf>,
    config_args: Option<T>,
) -> Result<T> {
    if let Some(config_path) = config_path {
        let config_str = std::fs::read_to_string(config_path)?;
        serde_yaml::from_str(&config_str).map_err(Into::into)
    } else {
        config_args.ok_or_else(|| Error::AnyhowError(anyhow::anyhow!("Either config_path or config_args must be provided")))
    }
}

fn main() -> Result<()> {
    info!(target: INDEXER_NAMESPACE, "Starting...");

    // We use it to automatically search the for root certificates to perform HTTPS calls
    // (sending telemetry and downloading genesis)
    openssl_probe::init_ssl_cert_env_vars();
    let env_filter = near_o11y::tracing_subscriber::EnvFilter::new(
        "nearcore=info,publisher=info,indexer=info,candidates_validator=info,\
         metrics=info,tokio_reactor=info,near=info,stats=info,telemetry=info,\
         near-performance-metrics=info,fastnear_indexer=info",
    );
    let _subscriber = near_o11y::default_subscriber(env_filter, &Default::default()).global();
    let opts: Opts = Opts::parse();

    let home_dir = opts.home_dir.unwrap_or_else(near_indexer::get_default_home);
    match opts.subcmd {
        SubCommand::Init(params) => {
            near_indexer::indexer_init_configs(&home_dir, read_config(params.config, params.args)?.into())?;
            Ok(())
        }
        SubCommand::Run(params) => run(home_dir, read_config(params.config, params.run_config_args)?),
    }
}