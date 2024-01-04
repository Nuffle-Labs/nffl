use near_indexer::near_primitives::types::TransactionOrReceiptId;
use near_indexer::near_primitives::views::ReceiptView;
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;

pub(crate) async fn listen_receipt_candidates(
    view_client: actix::Addr<near_client::ViewClientActor>,
    mut receiver: mpsc::Receiver<ReceiptView>,
) {
    'main: while let Some(receipt) = receiver.recv().await {
        let execution_outcome = match view_client
            .send(
                near_client::GetExecutionOutcome {
                    id: TransactionOrReceiptId::Receipt {
                        receipt_id: receipt.receipt_id,
                        receiver_id: receipt.receiver_id,
                    },
                }
                .with_span_context(),
            )
            .await
        {
            Ok(Ok(response)) => response,
            Ok(Err(err)) => {
                eprintln!("{}", err.to_string());
                receiver.close();
                break 'main;
            }
            Err(err) => {
                eprintln!("{}", err);
                receiver.close();
                break 'main;
            }
        };

        println!("listen_execution_outcomes {:?}", execution_outcome.outcome_proof);
    }

    // Drain messages
    while let Some(_) = receiver.recv().await {}
}
