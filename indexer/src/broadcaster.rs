use near_indexer::near_primitives::types::TransactionOrReceiptId;
use near_indexer::near_primitives::views::ReceiptView;
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;

pub(crate) async fn listen_tx_candidates(
    view_client: actix::Addr<near_client::ViewClientActor>,
    mut receiver: mpsc::Receiver<near_client::TxStatus>,
) {
    while let Some(tx_status) = receiver.recv().await {
        // TODO: handle errors
        let tx_status = view_client.send(tx_status.with_span_context()).await.unwrap().unwrap();

        // let execution_outcome = if let Some(execution_outcome) = tx_status.execution_outcome {
        //     return execution_outcome
        // }

        println!(
            "execution_outcome: {:?} \n status {:?}",
            tx_status.execution_outcome, tx_status.status
        );
    }
}

pub(crate) async fn listen_execution_outcomes(
    view_client: actix::Addr<near_client::ViewClientActor>,
    mut receiver: mpsc::Receiver<near_client::GetExecutionOutcome>,
) {
    while let Some(execution_outcome) = receiver.recv().await {
        // TODO: handle errors
        let execution_outcome = view_client
            .send(execution_outcome.with_span_context())
            .await
            .unwrap()
            .unwrap();
        println!("listen_execution_outcomes {:?}", execution_outcome.outcome_proof);
    }
}
