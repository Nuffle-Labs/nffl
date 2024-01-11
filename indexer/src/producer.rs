use futures::future::join_all;
use near_indexer::near_primitives::{
    types::{AccountId, TransactionOrReceiptId},
    views::{ActionView, ExecutionStatusView, ReceiptEnumView},
};
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;

use crate::{
    errors::{Error, Result},
    rabbit_publisher::RabbitPublisher,
};

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

pub(crate) async fn listen_blocks(
    mut stream: mpsc::Receiver<near_indexer::StreamerMessage>,
    receipt_sender: mpsc::Sender<CandidateData>,
    da_contract_id: AccountId,
) -> Result<()> {
    while let Some(streamer_message) = stream.recv().await {
        let candidates_data: Vec<CandidateData> = streamer_message
            .shards
            .into_iter()
            .flat_map(|shard| shard.chunk)
            .flat_map(|chunk| {
                chunk.receipts.into_iter().filter_map(|receipt| {
                    if receipt.receiver_id != da_contract_id {
                        return None;
                    }

                    let actions = if let ReceiptEnumView::Action { actions, .. } = receipt.receipt {
                        actions
                    } else {
                        return None;
                    };

                    let payloads = actions
                        .into_iter()
                        .filter_map(|el| match el {
                            ActionView::FunctionCall { method_name, args, .. } if method_name == "submit" => {
                                Some(args.into())
                            }
                            _ => None,
                        })
                        .collect::<Vec<Vec<u8>>>();

                    if payloads.is_empty() {
                        return None;
                    }

                    Some(CandidateData {
                        execution_request: near_client::GetExecutionOutcome {
                            id: TransactionOrReceiptId::Receipt {
                                receipt_id: receipt.receipt_id,
                                receiver_id: receipt.receiver_id,
                            },
                        },
                        payloads,
                    })
                })
            })
            .collect();

        let results = join_all(candidates_data.into_iter().map(|receipt| receipt_sender.send(receipt))).await;

        // Receiver dropped or closed.
        if let Some(_) = results.iter().find_map(|result| result.as_ref().err()) {
            return Err(Error::SendError);
        }
    }

    Ok(())
}
