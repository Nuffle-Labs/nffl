use crate::config::WorkerConfig;
use alloy::eips::BlockNumberOrTag;
use alloy::network::Ethereum;
use alloy::primitives::B256;
use alloy::providers::{Provider, ProviderBuilder, ReqwestProvider};
use reqwest::{Client, ClientBuilder, Url};
use serde::{Deserialize, Serialize};
use std::time::Duration;
use tracing::error;

#[derive(Serialize, Deserialize, std::fmt::Debug)]
pub(crate) struct Message {
    #[serde(rename = "RollupId")]
    rollup_id: u32,
    #[serde(rename = "BlockHeight")]
    block_height: u64,
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
    l2_provider: ReqwestProvider<Ethereum>,
    aggregator_http_address: String,
    network_id: String,
}

impl NFFLVerifier {
    pub(crate) async fn new(agg_url: &str, eth_l2_url: &str, network_id: u64) -> eyre::Result<NFFLVerifier> {
        let client = ClientBuilder::new().build()?;
        let url: Url = eth_l2_url.parse()?;
        let provider = ProviderBuilder::new().on_http(url);
        let agg_http_addr = format!("{}/aggregation/state-root-update", agg_url);

        Ok(NFFLVerifier {
            l2_provider: provider,
            http_client: client,
            aggregator_http_address: agg_http_addr,
            network_id: network_id.to_string(),
        })
    }

    pub async fn new_from_config(cfg: &WorkerConfig) -> eyre::Result<NFFLVerifier> {
        Self::new(&cfg.aggregator_url, &cfg.http_rpc_url, cfg.target_network_eid).await
    }

    /// Verifies the state root of a block. In case any request future
    /// is interrupted, or finishes unsuccessfully, returns Ok(false).
    pub async fn verify(&self, block_height: u64) -> eyre::Result<bool> {
        const TIMEOUT: Duration = Duration::from_secs(10);
        match tokio::try_join!(
            tokio::time::timeout(TIMEOUT, self.get_aggregator_state_root(block_height)),
            tokio::time::timeout(TIMEOUT, self.get_block_state_root(block_height)),
        ) {
            Ok((Ok(agg_response), Ok(block_state_root))) => {
                let state_root_slice: &[u8] = agg_response.state_root.as_slice();
                let aggregator_state_root: B256 = B256::from_slice(state_root_slice);
                if agg_response.block_height != block_height {
                    error!(
                        "Block heights are not equal for comparing state roots.\
                     Aggregator height: {:?}, L2-provided height: {:?}",
                        agg_response.block_height, block_height
                    );
                    return Ok(false);
                }
                let comparison_result = block_state_root.eq(&aggregator_state_root);
                if !comparison_result {
                    error!(
                        "State roots are not equal. Aggregator: {:?}, L2-provided: {:?}",
                        aggregator_state_root, block_state_root
                    );
                }
                Ok(comparison_result)
            }
            Err(e) => {
                error!("Timeout error while verifying state root: {:?}", e);
                Ok(false)
            }
            _ => Ok(false),
        }
    }

    /// Fetches the root state from the NFFL aggregator via HTTP.
    pub(crate) async fn get_aggregator_state_root(&self, block_height: u64) -> eyre::Result<Message> {
        let params = [
            ("rollupId", &self.network_id),
            ("blockHeight", &block_height.to_string()),
        ];
        let url = Url::parse_with_params(&self.aggregator_http_address, &params)?;
        let response = self.http_client.get(url).send().await?;
        let mut result: serde_json::Value = response.json().await?;

        serde_json::from_value(result["Message"].take()).map_err(Into::into)
    }

