use alloy::providers::{ProviderBuilder, RootProvider, WsConnect};
use alloy::pubsub::PubSubFrontend;
use reqwest::{ClientBuilder, Url};
use serde::{Deserialize, Serialize};
use std::sync::Arc;

#[derive(Serialize, Deserialize, std::fmt::Debug)]
pub (crate) struct ResponseWrapper {
    pub message: Message,
}

#[derive(Serialize, Deserialize, std::fmt::Debug)]
pub struct Message {
    #[serde(rename = "RollupId")]
    pub rollup_id: i64,
    #[serde(rename = "BlockHeight")]
    pub block_height: i64,
    #[serde(rename = "Timestamp")]
    pub timestamp: i64,
    #[serde(rename = "NearDaTransactionId")]
    pub near_da_transaction_id: Vec<i64>,
    #[serde(rename = "NearDaCommitment")]
    pub near_da_commitment: Vec<i64>,
    #[serde(rename = "StateRoot")]
    pub state_root: Vec<i64>,
}

pub struct Verifier {
    eth_l2_provider: Arc<RootProvider<PubSubFrontend>>,
    http_client: reqwest::Client,
    aggregator_http_address: String,
}

impl Verifier {
    pub async fn new(agg_url: &str, rpc_url: &str) -> eyre::Result<Verifier> {
        let ws = WsConnect::new(rpc_url);
        let provider = ProviderBuilder::new().on_ws(ws).await?;
        let client = ClientBuilder::new().build()?;

        let agg_http_addr = format!("{}/aggregation/state-root-update", agg_url);

        Ok(Verifier {
            eth_l2_provider: Arc::new(provider),
            http_client: client,
            aggregator_http_address: agg_http_addr,
        })
    }

    pub async fn verify(&self, rollup_id: u32, block_height: u64) -> bool {
        // tokio::spawn(async {
        let params = [
            ("rollupId", rollup_id.to_string()),
            ("blockHeight", block_height.to_string()),
        ];
        let url = Url::parse_with_params(&self.aggregator_http_address, &params)?;
        let res = self.http_client.get(url).send().await?;
        let response_payload = res.json::<ResponseWrapper>().await?;
        let message = &response_payload.message;
        // TODO: take the state root from the LayerZero message.
        message.state_root.eq(&vec![0])
    }
}

/*
Expected JSON successful response body.
{
  "Message" : {
    "RollupId" : 1,
    "BlockHeight" : 2,
    "Timestamp" : 3,
    "NearDaTransactionId" : [ ..bytes.. ],
    "NearDaCommitment" : [ ..bytes.. ],
    "StateRoot" : [ ..bytes.. ]
  }
}
*/
