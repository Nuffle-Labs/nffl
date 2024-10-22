use wiremock::{MockServer, Mock, ResponseTemplate};
use wiremock::matchers::{method, path};
use workers::verifier::Verifier;

#[tokio::main]
#[test]
async fn main() {
    // Start a background HTTP server on a random local port
    let mock_server = MockServer::start().await;
    setup(&mock_server).await;
    
    let verifier = Verifier::new("", "").await?;
    assert_eq!(verifier.verify(1, 2).await, true);
}


async fn setup(mock_server: &MockServer) -> eyre::Result<()> {
    let state_root_message = workers::verifier::Message {
        rollup_id: 1,
        block_height: 2,
        timestamp : 3,
        near_da_transaction_id : vec![4],
        near_da_commitment : vec![5],
        state_root: vec![6]
    };

    Mock::given(method("GET"))
        .and(path("/state-root"))
        .respond_with(ResponseTemplate::new(200).set_body_json(state_root_message))
        .mount(&mock_server)
        .await;
    
    Ok(())
}