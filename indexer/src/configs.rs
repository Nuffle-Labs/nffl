use near_indexer::near_primitives::types::{AccountId, Gas};
use std::collections::HashMap;

use crate::errors::{Error, Result};

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
    Run(RunConfigArgs),
    /// Initialize necessary configs
    Init(InitConfigArgs),
}

#[derive(clap::Parser, Debug)]
pub(crate) struct RunConfigArgs {
    /// Rabbit mq address
    #[clap(long)]
    pub rmq_address: String,
    /// Data availability contract
    #[clap(short, long)]
    pub da_contract_ids: Vec<String>,
    /// Target Rollup ID
    #[clap(long)]
    pub rollup_ids: Vec<u32>,
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

#[derive(clap::Parser, Debug)]
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
    /// Download the verified NEAR config file automatically.
    pub download_config: bool,
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
