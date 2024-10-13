use anyhow::Result;
use clap::{Command, Arg};
use eigensdk::logging::logger::Logger;
use operator_rs::attestor::Attestor;
use operator_rs::consumer::Consumer;
use operator_rs::operator::Operator;
use operator_rs::types::NFFLNodeConfig;
use prometheus::Registry;
use std::path::PathBuf;
use tracing_subscriber::FmtSubscriber;
use serde_json::Value;
use eigensdk::crypto_bls::BlsKeyPair;


mod config;

#[tokio::main]
async fn main() -> Result<()> {
    // Print all arguments for debugging
    println!("Received arguments: {:?}", std::env::args().collect::<Vec<String>>());

    let matches = Command::new("sffl-operator")
        .about("SFFL Operator")
        .subcommand(
            Command::new("run-args")
                .about("Start the operator with direct CLI options")
                .arg(Arg::new("production").long("production").help("Run in production logging mode"))
                .arg(Arg::new("rpc-url").long("rpc-url").value_name("URL").help("Connect to the indicated RPC"))
                .arg(Arg::new("ws-url").long("ws-url").value_name("URL").help("Connect to the indicated WebSocket"))
                .arg(Arg::new("operator-address").long("operator-address").value_name("ADDRESS").help("Operator's Ethereum address"))
                .arg(Arg::new("bls-key-path").long("bls-key-path").value_name("FILE").help("Path to BLS key file"))
                .arg(Arg::new("ecdsa-key-path").long("ecdsa-key-path").value_name("FILE").help("Path to ECDSA key file"))
                .arg(Arg::new("avs-registry-coordinator").long("avs-registry-coordinator").value_name("ADDRESS").help("AVS Registry Coordinator address"))
                .arg(Arg::new("operator-state-retriever").long("operator-state-retriever").value_name("ADDRESS").help("Operator State Retriever address"))
                .arg(Arg::new("aggregator-server-ip-port").long("aggregator-server-ip-port").value_name("IP:PORT").help("Aggregator server IP and port"))
                .arg(Arg::new("enable-metrics").long("enable-metrics").help("Enable metrics collection"))
                .arg(Arg::new("eigen-metrics-ip-port").long("eigen-metrics-ip-port").value_name("IP:PORT").help("EigenMetrics IP and port"))
                .arg(Arg::new("node-api-ip-port").long("node-api-ip-port").value_name("IP:PORT").help("Node API IP and port"))
                .arg(Arg::new("task-response-wait-ms").long("task-response-wait-ms").value_name("MS").help("Task response wait time in milliseconds"))
                .arg(Arg::new("register-operator-on-startup").long("register-operator-on-startup").help("Register operator on startup"))
                .arg(Arg::new("token-strategy-addr").long("token-strategy-addr").value_name("ADDRESS").help("Token strategy address"))
                .arg(Arg::new("enable-node-api").long("enable-node-api").help("Enable Node API"))
                .arg(Arg::new("near-da-indexer-rmq-ip-port-address").long("near-da-indexer-rmq-ip-port-address").value_name("IP:PORT").help("NEAR DA Indexer RabbitMQ IP and port"))
                .arg(Arg::new("near-da-indexer-rollup-ids").long("near-da-indexer-rollup-ids").value_name("IDS").help("Comma-separated list of NEAR DA Indexer rollup IDs"))
                .arg(Arg::new("rollup-ids-to-rpc-urls").long("rollup-ids-to-rpc-urls").value_name("ID:URL").help("Comma-separated list of rollup ID to RPC URL mappings"))
        )
        .subcommand(
            Command::new("run-config")
                .about("Start the operator using a configuration file")
                .arg(Arg::new("path").short('p').long("path").required(true).value_name("FILE").help("Load configuration from FILE")),
        )
        .get_matches();

    match matches.subcommand() {
        Some(("run-args", args)) => {
            let config = NFFLNodeConfig {
                enable_node_api: args.contains_id("enable-node-api"),
                near_da_indexer_rmq_ip_port_address: args.get_one::<String>("near-da-indexer-rmq-ip-port-address")
                    .map(|s| s.to_string())
                    .unwrap_or_default(),
                near_da_indexer_rollup_ids: args.get_one::<String>("near-da-indexer-rollup-ids")
                    .map(|s| s.split(',').filter_map(|id| id.parse().ok()).collect())
                    .unwrap_or_default(),
                rollup_ids_to_rpc_urls: args.get_one::<String>("rollup-ids-to-rpc-urls")
                    .map(|s| s.split(',')
                        .filter_map(|pair| {
                            let mut parts = pair.split(':');
                            let id = parts.next()?.parse().ok()?;
                            let url = parts.next()?.to_string();
                            Some((id, url))
                        })
                        .collect())
                    .unwrap_or_default(),
                production: args.contains_id("production"),
                eth_rpc_url: args.get_one::<String>("rpc-url").map(|s| s.to_string()).unwrap_or_default(),
                eth_ws_url: args.get_one::<String>("ws-url").map(|s| s.to_string()).unwrap_or_default(),
                operator_address: args.get_one::<String>("operator-address").map(|s| s.to_string()).unwrap_or_default(),
                ecdsa_private_key_store_path: args.get_one::<String>("ecdsa-key-path").map(|s| s.to_string()).unwrap_or_default(),
                bls_keypair: args.get_one::<String>("bls-key-path").map(|s| s.to_string()).unwrap_or_default(),
                avs_registry_coordinator_address: args.get_one::<String>("avs-registry-coordinator").map(|s| s.to_string()).unwrap_or_default(),
                operator_state_retriever_address: args.get_one::<String>("operator-state-retriever").map(|s| s.to_string()).unwrap_or_default(),
                aggregator_server_ip_port_address: args.get_one::<String>("aggregator-server-ip-port").map(|s| s.to_string()).unwrap_or_default(),
                enable_metrics: args.contains_id("enable-metrics"),
                eigen_metrics_ip_port_address: args.get_one::<String>("eigen-metrics-ip-port").map(|s| s.to_string()).unwrap_or_default(),
                node_api_ip_port_address: args.get_one::<String>("node-api-ip-port").map(|s| s.to_string()).unwrap_or_default(),
                task_response_wait_ms: args.get_one::<String>("task-response-wait-ms")
                    .and_then(|s| s.parse().ok())
                    .unwrap_or(0),
                register_operator_on_startup: args.contains_id("register-operator-on-startup"),
                token_strategy_addr: args.get_one::<String>("token-strategy-addr").map(|s| s.to_string()).unwrap_or_default(),
            };
            operator_main(config).await
        }
        Some(("run-config", args)) => {
            let config_path = args.get_one::<String>("path").unwrap();
            let config = config::load_config(PathBuf::from(config_path))?;
            operator_main(config).await
        }
        _ => Err(anyhow::anyhow!("Invalid subcommand")),
    }
}

async fn operator_main(config: NFFLNodeConfig) -> Result<()> {
    let log_level = if config.production {
        tracing::Level::INFO
    } else {
        tracing::Level::DEBUG
    };

    let subscriber = FmtSubscriber::builder()
        .with_max_level(log_level)
        .finish();
    tracing::subscriber::set_global_default(subscriber)?;

    tracing::info!("Initializing Operator");
    tracing::info!("Read config: {:?}", config);

    let attestor_config = config.clone();
    let attestor = Attestor::new(
        attestor_config,
        BlsKeyPair::new(config.bls_keypair)?,
        eigensdk::types::operator::OperatorId::from([1u8; 32]),
        Registry::default(),
        operator_rs::types::create_default_logger()
    )?;   

    // Start the attestor
    attestor.start().await?;
    
    // Uncomment these lines when you're ready to use the Operator
    // let operator = Operator::new(config).await?;
    // tracing::info!("Starting operator");
    // let mut ctx = operator_rs::operator::Context::default();
    // operator.start(&mut ctx).await?;

    Ok(())
}
