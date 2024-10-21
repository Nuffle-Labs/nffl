use alloy::providers::{ProviderBuilder, RootProvider, WsConnect};
use alloy::pubsub::PubSubFrontend;
use reqwest::{ClientBuilder, Url};
use serde::{Deserialize, Serialize};
use std::sync::Arc;
use crate::config;
use crate::config::DVNConfig;

#[derive(Serialize, Deserialize, std::fmt::Debug)]
pub (crate) struct ResponseWrapper {
    pub message: Message,
}

#[derive(Serialize, Deserialize, std::fmt::Debug)]
pub struct Message {
    #[serde(rename = "RollupId")]
    pub rollup_id: u32,
    #[serde(rename = "BlockHeight")]
    pub block_height: u64,
    #[serde(rename = "Timestamp")]
    pub timestamp: u64,
    #[serde(rename = "NearDaTransactionId")]
    pub near_da_transaction_id: Vec<u8>,
    #[serde(rename = "NearDaCommitment")]
    pub near_da_commitment: Vec<u8>,
    #[serde(rename = "StateRoot")]
    pub state_root: Vec<u8>,
}

pub struct NFFLVerifier {
    // eth_l2_provider: Arc<RootProvider<PubSubFrontend>>,
    http_client: reqwest::Client,
    aggregator_http_address: String,
    network_id: String,
}

impl NFFLVerifier {
    pub async fn new(agg_url: &str, rpc_url: &str, network_id: u64) -> eyre::Result<NFFLVerifier> {
        // let ws = WsConnect::new(rpc_url);
        // let provider = ProviderBuilder::new().on_ws(ws).await?;
        let client = ClientBuilder::new().build()?;

        let agg_http_addr = format!("{}/aggregation/state-root-update", agg_url);

        Ok(NFFLVerifier {
            // eth_l2_provider: Arc::new(provider),
            http_client: client,
            aggregator_http_address: agg_http_addr,
            network_id: network_id.to_string()
        })
    }    
    
    pub async fn new_from_config(cfg: &DVNConfig) -> eyre::Result<NFFLVerifier> {
        // let ws = WsConnect::new(rpc_url);
        // let provider = ProviderBuilder::new().on_ws(ws).await?;
        let client = ClientBuilder::new().build()?;

        let agg_http_addr = format!("{}/aggregation/state-root", cfg.aggregator_url());

        Ok(NFFLVerifier {
            // eth_l2_provider: Arc::new(provider),
            http_client: client,
            aggregator_http_address: agg_http_addr,
            network_id: cfg.network_id().to_string()
        })
    }

    pub async fn verify(&self, payload_hash: &[u8], block_height: u64) -> eyre::Result<bool> {
        // tokio::spawn(async {
        let params = [
            ("rollupId", self.network_id.clone()),
            ("blockHeight", block_height.to_string()),
        ];
        let url = Url::parse_with_params(&self.aggregator_http_address, &params)?;
        let res = self.http_client.get(url).send().await?;
        let response_payload = res.json::<ResponseWrapper>().await?;
        let message = &response_payload.message;
        if (message.block_height != block_height) {
            return Ok(false);
        }
        // TODO: take the hash from the LayerZero message for Merkle proof verification.
        // Ok(message.state_root.eq(&vec![0]))
        Ok(true)
    }
}
