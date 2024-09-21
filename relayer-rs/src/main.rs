use crate::config::RelayerConfig;
use crate::relayer::Relayer;
use anyhow::Result;
use clap::{App, Arg, SubCommand};
use prometheus::Registry;
use std::path::PathBuf;
use tracing_subscriber::FmtSubscriber;

mod config;
mod metrics;
mod relayer;

#[tokio::main]
async fn main() -> Result<()> {
    let matches = App::new("sffl-test-relayer")
        .about("SFFL Test Relayer")
        .subcommand(
            SubCommand::with_name("run-args")
                .about("Start the relayer with direct CLI options")
                .arg(Arg::with_name("production").long("production").help("Run in production logging mode"))
                .arg(Arg::with_name("rpc-url").long("rpc-url").required(true).takes_value(true).help("Connect to the indicated RPC"))
                .arg(Arg::with_name("da-account-id").long("da-account-id").required(true).takes_value(true).help("Publish block data to the indicated NEAR account"))
                .arg(Arg::with_name("key-path").long("key-path").required(true).takes_value(true).help("Path to NEAR account's key file"))
                .arg(Arg::with_name("network").long("network").takes_value(true).default_value("http://127.0.0.1:3030").help("Network for NEAR client to use"))
                .arg(Arg::with_name("metrics-ip-port-address").long("metrics-ip-port-address").takes_value(true).help("Metrics scrape address")),
        )
        .subcommand(
            SubCommand::with_name("run-config")
                .about("Start the relayer using a configuration file")
                .arg(Arg::with_name("path").short("p").long("path").required(true).takes_value(true).help("Load configuration from FILE")),
        )
        .get_matches();

    match matches.subcommand() {
        ("run-args", Some(args)) => {
            let config = RelayerConfig {
                production: args.is_present("production"),
                rpc_url: args.value_of("rpc-url").unwrap().to_string(),
                da_account_id: args.value_of("da-account-id").unwrap().to_string(),
                key_path: args.value_of("key-path").unwrap().to_string(),
                network: args.value_of("network").unwrap().to_string(),
                metrics_ip_port_addr: args.value_of("metrics-ip-port-address").map(String::from),
            };
            relayer_main(config).await
        }
        ("run-config", Some(args)) => {
            let config_path = args.value_of("path").unwrap();
            let config = config::load_config(PathBuf::from(config_path))?;
            relayer_main(config).await
        }
        _ => Err(anyhow::anyhow!("Invalid subcommand")),
    }
}

async fn relayer_main(config: RelayerConfig) -> Result<()> {
    let log_level = if config.production {
        tracing::Level::INFO
    } else {
        tracing::Level::DEBUG
    };

    let subscriber = FmtSubscriber::builder()
        .with_max_level(log_level)
        .finish();
    tracing::subscriber::set_global_default(subscriber)?;

    tracing::info!("Initializing Relayer");
    tracing::info!("Read config: {:?}", config);

    let mut relayer = Relayer::new(config.clone()).await?;

    if let Some(metrics_addr) = config.metrics_ip_port_addr {
        let registry = Registry::new();
        relayer.enable_metrics(&registry)?;
        metrics::start_metrics_server(metrics_addr, registry);
    }

    tracing::info!("Starting relayer");
    relayer.start().await?;

    Ok(())
}