#[derive(Clone, Deserialize)]
pub struct NodeConfig {
    pub production: bool,
    pub eth_rpc_url: String,
    pub eth_ws_url: String,
    pub operator_address: String,
    pub bls_private_key_store_path: String,
    pub ecdsa_private_key_store_path: String,
    pub avs_registry_coordinator_address: String,
    pub operator_state_retriever_address: String,
    pub aggregator_server_ip_port_address: String,
    pub enable_metrics: bool,
    pub eigen_metrics_ip_port_address: Option<String>,
    pub node_api_ip_port_address: Option<String>,
    pub task_response_wait_ms: u64,
    pub register_operator_on_startup: bool,
    pub token_strategy_addr: Option<String>,
}
