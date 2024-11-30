use near_config_utils::DownloadConfigType;
use near_indexer::near_primitives::types::{AccountId, Gas};
use std::collections::HashMap;
use std::net::SocketAddr;
use std::str::FromStr;

use crate::errors::{Error, Result};

use serde::Deserialize;

#[derive(clap::Parser, Debug)]
#[clap(version = "0.0.1")]
#[clap(subcommand_required = true, arg_required_else_help = true)]
pub(crate) struct Opts {
    /// Sets a custom config dir. Defaults to ~/.near/
    #[clap(long)]
    pub home_dir: Option<std::path::PathBuf>,
    #[clap(subcommand)]
    pub subcmd: SubCommand,
}

#[derive(clap::Parser, Debug)]
pub(crate) enum SubCommand {
    /// Run NEAR Indexer Example. Start observe the network
    Run(RunConfigParams),
    /// Initialize necessary configs
    Init(InitConfigParams),
}

#[derive(clap::Parser, Deserialize, Debug)]
#[command(group = clap::ArgGroup::new("config_path").conflicts_with("config_args").multiple(false))]
pub(crate) struct RunConfigParams {
    #[clap(long)]
    #[arg(group = "config_path")]
    pub config: Option<std::path::PathBuf>,

    #[clap(flatten)]
    pub run_config_args: Option<RunConfigArgs>,
}

#[derive(clap::Parser, Deserialize, Debug)]
#[group(id = "config_args", conflicts_with = "config_path")]
pub(crate) struct RunConfigArgs {
    /// Rabbit mq address
    #[clap(long, default_value = "amqp://localhost:5672")]
    pub rmq_address: String,
    /// Data availability contract
    #[clap(short, long)]
    pub da_contract_ids: Vec<String>,
    /// Target Rollup ID
    #[clap(long)]
    pub rollup_ids: Vec<u32>,
    /// Metrics socket addr
    #[clap(long)]
    pub metrics_ip_port_address: Option<SocketAddr>,
    /// Address of fastnear block producer.
    #[clap(long, default_value = "https://testnet.neardata.xyz/v0/last_block/final")]
    pub fastnear_address: String,    
    /// Internal FastIndexer channels width. 
    #[clap(long, default_value = "256")]
    pub channel_width: usize
}

impl RunConfigArgs {
    pub(crate) fn compile_addresses_to_ids_map(&self) -> Result<HashMap<AccountId, u32>> {
        if self.rollup_ids.len() != self.da_contract_ids.len() {
            return Err(Error::IDsAndContractAddressesError);
        }

        let addresses: Vec<AccountId> = self
            .da_contract_ids
            .iter()
            .map(|el| el.parse())
            .collect::<Result<_, _>>()?;

        let map = self
            .rollup_ids
            .iter()
            .zip(addresses)
            .map(|(id, addr)| (addr, *id))
            .collect::<HashMap<AccountId, u32>>();
        Ok(map)
    }
}

#[derive(clap::Parser, Deserialize, Debug)]
#[command(group = clap::ArgGroup::new("config_path").conflicts_with("config_args").multiple(false))]
pub(crate) struct InitConfigParams {
    #[clap(long)]
    #[arg(group = "config_path")]
    pub config: Option<std::path::PathBuf>,

    #[clap(flatten)]
    pub args: Option<InitConfigArgs>,
}

mod download_config_type_option_serde {
    use super::*;
    use serde::{Deserialize, Deserializer, Serializer};

    #[allow(dead_code)]
    pub fn serialize<S>(value: &Option<DownloadConfigType>, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: Serializer,
    {
        match value {
            Some(v) => serializer.serialize_str(&v.to_string()),
            None => serializer.serialize_none(),
        }
    }

    pub fn deserialize<'de, D>(deserializer: D) -> Result<Option<DownloadConfigType>, D::Error>
    where
        D: Deserializer<'de>,
    {
        let s: Option<String> = Option::deserialize(deserializer)?;
        s.map(|s| DownloadConfigType::from_str(&s).map_err(serde::de::Error::custom))
            .transpose()
    }
}

#[derive(clap::Parser, Deserialize, Debug)]
#[group(id = "config_args", conflicts_with = "config_path")]
pub(crate) struct InitConfigArgs {
    /// chain/network id (localnet, testnet, devnet, betanet)
    #[clap(short, long)]
    pub chain_id: Option<String>,
    /// Account ID for the validator key
    #[clap(long)]
    pub account_id: Option<String>,
    /// Specify private key generated from seed (TESTING ONLY)
    #[clap(long)]
    pub test_seed: Option<String>,
    /// Number of shards to initialize the chain with
    #[clap(short, long, default_value = "1")]
    pub num_shards: u64,
    /// Makes block production fast (TESTING ONLY)
    #[clap(short, long)]
    pub fast: bool,
    /// Genesis file to use when initialize testnet (including downloading)
    #[clap(short, long)]
    pub genesis: Option<String>,
    #[clap(long)]
    /// Download the verified NEAR genesis file automatically.
    pub download_genesis: bool,
    /// Specify a custom download URL for the genesis-file.
    #[clap(long)]
    pub download_genesis_url: Option<String>,
    /// Specify a custom download URL for the records-file.
    #[clap(long)]
    pub download_records_url: Option<String>,
    #[clap(long)]
    #[serde(with = "download_config_type_option_serde")]
    /// Download the verified NEAR config file automatically.
    pub download_config: Option<DownloadConfigType>,
    /// Specify a custom download URL for the config file.
    #[clap(long)]
    pub download_config_url: Option<String>,
    /// Specify the boot nodes to bootstrap the network
    pub boot_nodes: Option<String>,
    /// Specify a custom max_gas_burnt_view limit.
    #[clap(long)]
    pub max_gas_burnt_view: Option<Gas>,
}

impl From<InitConfigArgs> for near_indexer::InitConfigArgs {
    fn from(config_args: InitConfigArgs) -> Self {
        Self {
            chain_id: config_args.chain_id,
            account_id: config_args.account_id,
            test_seed: config_args.test_seed,
            num_shards: config_args.num_shards,
            fast: config_args.fast,
            genesis: config_args.genesis,
            download_genesis: config_args.download_genesis,
            download_genesis_url: config_args.download_genesis_url,
            download_records_url: config_args.download_records_url,
            download_config: config_args.download_config,
            download_config_url: config_args.download_config_url,
            boot_nodes: config_args.boot_nodes,
            max_gas_burnt_view: config_args.max_gas_burnt_view,
        }
    }
}
