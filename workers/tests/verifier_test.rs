use wiremock::matchers::{method, path};
use wiremock::{Mock, MockServer, ResponseTemplate};
use workers::verifier::NFFLVerifier;

#[tokio::main]
#[test]
async fn main() -> eyre::Result<()> {
    // Start a background HTTP server on a random local port
    let mock_server = MockServer::start().await;
    setup(&mock_server).await;

    let verifier = NFFLVerifier::new(
        "", 
        mock_server.address().to_string().as_str(),
        1,
    ).await?;
    let verify_result = verifier.verify(vec![].as_slice(),2).await;
    assert_eq!(verify_result.is_ok(), true);
    assert_eq!(verify_result.unwrap(), true);

    Ok(())
}

async fn setup(mock_server: &MockServer) -> eyre::Result<()> {
    let state_root_message = workers::verifier::Message {
        rollup_id: 1,
        block_height: 2,
        timestamp: 3,
        near_da_transaction_id: vec![4],
        near_da_commitment: vec![5],
        state_root: vec![0],
    };

    Mock::given(method("GET"))
        .and(path("/state-root"))
        .respond_with(ResponseTemplate::new(200).set_body_json(state_root_message))
        .mount(&mock_server)
        .await;

    Ok(())
}
