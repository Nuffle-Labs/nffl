use near_indexer::near_primitives::views::{ExecutionStatusView, FinalExecutionStatus};
use near_o11y::WithSpanContextExt;
use std::{collections::VecDeque, sync, time::Duration};
use tokio::{
    sync::{mpsc, Mutex},
    time,
};
use tracing::info;

use crate::{
    block_listener::CandidateData,
    errors::Result,
    rabbit_publisher::RabbitPublisher,
    rabbit_publisher::{get_routing_key, PublishData, PublishOptions, PublishPayload, PublisherContext},
};

const CANDIDATES_VALIDATOR: &str = "candidates_validator";

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
                    let mut queue = queue_protected.lock().await;
                    let _ = Self::flush(&mut queue, publisher_protected.clone(), &view_client).await;
                    interval.reset();
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
        if queue.is_empty() {
            return Ok(true);
        }

        info!(target: CANDIDATES_VALIDATOR, "Flushing");
        while let Some(candidate_data) = queue.front() {
            let final_status = Self::fetch_execution_outcome(view_client, candidate_data).await?;
            match final_status {
                FinalExecutionStatus::NotStarted | FinalExecutionStatus::Started => return Ok(false),
                FinalExecutionStatus::SuccessValue(_) => {
                    info!(target: CANDIDATES_VALIDATOR, "Candidate status now successful, candidate_data: {}", candidate_data);
                    Self::send(candidate_data, publisher_protected.clone()).await?;
                    queue.pop_front();
                }
                FinalExecutionStatus::Failure(_) => {
                    queue.pop_front();
                }
            };
        }

        Ok(true)
    }

    async fn fetch_execution_outcome(
        view_client: &actix::Addr<near_client::ViewClientActor>,
        candidate_data: &CandidateData,
    ) -> Result<FinalExecutionStatus> {
        Ok(view_client
            .send(
                near_client::TxStatus {
                    tx_hash: candidate_data.transaction.transaction.hash,
                    signer_account_id: candidate_data.transaction.transaction.signer_id.clone(),
                    fetch_receipt: true,
                }
                .with_span_context(),
            )
            .await??
            .execution_outcome
            .map(|x| x.into_outcome().status)
            .unwrap_or(FinalExecutionStatus::NotStarted))
    }

    async fn send(candidate_data: &CandidateData, publisher_protected: ProtectedPublisher) -> Result<()> {
        // TODO: is sequential order important here?
        let mut rabbit_publisher = publisher_protected.lock().await;
        for data in candidate_data.clone().payloads {
            rabbit_publisher
                .publish(PublishData {
                    publish_options: PublishOptions {
                        routing_key: get_routing_key(candidate_data.rollup_id),
                        ..PublishOptions::default()
                    },
                    cx: PublisherContext {
                        block_hash: candidate_data.transaction.outcome.execution_outcome.block_hash,
                    },
                    payload: PublishPayload {
                        transaction_id: candidate_data.transaction.transaction.hash,
                        data,
                    },
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
        tokio::spawn(Self::ticker(
            done_receiver,
            queue_protected.clone(),
            publisher_protected.clone(),
            view_client.clone(),
        ));

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

            let final_status = match candidate_data.transaction.outcome.execution_outcome.outcome.status {
                ExecutionStatusView::Failure(_) => continue,
                _ => Self::fetch_execution_outcome(&view_client, &candidate_data).await?,
            };

            match final_status {
                FinalExecutionStatus::NotStarted | FinalExecutionStatus::Started => {
                    let mut queue = queue_protected.lock().await;
                    queue.push_back(candidate_data);
                }
                FinalExecutionStatus::SuccessValue(_) => {
                    info!(target: CANDIDATES_VALIDATOR, "Candidate executed successfully, candidate_data: {}", candidate_data);
                    Self::send(&candidate_data, publisher_protected.clone()).await?;
                }
                FinalExecutionStatus::Failure(_) => {}
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
