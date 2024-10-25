use crate::config::DVNConfig;
use alloy::eips::BlockNumberOrTag;
use alloy::network::Ethereum;
use alloy::primitives::B256;
use alloy::providers::{Provider, ProviderBuilder, ReqwestProvider};
use reqwest::{Client, ClientBuilder, Url};
use serde::{Deserialize, Serialize};
use std::time::Duration;
use tracing::error;

#[derive(Serialize, Deserialize, std::fmt::Debug)]
pub(crate) struct ResponseWrapper {
    pub message: Message,
}

#[derive(Serialize, Deserialize, std::fmt::Debug)]
pub(crate) struct Message {
    #[serde(rename = "RollupId")]
    rollup_id: u32,
    #[serde(rename = "BlockHeight")]
    block_height: u64,
    #[serde(rename = "Timestamp")]
    timestamp: u64,
    #[serde(rename = "NearDaTransactionId")]
    near_da_transaction_id: Vec<u8>,
    #[serde(rename = "NearDaCommitment")]
    near_da_commitment: Vec<u8>,
    #[serde(rename = "StateRoot")]
    pub(crate) state_root: Vec<u8>,
}

// TODO: Generify in a future for other networks, like Solana.
/// Verifies state roots for the Decentralized Verification Network (DVN).
///
/// This verifier implements the V1 verification algorithm which:
/// 1. Fetches the state root from the aggregator
/// 2. Retrieves the block state root using the chain RPC API
/// 3. Compares the two state roots to determine message validity
pub struct NFFLVerifier {
    http_client: Client,
    eth_l2_provider: ReqwestProvider<Ethereum>,
    aggregator_http_address: String,
    network_id: String,
}

impl NFFLVerifier {
    pub(crate) async fn new(agg_url: &str, eth_l2_url: &str, network_id: u64) -> eyre::Result<NFFLVerifier> {
        let client = ClientBuilder::new().build()?;
        let url: Url = eth_l2_url.parse()?;
        let provider = ProviderBuilder::new().on_http(url);
        let agg_http_addr = format!("{}/state-root", agg_url);

        Ok(NFFLVerifier {
            eth_l2_provider: provider,
            http_client: client,
            aggregator_http_address: agg_http_addr,
            network_id: network_id.to_string(),
        })
    }

    pub async fn new_from_config(cfg: &DVNConfig) -> eyre::Result<NFFLVerifier> {
        Self::new(&cfg.aggregator_url, &cfg.http_rpc_url, cfg.network_eid).await
    }

    /// Verifies the state root of a block. In case any request future
    /// is interrupted, or finishes unsuccessfully, returns Ok(false).
    pub async fn verify(&self, block_height: u64) -> eyre::Result<bool> {
        const TIMEOUT: Duration = Duration::from_secs(10);
        match tokio::try_join!(
            tokio::time::timeout(TIMEOUT, self.get_aggregator_root_state(block_height)),
            tokio::time::timeout(TIMEOUT, self.get_block_state_root(block_height)),
        ) {
            Ok((Ok(agg_response), Ok(block_state_root))) => {
                let state_root_slice: &[u8] = agg_response.message.state_root.as_slice();
                let aggregator_state_root: B256 = B256::from_slice(state_root_slice);
                if agg_response.message.block_height != block_height {
                    return Ok(false);
                }
                Ok(block_state_root.eq(&aggregator_state_root))
            }
            Err(e) => {
                error!("Error while verifying state root: {:?}", e);
                Ok(false)
            }
            _ => Ok(false),
        }
    }

    /// Fetches the root state from the NFFL aggregator via HTTP.
    pub(crate) async fn get_aggregator_root_state(&self, block_height: u64) -> eyre::Result<ResponseWrapper> {
        let params = [
            ("rollupId", self.network_id.clone()),
            ("blockHeight", block_height.to_string()),
        ];
        let url = Url::parse_with_params(&self.aggregator_http_address, &params)?;
        let response = self.http_client.get(url).send().await?;
        let result = response.json::<ResponseWrapper>().await?;
        Ok(result)
    }

    /// Fetches the block state root from the Ethereum L2 provider
    /// via JSON-RPC API, backed by alloy-rs.
    /// Note: an author didn't write a test for that, because he doesn't know how to mock RPC :(
    pub(crate) async fn get_block_state_root(&self, block_number: u64) -> eyre::Result<B256> {
        let b_number = BlockNumberOrTag::from(block_number);
        match self.eth_l2_provider.get_block_by_number(b_number, true).await? {
            Some(block) => Ok(block.header.state_root),
            None => Err(eyre::eyre!("Block {block_number} not found")),
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::verifier::{Message, NFFLVerifier, ResponseWrapper};
    use wiremock::matchers::{method, path, query_param_contains};
    use wiremock::{Mock, MockServer, ResponseTemplate};

    #[tokio::test]
    async fn test_aggregator_root_state_mock_ok() {
        let mock_server = MockServer::start().await;
        setup(&mock_server).await;

        let verifier_result = NFFLVerifier::new(mock_server.uri().as_str(), "https://arbitrum.drpc.org", 1).await;

        assert!(verifier_result.is_ok());

        let verifier = verifier_result.unwrap();

        let state_root_resp_res = verifier.get_aggregator_root_state(2).await;
        assert!(state_root_resp_res.is_ok());

        let state_root = state_root_resp_res.unwrap().message.state_root;
        assert_eq!(state_root, vec![1, 1, 1]);
    }

    async fn setup(mock_server: &MockServer) {
        let state_root_message = Message {
            rollup_id: 1,
            block_height: 2,
            timestamp: 3,
            near_da_transaction_id: vec![4],
            near_da_commitment: vec![5],
            state_root: vec![1, 1, 1],
        };

        let state_root_resp = ResponseWrapper {
            message: state_root_message,
        };

        Mock::given(method("GET"))
            .and(path("/state-root"))
            .and(query_param_contains("blockHeight", "2"))
            .respond_with(ResponseTemplate::new(200).set_body_json(state_root_resp))
            .mount(mock_server)
            .await;
    }
}
