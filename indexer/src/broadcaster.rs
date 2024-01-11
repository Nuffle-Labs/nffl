use crate::errors::Result;
use crate::rabbit_publisher::RabbitPublisher;
use std::ops::Deref;

use near_indexer::near_primitives::types::{FunctionArgs, TransactionOrReceiptId};
use near_indexer::near_primitives::views::{ActionView, ExecutionStatusView, ReceiptEnumView, ReceiptView};
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;

pub(crate) struct CandidateData {
    pub execution_request: near_client::GetExecutionOutcome,
    pub payloads: Vec<Vec<u8>>,
}

pub(crate) async fn process_receipt_candidates(
    view_client: actix::Addr<near_client::ViewClientActor>,
    mut receiver: mpsc::Receiver<CandidateData>,
    mut rabbit_publisher: RabbitPublisher,
) -> Result<()> {
    'main: while let Some(receipt) = receiver.recv().await {
        let execution_outcome = match view_client.send(receipt.execution_request.with_span_context()).await {
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

        if !matches!(
            execution_outcome.outcome_proof.outcome.status,
            ExecutionStatusView::SuccessValue(_)
        ) {
            // TODO: log this?
            continue;
        };

        // TODO: is sequential order important here?
        for el in receipt.payloads {
            match rabbit_publisher.publish(&el).await {
                Ok(_) => (),
                Err(err) => {
                    eprintln!("{}", err);
                    receiver.close();
                    break 'main;
                }
            };
        }
    }

    // Drain messages
    while let Some(_) = receiver.recv().await {}

    Ok(())
}
