use crate::errors::Result;
use crate::rabbit_publisher::RabbitPublisher;
use std::ops::Deref;

use near_indexer::near_primitives::types::{FunctionArgs, TransactionOrReceiptId};
use near_indexer::near_primitives::views::{ActionView, ExecutionStatusView, ReceiptEnumView, ReceiptView};
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;

// pub(crate) struct CandidateData {
//     pub execution_request: near_client::GetExecutionOutcome,
//     pub
// }

pub(crate) async fn process_receipt_candidates(
    view_client: actix::Addr<near_client::ViewClientActor>,
    // TODO: supply only needed data: receiver_id, receipt_id and ActionView::FunctionCall
    mut receiver: mpsc::Receiver<ReceiptView>,
    mut rabbit_publisher: RabbitPublisher,
) -> Result<()> {
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

        let payloads = if let ExecutionStatusView::SuccessValue(_) = execution_outcome.outcome_proof.outcome.status {
            if let ReceiptEnumView::Action { actions, .. } = &receipt.receipt {
                actions
                    .iter()
                    .filter_map(|el| match &el {
                        ActionView::FunctionCall { method_name, args, .. } if method_name == "submit" => Some(args),
                        _ => None,
                    })
                    .collect::<Vec<&FunctionArgs>>()
            } else {
                unreachable!();
            }
        } else {
            continue;
        };

        // TODO: is sequential order important here?
        for el in payloads {
            match rabbit_publisher.publish(el.deref()).await {
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
