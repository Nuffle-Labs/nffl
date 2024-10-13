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

mod config;

#[tokio::main]
async fn main() -> Result<()> {
    let matches = Command::new("sffl-operator")
        .about("SFFL Operator")
        .subcommand(
            Command::new("run-args")
                .about("Start the operator with direct CLI options")
                .arg(Arg::new("production").long("production").help("Run in production logging mode"))
                .arg(Arg::new("rpc-url").long("rpc-url").required(true).value_name("URL").help("Connect to the indicated RPC"))
                .arg(Arg::new("ws-url").long("ws-url").required(true).value_name("URL").help("Connect to the indicated WebSocket"))
                .arg(Arg::new("operator-address").long("operator-address").required(true).value_name("ADDRESS").help("Operator's Ethereum address"))
                .arg(Arg::new("bls-key-path").long("bls-key-path").required(true).value_name("FILE").help("Path to BLS key file"))
                .arg(Arg::new("ecdsa-key-path").long("ecdsa-key-path").required(true).value_name("FILE").help("Path to ECDSA key file"))
                .arg(Arg::new("avs-registry-coordinator").long("avs-registry-coordinator").required(true).value_name("ADDRESS").help("AVS Registry Coordinator address"))
                .arg(Arg::new("operator-state-retriever").long("operator-state-retriever").required(true).value_name("ADDRESS").help("Operator State Retriever address"))
                .arg(Arg::new("aggregator-server-ip-port").long("aggregator-server-ip-port").required(true).value_name("IP:PORT").help("Aggregator server IP and port"))
                .arg(Arg::new("enable-metrics").long("enable-metrics").help("Enable metrics collection"))
                .arg(Arg::new("eigen-metrics-ip-port").long("eigen-metrics-ip-port").value_name("IP:PORT").help("EigenMetrics IP and port"))
                .arg(Arg::new("node-api-ip-port").long("node-api-ip-port").value_name("IP:PORT").help("Node API IP and port"))
                .arg(Arg::new("task-response-wait-ms").long("task-response-wait-ms").value_name("MS").default_value("0").help("Task response wait time in milliseconds"))
                .arg(Arg::new("register-operator-on-startup").long("register-operator-on-startup").help("Register operator on startup"))
                .arg(Arg::new("token-strategy-addr").long("token-strategy-addr").value_name("ADDRESS").help("Token strategy address")),
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
                near_da_indexer_rmq_ip_port_address: args.get_one::<String>("near-da-indexer-rmq-ip-port-address").unwrap().to_string(),
                near_da_indexer_rollup_ids: args.get_one::<String>("near-da-indexer-rollup-ids").unwrap().split(',').map(|s| s.parse().unwrap()).collect(),
                rollup_ids_to_rpc_urls: args.get_one::<String>("rollup-ids-to-rpc-urls")
                    .unwrap()
                    .split(',')
                    .map(|s| {
                        let mut parts = s.split(':');
                        (parts.next().unwrap().parse().unwrap(), parts.next().unwrap().to_string())
                    })
                    .collect(),
                production: args.contains_id("production"),
                eth_rpc_url: args.get_one::<String>("rpc-url").unwrap().to_string(),
                eth_ws_url: args.get_one::<String>("ws-url").unwrap().to_string(),
                operator_address: args.get_one::<String>("operator-address").unwrap().to_string(),
                bls_private_key_store_path: args.get_one::<String>("bls-key-path").unwrap().to_string(),
                ecdsa_private_key_store_path: args.get_one::<String>("ecdsa-key-path").unwrap().to_string(),
                avs_registry_coordinator_address: args.get_one::<String>("avs-registry-coordinator").unwrap().to_string(),
                operator_state_retriever_address: args.get_one::<String>("operator-state-retriever").unwrap().to_string(),
                aggregator_server_ip_port_address: args.get_one::<String>("aggregator-server-ip-port").unwrap().to_string(),
                enable_metrics: args.contains_id("enable-metrics"),
                eigen_metrics_ip_port_address: args.get_one::<String>("eigen-metrics-ip-port").unwrap().to_string(),
                node_api_ip_port_address: args.get_one::<String>("node-api-ip-port").unwrap().to_string(),
                task_response_wait_ms: args.get_one::<String>("task-response-wait-ms").unwrap().parse().unwrap(),
                register_operator_on_startup: args.contains_id("register-operator-on-startup"),
                token_strategy_addr: args.get_one::<String>("token-strategy-addr").unwrap().to_string(),
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
        eigensdk::crypto_bls::BlsKeyPair::new(config.bls_private_key_store_path.clone())?,
        eigensdk::types::operator::OperatorId::from([1u8; 32]),
        Registry::default(),
        operator_rs::types::create_default_logger()
    )?;   
    //Start the attestor
    attestor.start().await?;
    
    
    // let operator = Operator::new(config).await?;
    // tracing::info!("Starting operator");

    // let mut ctx = operator_rs::operator::Context;
    // operator.start(&mut ctx).await?;

    Ok(())
}