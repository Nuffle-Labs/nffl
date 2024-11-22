/// Docker container availability test for our test-containers suite
mod containers;

use testcontainers::runners::AsyncRunner;
use crate::containers::*;

#[cfg(test)]
#[tokio::test]
pub async fn check_rabbitmq_available() {
    let rabbitmq = rabbitmq();
    rabbitmq.start().await.expect("");
}

#[cfg(test)]
#[tokio::test]
pub async fn check_anvil_node_available() {
    let anvil_node = anvil_node();
    anvil_node.start().await.expect("");
}

#[cfg(test)]
#[tokio::test]
pub async fn check_anvil_node_setup_available() {
    let setup = anvil_node_setup();
    setup.start().await.expect("");
}

#[cfg(test)]
#[tokio::test]
pub async fn check_anvil_rollup_node_available() {
    let rollup_node = anvil_rollup_node(8546); // Example port number
    rollup_node.start().await.expect("");
}

#[cfg(test)]
#[tokio::test]
pub async fn check_near_da_deployer_available() {
    let deployer = near_da_deployer(3030); // Example indexer port number
    deployer.start().await.expect("");
}

#[cfg(all(test, target_arch = "x86_64"))]
#[tokio::test]
pub async fn check_rollup_relayer_available() {
    let relayer = rollup_relayer(8546); // Example rollup node port number
    relayer.start().await.expect("");
}

#[cfg(all(test, target_arch = "x86_64"))]
#[tokio::test]
pub async fn check_indexer_available() {
    let indexer = indexer();
    indexer.start().await.expect("");
}

#[cfg(all(test, target_arch = "x86_64"))]
#[tokio::test]
pub async fn check_aggregator_available() {
    let aggregator = aggregator();
    aggregator.start().await.expect("");
}

#[cfg(all(test, target_arch = "x86_64"))]
#[tokio::test]
pub async fn check_operator_available() {
    let operator = operator("../../../config-files/operator1-docker-compose.anvil.yaml"); // Example config path
    operator.start().await.expect("");
}
