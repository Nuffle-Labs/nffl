use near_indexer::near_primitives;
use near_indexer::near_primitives::views::ExecutionOutcomeWithIdView;
use near_indexer::near_primitives::{types::TransactionOrReceiptId, views::ExecutionStatusView};
use near_o11y::WithSpanContextExt;
use std::{collections::VecDeque, sync, time::Duration};
use tokio::{
    sync::{mpsc, Mutex},
    time,
};
use tracing::info;

use crate::block_listener::CandidateData;
use crate::rabbit_publisher::{get_routing_key, PublishData, PublishOptions, PublisherContext};
use crate::{errors::Result, rabbit_publisher::RabbitPublisher};

type ProtectedQueue = sync::Arc<Mutex<VecDeque<CandidateData>>>;
type ProtectedPublisher = sync::Arc<Mutex<RabbitPublisher>>;

pub(crate) struct CandidatesValidator {
    view_client: actix::Addr<near_client::ViewClientActor>,
    receiver: mpsc::Receiver<CandidateData>,
    rabbit_publisher: RabbitPublisher,
}

impl CandidatesValidator {
    pub(crate) fn new(
        view_client: actix::Addr<near_client::ViewClientActor>,
        receiver: mpsc::Receiver<CandidateData>,
        rabbit_publisher: RabbitPublisher,
    ) -> Self {
        Self {
            view_client,
            receiver,
            rabbit_publisher,
        }
    }

    async fn ticker(
        mut done: mpsc::Receiver<()>,
        queue_protected: ProtectedQueue,
        publisher_protected: ProtectedPublisher,
        view_client: actix::Addr<near_client::ViewClientActor>,
    ) {
        let mut interval = time::interval(Duration::from_secs(2));

        loop {
            tokio::select! {
                _ = interval.tick() => {
                    info!(target: "ticker", "trying to flush");
                    let mut queue = queue_protected.lock().await;
                    let _ = Self::flush(&mut queue, publisher_protected.clone(), &view_client).await;
                },
                _ = done.recv() => {
                    return
                }
            }
        }
    }

    // Assumes queue is under mutex
    async fn flush(
        queue: &mut VecDeque<CandidateData>,
        publisher_protected: ProtectedPublisher,
        view_client: &actix::Addr<near_client::ViewClientActor>,
    ) -> Result<bool> {
        while let Some(candidate_data) = queue.front() {
            let execution_outcome = Self::fetch_execution_outcome(view_client, candidate_data).await?;
            match execution_outcome.outcome.status {
                ExecutionStatusView::Unknown | ExecutionStatusView::SuccessReceiptId(_) => return Ok(false),
                ExecutionStatusView::SuccessValue(_) => {
                    Self::send(candidate_data, publisher_protected.clone(), execution_outcome).await?;
                    queue.pop_front();
                }
                ExecutionStatusView::Failure(_) => {
                    queue.pop_front();
                }
            };
        }

        Ok(true)
    }

    async fn fetch_execution_outcome(
        view_client: &actix::Addr<near_client::ViewClientActor>,
        candidate_data: &CandidateData,
    ) -> Result<ExecutionOutcomeWithIdView> {
        Ok(view_client
            .send(
                near_client::GetExecutionOutcome {
                    id: TransactionOrReceiptId::Transaction {
                        transaction_hash: candidate_data.transaction.transaction.hash,
                        sender_id: candidate_data.clone().transaction.transaction.signer_id,
                    },
                }
                .with_span_context(),
            )
            .await??
            .outcome_proof)
    }

    async fn send(
        candidate_data: &CandidateData,
        publisher_protected: ProtectedPublisher,
        execution_outcome: near_primitives::views::ExecutionOutcomeWithIdView,
    ) -> Result<()> {
        // TODO: is sequential order important here?
        let mut rabbit_publisher = publisher_protected.lock().await;
        for payload in candidate_data.clone().payloads {
            rabbit_publisher
                .publish(PublishData {
                    publish_options: PublishOptions {
                        routing_key: get_routing_key(candidate_data.rollup_id),
                        ..PublishOptions::default()
                    },
                    cx: PublisherContext {
                        block_hash: execution_outcome.block_hash,
                    },
                    payload,
                })
                .await?;
        }

        Ok(())
    }

    async fn process_candidates(self) -> Result<()> {
        let Self {
            mut receiver,
            rabbit_publisher,
            view_client,
        } = self;

        let queue_protected = sync::Arc::new(Mutex::new(VecDeque::new()));
        let publisher_protected = sync::Arc::new(Mutex::new(rabbit_publisher));

        let (done_sender, done_receiver) = mpsc::channel(1);
        let _ = Self::ticker(
            done_receiver,
            queue_protected.clone(),
            publisher_protected.clone(),
            view_client.clone(),
        );

        while let Some(candidate_data) = receiver.recv().await {
            {
                let mut queue = queue_protected.lock().await;
                // TODO(edwin): handle errors/ unwrap_or(false)?
                let flushed = Self::flush(&mut queue, publisher_protected.clone(), &view_client).await?;

                if !flushed {
                    queue.push_back(candidate_data);
                    continue;
                }
            }

            let execution_outcome = candidate_data.transaction.outcome.execution_outcome.clone();
            let current_execution_outcome = match execution_outcome.outcome.status {
                ExecutionStatusView::SuccessValue(_) => {
                    Self::send(&candidate_data, publisher_protected.clone(), execution_outcome).await?;
                    continue;
                }
                ExecutionStatusView::Failure(_) => continue,
                _ => Self::fetch_execution_outcome(&view_client, &candidate_data).await?,
            };

            match current_execution_outcome.outcome.status {
                ExecutionStatusView::Unknown | ExecutionStatusView::SuccessReceiptId(_) => {
                    let mut queue = queue_protected.lock().await;
                    queue.push_back(candidate_data);
                }
                ExecutionStatusView::SuccessValue(_) => {
                    Self::send(&candidate_data, publisher_protected.clone(), current_execution_outcome).await?;
                }
                ExecutionStatusView::Failure(_) => {}
            }
        }

        Ok(done_sender.send(()).await?)
    }

    pub(crate) async fn start(self) -> Result<()> {
        let rabbit_publisher = self.rabbit_publisher.clone();
        tokio::select! {
            result = self.process_candidates() => result,
            _ = rabbit_publisher.closed() => {
                Ok(())
            }
        }
    }
}
