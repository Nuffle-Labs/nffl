use crate::config::RelayerConfig;
use crate::relayer::Relayer;
use anyhow::Result;
use clap::{Command, Arg};
use prometheus::Registry;
use std::path::PathBuf;
use tracing_subscriber::FmtSubscriber;
use crate::metrics::RelayerMetrics;

mod config;
mod metrics;
mod relayer;

#[tokio::main]
async fn main() -> Result<()> {
    let matches = Command::new("sffl-test-relayer")
        .about("SFFL Test Relayer")
        .subcommand(
            Command::new("run-args")
                .about("Start the relayer with direct CLI options")
                .arg(Arg::new("production").long("production").help("Run in production logging mode"))
                .arg(Arg::new("rpc-url").long("rpc-url").required(true).value_name("URL").help("Connect to the indicated RPC"))
                .arg(Arg::new("da-account-id").long("da-account-id").required(true).value_name("ACCOUNT").help("Publish block data to the indicated NEAR account"))
                .arg(Arg::new("key-path").long("key-path").required(true).value_name("FILE").help("Path to NEAR account's key file"))
                .arg(Arg::new("network").long("network").value_name("URL").default_value("http://127.0.0.1:3030").help("Network for NEAR client to use"))
                .arg(Arg::new("metrics-ip-port-address").long("metrics-ip-port-address").value_name("ADDRESS").help("Metrics scrape address")),
        )
        .subcommand(
            Command::new("run-config")
                .about("Start the relayer using a configuration file")
                .arg(Arg::new("path").short('p').long("path").required(true).value_name("FILE").help("Load configuration from FILE")),
        )
        .get_matches();

    match matches.subcommand() {
        Some(("run-args", args)) => {
            let config = RelayerConfig {
                production: args.contains_id("production"),
                rpc_url: args.get_one::<String>("rpc-url").unwrap().to_string(),
                da_account_id: args.get_one::<String>("da-account-id").unwrap().to_string(),
                key_path: args.get_one::<String>("key-path").unwrap().to_string(),
                network: args.get_one::<String>("network").unwrap().to_string(),
                metrics_ip_port_addr: args.get_one::<String>("metrics-ip-port-address").map(|s| s.to_string()),
            };
            relayer_main(config).await
        }
        Some(("run-config", args)) => {
            let config_path = args.get_one::<String>("path").unwrap();
            let config = config::load_config(PathBuf::from(config_path)).map_err(|e| anyhow::anyhow!(e))?;
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

    let mut relayer = Relayer::new(&config).await?;

    if let Some(metrics_addr) = config.metrics_ip_port_addr {
        let registry = Registry::new();
        tokio::spawn(metrics::start_metrics_server(metrics_addr, registry));
    }

    tracing::info!("Starting relayer");
    relayer.start().await?;

    Ok(())
}