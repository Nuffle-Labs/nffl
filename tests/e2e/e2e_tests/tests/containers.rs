use std::path::PathBuf;
use testcontainers::core::WaitFor;
use testcontainers::{ContainerRequest, GenericImage, ImageExt};
use testcontainers_modules::rabbitmq::RabbitMq;
const SHELL_ENTRYPOINT: &str = "sh";
const ANVIL_ENTRYPOINT: &str = "anvil";
const NETWORK_NAME: &str = "nffl";
const ANVIL_STATE_PATH: &str = "../../anvil/data/avs-and-eigenlayer-deployed-anvil-state.json";
const NEAR_KEYS: &str = "near_cli_keys";
const CURRENT_DIR: &str = "./";

#[cfg(test)]
pub fn rabbitmq() -> ContainerRequest<RabbitMq> {
    RabbitMq::default()
        .with_env_var("RABBITMQ_DEFAULT_USER", "guest")
        .with_env_var("RABBITMQ_DEFAULT_PASSWORD", "guest")
        .with_network(NETWORK_NAME)
}

#[cfg(test)]
pub fn anvil_node() -> ContainerRequest<GenericImage> {
    GenericImage::new(
        "ghcr.io/foundry-rs/foundry",
        "latest@sha256:8b843eb65cc7b155303b316f65d27173c862b37719dc095ef3a2ef27ce8d3c00",
    )
        .with_wait_for(WaitFor::message_on_stdout("Listening on 0.0.0.0:8545"))
        .with_entrypoint(ANVIL_ENTRYPOINT)
        .with_copy_to("/root/.anvil/state.json", PathBuf::from(ANVIL_STATE_PATH))
        .with_cmd(vec![
            "--host",
            "0.0.0.0",
            "--port",
            "8545",
            "--chain-id",
            "1",
            "--block-time",
            "5",
            "--load-state",
            "/root/.anvil/state.json",
        ])
        .with_network(NETWORK_NAME)
}

#[cfg(test)]
pub fn anvil_node_setup() -> ContainerRequest<GenericImage> {
    GenericImage::new(
        "ghcr.io/foundry-rs/foundry",
        "latest@sha256:8b843eb65cc7b155303b316f65d27173c862b37719dc095ef3a2ef27ce8d3c00",
    )
        .with_entrypoint(SHELL_ENTRYPOINT)
        .with_cmd(vec![
            "sh",
            "-c",
            "cast rpc anvil_setBalance 0xD5A0359da7B310917d7760385516B2426E86ab7f 0x8ac7230489e80000 \\ \
            cast rpc anvil_setBalance 0x9441540E8183d416f2Dc1901AB2034600f17B65a 0x8ac7230489e80000",
        ])
        .with_env_var("ETH_RPC_URL", "http://mainnet-anvil:8545")
        .with_network(NETWORK_NAME)
    // Execute after start
    //   cast rpc anvil_setBalance 0xD5A0359da7B310917d7760385516B2426E86ab7f 0x8ac7230489e80000
    //   cast rpc anvil_setBalance 0x9441540E8183d416f2Dc1901AB2034600f17B65a 0x8ac7230489e80000
}

#[cfg(test)]
pub fn anvil_rollup_node(port: i32) -> ContainerRequest<GenericImage> {
    GenericImage::new(
        "ghcr.io/foundry-rs/foundry",
        "latest@sha256:8b843eb65cc7b155303b316f65d27173c862b37719dc095ef3a2ef27ce8d3c00",
    )
        .with_entrypoint(ANVIL_ENTRYPOINT)
        .with_cmd(vec![
            "--host",
            "0.0.0.0",
            "--port",
            port.to_string().as_str(),
            "--chain-id",
            "2",
            "--block-time",
            "5",
            "--load-state",
            "/root/.anvil/state.json",
        ])
        .with_copy_to("/root/.anvil/state.json", PathBuf::from(ANVIL_STATE_PATH))
        .with_network(NETWORK_NAME)
}

