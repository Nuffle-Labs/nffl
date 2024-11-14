use alloy::{
    primitives::{address, Address, Bytes, FixedBytes},
    providers::ProviderBuilder,
};
use axum::{routing::post, Json, Router};
use offchain::{
    abi::L0V2EndpointAbi::{Origin, PacketSent, PacketVerified},
    chain::{connections::get_abi_from_path, contracts::create_contract_instance, ContractInst},
    workers::executor::NFFLExecutor,
};
use std::{
    sync::atomic::{AtomicI32, Ordering},
    sync::Arc,
};
use tokio::task::JoinHandle;
use tracing::{debug, level_filters::LevelFilter};
use tracing_subscriber::EnvFilter;

#[derive(serde::Serialize, serde::Deserialize, Debug)]
struct EthCallRequest {
    method: String,
    params: Vec<serde_json::Value>,
    id: u32,
    jsonrpc: String,
}

#[derive(serde::Serialize, serde::Deserialize, Debug)]
struct EthCallResponse {
    result: String,
    id: u32,
    jsonrpc: String,
}

#[tokio::test]
async fn test_handle_verified_packet_success() -> eyre::Result<()> {
    // Initialize tracing
    tracing_subscriber::fmt()
        .with_target(true)
        .with_env_filter(
            EnvFilter::builder()
                .with_default_directive(LevelFilter::DEBUG.into())
                .from_env_lossy(),
        )
        .init();

    let counter: Arc<AtomicI32> = Arc::new(AtomicI32::new(0));

    let packet_sent = PacketSent {
        encodedPayload: Bytes::from(&[1; 256]),
        options: Bytes::from(&[1; 32]),
        sendLibrary: Address::from_slice(&[2; 20]),
    };

    let verified_packet = PacketVerified {
        origin: Origin {
            srcEid: 1,
            sender: FixedBytes::from(&[1; 32]),
            nonce: 101010,
        },
        receiver: Address::from_slice(&[1; 20]),
        payloadHash: FixedBytes::from(&[2; 32]),
    };

    let _join_handle = prepare_server(counter.clone()).await;
    let contract = setup_contract().await?;

    NFFLExecutor::handle_verified_packet(&contract, packet_sent, verified_packet).await?;

    assert_eq!(counter.load(Ordering::Acquire), 2);
    Ok(())
}

async fn prepare_server(counter: Arc<AtomicI32>) -> JoinHandle<()> {
    const SERVER_ADDRESS_SHORT: &str = "127.0.0.1:8081";

    // Define the handler for the POST request.
    let app = Router::new().route(
        "/",
        post(|| async move {
            debug!("Server : POST request accepted");
            counter.fetch_add(1, Ordering::Release);
            Json(EthCallResponse {
                result: "0x0000000000000000000000000000000000000000000000000000000000000002".to_string(),
                id: 1,
                jsonrpc: "2.0".to_string(),
            })
        }),
    );
    // Spawn the server on a background task.
    let listener = tokio::net::TcpListener::bind(SERVER_ADDRESS_SHORT).await.unwrap();

    tokio::spawn(async move { axum::serve(listener, app).await.unwrap() })
}

async fn setup_contract() -> eyre::Result<ContractInst> {
    const SERVER_ADDRESS: &str = "http://127.0.0.1:8081";

    let http_provider = ProviderBuilder::new().on_http(SERVER_ADDRESS.parse()?);
    let l0_abi = get_abi_from_path("offchain/abi/L0V2Endpoint.json")?;

    create_contract_instance(
        address!("d8da6bf26964af9d7eed9e03e53415d37aa96045"),
        http_provider,
        l0_abi,
    )
}