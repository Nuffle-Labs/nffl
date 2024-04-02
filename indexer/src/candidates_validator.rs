use near_indexer::near_primitives::types::TransactionOrReceiptId;
use near_indexer::near_primitives::views::ExecutionStatusView;
use near_o11y::WithSpanContextExt;
use serde_json::de::Read;
use std::collections::VecDeque;
use std::sync;
use std::time::Duration;
use tokio::sync::Mutex;
use tokio::{sync::mpsc, time};
use tracing::info;

use crate::block_listener::CandidateData;
use crate::rabbit_publisher::{get_routing_key, PublishData, PublishOptions, PublisherContext};
use crate::{errors::Result, rabbit_publisher::RabbitPublisher};

type ProtectedQueue = sync::Arc<tokio::sync::Mutex<VecDeque<CandidateData>>>;
type ProtectedPublisher = sync::Arc<tokio::sync::Mutex<RabbitPublisher>>;

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

    async fn toilet(
        mut done: mpsc::Receiver<bool>,
        queue: ProtectedQueue,
        publisher: ProtectedPublisher,
        view_client: actix::Addr<near_client::ViewClientActor>,
    ) {
        let mut interval = time::interval(Duration::from_secs(2));

        loop {
            tokio::select! {
                _ = interval.tick() => {
                    let _ = Self::flush(queue.clone(), publisher.clone(), &view_client).await;
                },
                _ = done.recv() => {
                    return
                }
            }
        }
    }

    async fn flush(
        queue_protected: ProtectedQueue,
        protected_publisher: ProtectedPublisher,
        view_client: &actix::Addr<near_client::ViewClientActor>,
    ) -> Result<bool> {
        let mut queue = queue_protected.lock().await;
        while let Some(candidate_data) = queue.front() {
            let mut publisher = protected_publisher.lock().await;

            let execution_status = Self::check_execution_and_send(view_client, candidate_data, &mut publisher).await?;
            match execution_status {
                ExecutionStatusView::Unknown | ExecutionStatusView::SuccessReceiptId(_) => return Ok(false),
                ExecutionStatusView::Failure(_) | ExecutionStatusView::SuccessValue(_) => {
                    queue.pop_front();
                }
            };
        }

        Ok(true)
    }

    async fn check_execution_and_send(
        view_client: &actix::Addr<near_client::ViewClientActor>,
        candidate_data: &CandidateData,
        rabbit_publisher: &mut RabbitPublisher,
    ) -> Result<ExecutionStatusView> {
        let execution_outcome = view_client
            .send(
                near_client::GetExecutionOutcome {
                    id: TransactionOrReceiptId::Transaction {
                        transaction_hash: candidate_data.transaction.transaction.hash,
                        sender_id: candidate_data.clone().transaction.transaction.signer_id,
                    },
                }
                .with_span_context(),
            )
            .await??;

        let execution_status = execution_outcome.outcome_proof.outcome.status;
        if !matches!(execution_status, ExecutionStatusView::SuccessValue(_)) {
            info!(target: "candidates_validator", "unsuccessful value");
            return Ok(execution_status);
        };

        // TODO: is sequential order important here?
        for payload in candidate_data.clone().payloads {
            rabbit_publisher
                .publish(PublishData {
                    publish_options: PublishOptions {
                        routing_key: get_routing_key(candidate_data.rollup_id),
                        ..PublishOptions::default()
                    },
                    cx: PublisherContext {
                        block_hash: execution_outcome.outcome_proof.block_hash,
                    },
                    payload,
                })
                .await?;
        }

        Ok(execution_status)
    }

    async fn process_candidates(self) -> Result<()> {
        let Self {
            mut receiver,
            mut rabbit_publisher,
            view_client,
        } = self;

        let queue_protected = sync::Arc::new(Mutex::new(VecDeque::new()));
        let publisher_protected = sync::Arc::new(Mutex::new(rabbit_publisher));
        while let Some(candidate_data) = receiver.recv().await {
            let flushed = Self::flush(queue_protected.clone(), publisher_protected.clone(), &view_client).await?;
            if !flushed {
                let mut queue_protected = queue_protected.lock().await;
                queue_protected.push_back(candidate_data);

                continue;
            }

            let mut rabbit_publisher = publisher_protected.lock().await;
            let execution_status =
                Self::check_execution_and_send(&view_client, &candidate_data, &mut rabbit_publisher).await?;

            match execution_status {
                ExecutionStatusView::Unknown | ExecutionStatusView::SuccessReceiptId(_) => {
                    let mut queue = queue_protected.lock().await;
                    queue.push_back(candidate_data);
                }
                ExecutionStatusView::Failure(_) | ExecutionStatusView::SuccessValue(_) => {}
            }
        }

        Ok(())
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