#[cfg(test)]
pub fn near_da_deployer(indexer_port: i32) -> ContainerRequest<GenericImage> {
    GenericImage::new("node", "16")
        .with_entrypoint(SHELL_ENTRYPOINT)
        .with_cmd(vec![
            "sh", "-c",
            "npm i -g near-cli@3.0.0 \\\
             near create-account da2.test.near --masterAccount test.near \\\
             near deploy da2.test.near /nffl/tests/near/near_da_blob_store.wasm --initFunction new --initArgs {} --masterAccount test.near -f \\\
             near create-account da3.test.near --masterAccount test.near \\\
             near deploy da2.test.near /nffl/tests/near/near_da_blob_store.wasm --initFunction new --initArgs {} --masterAccount test.near -f"])
        .with_env_var("NEAR_ENV", "localhost")
        .with_env_var("NEAR_CLI_LOCALNET_NETWORK_ID", "localhost")
        .with_env_var("NEAR_CLI_LOCALNET_KEY_PATH", "/near-cli/validator_key.json")
        .with_env_var("NEAR_HELPER_ACCOUNT", "near")
        .with_env_var("NEAR_NODE_URL", format!("http://nffl-indexer:{indexer_port}").as_str())
        .with_network(NETWORK_NAME)
}

#[cfg(test)]
#[cfg(target_arch = "x86_64")]
pub fn rollup_relayer(rollup_node_port: i32) -> ContainerRequest<GenericImage> {
    GenericImage::new("ghcr.io/nuffle-labs/nffl/relayer", "66dcb37e32e34f552a63c1e638a57dd251846f63")
        .with_cmd(vec![
            "--rpc-url",
            format!("ws://rollup0-anvil:{rollup_node_port}").as_str(),
            "--da-account-id",
            "da2.test.near",
            "--key-path",
            "/root/.near-credentials/localnet/da2.test.near.json",
            "--network",
            "http://nffl-indexer:3030",
            "--metrics-ip-port-address",
            "rollup0-relayer:9091",
        ])
        .with_copy_to("/root/.near-credentials", PathBuf::from(ANVIL_STATE_PATH))
        .with_network(NETWORK_NAME)
}

#[cfg(test)]
#[cfg(target_arch = "x86_64")]
pub fn indexer() -> ContainerRequest<GenericImage> {
    GenericImage::new(
        "ghcr.io/nuffle-labs/nffl/indexer",
        "66dcb37e32e34f552a63c1e638a57dd251846f63",
    )
        .with_cmd(vec![
            "--rmq-address",
            "amqp://rmq:5672",
            "--da-contract-ids",
            "da2.test.near",
            "--da-contract-ids",
            "da3.test.near",
            "--rollup-ids",
            "2",
            "--rollup-ids",
            "3",
        ])
        .with_network(NETWORK_NAME)
}

#[cfg(test)]
#[cfg(target_arch = "x86_64")]
pub fn aggregator() -> ContainerRequest<GenericImage> {
    GenericImage::new(
        "ghcr.io/nuffle-labs/nffl/aggregator",
        "66dcb37e32e34f552a63c1e638a57dd251846f63",
    )
        .with_copy_to("/nffl", PathBuf::from(CURRENT_DIR))
        .with_working_dir("/nffl")
        .with_cmd(vec![
            "--config",
            "config-files/aggregator-docker-compose.yaml",
            "--nffl-deployment",
            "contracts/evm/script/output/31337/sffl_avs_deployment_output.json",
            "--ecdsa-private-key",
            "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
        ])
        .with_network(NETWORK_NAME)
}

#[cfg(test)]
#[cfg(target_arch = "x86_64")]
pub fn operator(config_path: &str) -> ContainerRequest<GenericImage> {
    GenericImage::new(
        "ghcr.io/nuffle-labs/nffl/operator",
        "66dcb37e32e34f552a63c1e638a57dd251846f63",
    )
        .with_cmd(vec!["--config", config_path])
        .with_copy_to("/nffl", PathBuf::from(CURRENT_DIR))
        .with_working_dir("/nffl")
        .with_env_var("OPERATOR_ECDSA_KEY_PASSWORD", "EnJuncq01CiVk9UbuBYl")
        .with_env_var("OPERATOR_BLS_KEY_PASSWORD", "fDUMDLmBROwlzzPXyIcy")
        .with_network(NETWORK_NAME)
}
