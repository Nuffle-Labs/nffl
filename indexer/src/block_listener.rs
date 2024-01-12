use futures::future::join_all;
use near_indexer::near_primitives::{
    types::{AccountId, TransactionOrReceiptId},
    views::{ActionView, ReceiptEnumView},
};
use tokio::sync::mpsc;

use crate::errors::{Error, Result};

#[derive(Clone)]
pub(crate) struct CandidateData {
    pub transaction_or_receipt_id: TransactionOrReceiptId,
    pub payloads: Vec<Vec<u8>>,
}

pub(crate) struct BlockListener {
    stream: mpsc::Receiver<near_indexer::StreamerMessage>,
    receipt_sender: mpsc::Sender<CandidateData>,
    da_contract_id: AccountId,
}

impl BlockListener {
    pub(crate) fn new(
        stream: mpsc::Receiver<near_indexer::StreamerMessage>,
        receipt_sender: mpsc::Sender<CandidateData>,
        da_contract_id: AccountId,
    ) -> Self {
        Self {
            stream,
            receipt_sender,
            da_contract_id,
        }
    }

    pub(crate) async fn start(self) -> Result<()> {
        let Self {
            mut stream,
            receipt_sender,
            da_contract_id,
        } = self;

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
                            transaction_or_receipt_id: TransactionOrReceiptId::Receipt {
                                receipt_id: receipt.receipt_id,
                                receiver_id: receipt.receiver_id,
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
}