    /// Fetches the block state root from the Ethereum L2 provider
    /// via JSON-RPC API, backed by alloy-rs.
    pub(crate) async fn get_block_state_root(&self, block_number: u64) -> eyre::Result<B256> {
        let b_number = BlockNumberOrTag::from(block_number);
        match self.l2_provider.get_block_by_number(b_number, true).await? {
            Some(block) => Ok(block.header.state_root),
            None => Err(eyre::eyre!("Block {block_number} not found")),
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::verifier::NFFLVerifier;
    use std::time::Duration;
    use wiremock::matchers::{method, path, query_param_contains};
    use wiremock::{Mock, MockServer, ResponseTemplate};

    #[tokio::test]
    async fn test_verify_mock_ok() {
        let mock_server = MockServer::start().await;
        setup(&mock_server, true).await;

        let verifier_result = NFFLVerifier::new(mock_server.uri().as_str(), mock_server.uri().as_str(), 1).await;

        assert!(verifier_result.is_ok());

        let verifier = verifier_result.unwrap();
        assert!(verifier.verify(2).await.unwrap());
    }

    #[tokio::test]
    async fn test_verify_mock_fails() {
        let mock_server = MockServer::start().await;
        setup(&mock_server, false).await;

        let verifier_result = NFFLVerifier::new(mock_server.uri().as_str(), mock_server.uri().as_str(), 1).await;

        assert!(verifier_result.is_ok());

        let verifier = verifier_result.unwrap();
        assert!(!verifier.verify(2).await.unwrap());
    }

    #[tokio::test]
    async fn test_verify_timeout_fail() {
        let mock_server = MockServer::start().await;

        // Mock that delays longer than timeout
        Mock::given(method("GET"))
            .respond_with(ResponseTemplate::new(200).set_delay(Duration::from_secs(11)))
            .mount(&mock_server)
            .await;

        let verifier = NFFLVerifier::new(mock_server.uri().as_str(), mock_server.uri().as_str(), 1)
            .await
            .unwrap();

        assert!(!verifier.verify(2).await.unwrap());
    }

    #[tokio::test]
    async fn test_aggregator_state_root_mock_ok() {
        let mock_server = MockServer::start().await;
        setup(&mock_server, true).await;

        let verifier_result = NFFLVerifier::new(mock_server.uri().as_str(), mock_server.uri().as_str(), 1).await;

        assert!(verifier_result.is_ok());

        let verifier = verifier_result.unwrap();

        let state_root_resp_res = verifier.get_aggregator_state_root(2).await;
        assert!(state_root_resp_res.is_ok());

        let state_root = state_root_resp_res.unwrap().state_root;
        assert_eq!(
            state_root,
            vec![
                99, 80, 208, 69, 66, 69, 251, 65, 15, 192, 251, 147, 246, 100, 140, 91, 144, 71, 166, 8, 20, 65, 227,
                111, 15, 243, 171, 37, 156, 154, 71, 240
            ]
        );
    }

    #[tokio::test]
    async fn test_block_state_root_mock_ok() {
        let mock_server = MockServer::start().await;
        setup(&mock_server, true).await;

        let verifier_result = NFFLVerifier::new(mock_server.uri().as_str(), mock_server.uri().as_str(), 1).await;

        assert!(verifier_result.is_ok());

        let verifier = verifier_result.unwrap();

        let state_root_resp_res = verifier.get_block_state_root(2).await;
        assert!(state_root_resp_res.is_ok());

        let state_root = state_root_resp_res.unwrap().to_vec();
        assert_eq!(
            state_root,
            vec![
                99, 80, 208, 69, 66, 69, 251, 65, 15, 192, 251, 147, 246, 100, 140, 91, 144, 71, 166, 8, 20, 65, 227,
                111, 15, 243, 171, 37, 156, 154, 71, 240
            ]
        );
    }

    async fn setup(mock_server: &MockServer, should_pass: bool) {
        let expected_state_root: Vec<u8> = match should_pass {
            true => vec![
                99, 80, 208, 69, 66, 69, 251, 65, 15, 192, 251, 147, 246, 100, 140, 91, 144, 71, 166, 8, 20, 65, 227,
                111, 15, 243, 171, 37, 156, 154, 71, 240,
            ],
            false => vec![0; 32],
        };

        Mock::given(method("GET"))
            .and(path("/aggregation/state-root-update"))
            .and(query_param_contains("blockHeight", "2"))
            .respond_with(ResponseTemplate::new(200).set_body_json(serde_json::json!(
                            {
              "Message" : {
                "RollupId" : 1,
                "BlockHeight" : 2,
                "Timestamp" : 3,
                "NearDaTransactionId" : [ 4 ],
                "NearDaCommitment" : [ 5 ],
                "StateRoot" : expected_state_root,
              },
              "Aggregation" : {
                "EthBlockNumber" : 0,
                "MessageDigest" : [ 141 ],
                "NonSignersPubkeysG1" : null,
                "QuorumApksG1" : null,
                "SignersApkG2" : null,
                "SignersAggSigG1" : null,
                "NonSignerQuorumBitmapIndices" : null,
                "QuorumApkIndices" : null,
                "TotalStakeIndices" : null,
                "NonSignerStakeIndices" : null
              }
            })))
            .mount(mock_server)
            .await;

        Mock::given(method("POST"))
            .and(path("/"))
            .respond_with(ResponseTemplate::new(200).set_body_json(serde_json::json!({
              "jsonrpc": "2.0",
              "id": 1,
              "result": {
                "difficulty": "0x1913ff69551dac",
                "extraData": "0xe4b883e5bda9e7a59ee4bb99e9b1bc000921",
                "gasLimit": "0xe4e1b2",
                "gasUsed": "0xe4d737",
                "hash": "0xa917fcc721a5465a484e9be17cda0cc5493933dd3bc70c9adbee192cb419c9d7",
                "logsBloom": "0x00af00124b82093253a6960ab5a003170000318c0a00c18d418505009c10c905810e05d4a4511044b6245a062122010233958626c80039250781851410a468418101040c0100f178088a4e89000140e00001880c1c601413ac47bc5882854701180b9404422202202521584000808843030a552488a80e60c804c8d8004d0480422585320e068028d2e190508130022600024a51c116151a07612040081000088ba5c891064920a846b36288a40280820212b20940280056b233060818988945f33460426105024024040923447ad1102000028b8f0e001e810021031840a2801831a0113b003a5485843004c10c4c10d6a04060a84d88500038ab10875a382c",
                "miner": "0x829bd824b016326a401d083b33d092293333a830",
                "mixHash": "0x7d416c4a24dc3b43898040ea788922d8563d44a5193e6c4a1d9c70990775c879",
                "nonce": "0xe6e41732385c71d6",
                "number": "0xc5043f",
                "parentHash": "0xd1c4628a6710d8dec345e5bca6b8093abf3f830516e05e36f419f993334d10ef",
                "receiptsRoot": "0x7eadd994da137c7720fe2bf2935220409ed23a06ec6470ffd2d478e41af0255b",
                "sha3Uncles": "0x7d9ce61d799ddcb5dfe1644ec7224ae7018f24ecb682f077b4c477da192e8553",
                "size": "0xa244",
                "stateRoot": "0x6350d0454245fb410fc0fb93f6648c5b9047a6081441e36f0ff3ab259c9a47f0",
                "timestamp": "0x6100bc82",
                "totalDifficulty": "0x5f35fb5663cdc988403",
                "transactions": [
                  "0x3dc2776aa483c0eee09c2ccc654bf81dccebead40e9bb664289637bfb5e7e954"
                ],
                "transactionsRoot": "0xa17c2a87a6ff2fd790d517e48279e02f2e092a05309300c976363e47e0012672",
                "uncles": [
                  "0xd3946359c70281162cf00c8164d99ca14801e8008715cb1fad93b9cecaf9f7d8"
                ]
              }
            })))
            .mount(mock_server)
            .await;
    }
}
