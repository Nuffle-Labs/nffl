use std::ops::Deref;
use crate::errors::Result;

use deadpool_lapin::Pool;
use lapin::options::BasicPublishOptions;
use lapin::BasicProperties;
use near_indexer::near_primitives::types::{FunctionArgs, TransactionOrReceiptId};
use near_indexer::near_primitives::views::{ActionView, ExecutionStatusView, ReceiptEnumView, ReceiptView};
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;

pub(crate) async fn process_receipt_candidates(
    view_client: actix::Addr<near_client::ViewClientActor>,
    // TODO: supply only needed data: receiver_id, receipt_id and ActionView::FunctionCall
    mut receiver: mpsc::Receiver<ReceiptView>,
    connection_pool: Pool,
) -> Result<()> {
    'main: while let Some(receipt) = receiver.recv().await {
        println!("receipt: {:?}", serde_json::to_string(&receipt).unwrap());

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

        println!(
            "listen_execution_outcomes {:?}",
            serde_json::to_string(&execution_outcome.outcome_proof).unwrap()
        );

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

        // TODO: move to actor
        let connection = connection_pool.get().await?;
        let channel = connection.create_channel().await?;

        for payload in payloads {
            channel
                .basic_publish(
                    "",
                    "hello",
                    BasicPublishOptions::default(),
                    payload.deref(),
                    BasicProperties::default(),
                )
                .await?;
        }
    }

    // Drain messages
    while let Some(_) = receiver.recv().await {}

    Ok(())
}
