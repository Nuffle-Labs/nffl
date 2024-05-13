use near_indexer::near_primitives::views::{ExecutionStatusView, FinalExecutionStatus};
use near_o11y::WithSpanContextExt;
use prometheus::Registry;
use std::{collections::VecDeque, sync, time::Duration};
use tokio::{
    sync::{mpsc, oneshot, Mutex},
    time,
};
use tracing::info;

use crate::types::CandidateData;
use crate::{
    errors::Result,
    metrics::{make_candidates_validator_metrics, CandidatesListener, Metricable},
    rabbit_publisher::RabbitPublisherHandle,
    rabbit_publisher::{get_routing_key, PublishData, PublishOptions, PublishPayload, PublisherContext},
    types,
};

const CANDIDATES_VALIDATOR: &str = "candidates_validator";

#[derive(Clone)]
pub(crate) struct CandidatesValidator {
    view_client: actix::Addr<near_client::ViewClientActor>,
    listener: Option<CandidatesListener>,
}

impl CandidatesValidator {
    pub(crate) fn new(view_client: actix::Addr<near_client::ViewClientActor>) -> Self {
        Self {
            view_client,
            listener: None,
        }
    }

    async fn ticker(
        mut done: oneshot::Receiver<()>,
        queue_protected: types::ProtectedQueue<CandidateData>,
        mut rmq_handle: RabbitPublisherHandle,
        view_client: actix::Addr<near_client::ViewClientActor>,
        listener: Option<CandidatesListener>,
    ) {
        info!(target: CANDIDATES_VALIDATOR, "Starting ticker");
        let mut interval = time::interval(Duration::from_secs(2));

        loop {
            tokio::select! {
                _ = interval.tick() => {
                    let mut queue = queue_protected.lock().await;
                    let _ = Self::flush(&mut queue, &mut rmq_handle, &view_client, listener.clone()).await;
                    interval.reset();
                },
                _ = &mut done => {
                    return
                }
            }
        }
    }

    // Assumes queue is under mutex
    async fn flush(
        queue: &mut VecDeque<CandidateData>,
        rmq_handle: &mut RabbitPublisherHandle,
        view_client: &actix::Addr<near_client::ViewClientActor>,
        listener: Option<CandidatesListener>,
    ) -> Result<bool> {
        if queue.is_empty() {
            return Ok(true);
        }

        info!(target: CANDIDATES_VALIDATOR, "Flushing");
        while let Some(candidate_data) = queue.front() {
            let final_status = Self::fetch_execution_outcome(view_client, candidate_data).await?;
            match final_status {
                FinalExecutionStatus::NotStarted | FinalExecutionStatus::Started => {
                    info!(target: CANDIDATES_VALIDATOR, "Execution not finished, not sending to RabbitMQ");
                    return Ok(false);
                }
                FinalExecutionStatus::SuccessValue(_) => {
                    info!(target: CANDIDATES_VALIDATOR, "Candidate status now successful, candidate_data: {}", candidate_data);
                    listener.as_ref().map(|l| l.num_successful.inc());
                    Self::send(candidate_data, rmq_handle).await?;
                    queue.pop_front();
                }
                FinalExecutionStatus::Failure(_) => {
                    info!(target: CANDIDATES_VALIDATOR, "Execution failed, not sending to RabbitMQ");
                    listener.as_ref().map(|l| l.num_failed.inc());

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
        info!(target: CANDIDATES_VALIDATOR, "Fetching execution outcome for candidate data");
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

    async fn send(candidate_data: &CandidateData, rmq_handle: &mut RabbitPublisherHandle) -> Result<()> {
        // TODO: is sequential order important here?
        for data in candidate_data.clone().payloads {
            rmq_handle
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

    async fn process_candidates(
        self,
        mut receiver: mpsc::Receiver<CandidateData>,
        mut rmq_handle: RabbitPublisherHandle,
    ) -> Result<()> {
        let Self { view_client, listener } = self;

        let queue_protected = sync::Arc::new(Mutex::new(VecDeque::new()));
        let (done_sender, done_receiver) = oneshot::channel();
        actix::spawn(Self::ticker(
            done_receiver,
            queue_protected.clone(),
            rmq_handle.clone(),
            view_client.clone(),
            listener.clone(),
        ));

        while let Some(candidate_data) = receiver.recv().await {
            info!(target: CANDIDATES_VALIDATOR, "Received candidate data");

            {
                let mut queue = queue_protected.lock().await;
                // TODO(edwin): handle errors/ unwrap_or(false)?
                info!(target: CANDIDATES_VALIDATOR, "Flushing enqueued candidates");
                let flushed = Self::flush(&mut queue, &mut rmq_handle, &view_client, listener.clone()).await?;

                if !flushed {
                    info!(target: CANDIDATES_VALIDATOR, "Not flushed, so enqueuing candidate data");
                    queue.push_back(candidate_data);
                    continue;
                }
            }

            let final_status = match candidate_data.transaction.outcome.execution_outcome.outcome.status {
                ExecutionStatusView::Failure(_) => {
                    info!(target: CANDIDATES_VALIDATOR, "Execution failed, not sending to RabbitMQ");
                    listener.as_ref().map(|l| l.num_failed.inc());

                    continue;
                }
                _ => Self::fetch_execution_outcome(&view_client, &candidate_data).await?,
            };

            match final_status {
                FinalExecutionStatus::NotStarted | FinalExecutionStatus::Started => {
                    info!(target: CANDIDATES_VALIDATOR, "Execution not finished, pushing back to queue");
                    let mut queue = queue_protected.lock().await;
                    queue.push_back(candidate_data);
                }
                FinalExecutionStatus::SuccessValue(_) => {
                    info!(target: CANDIDATES_VALIDATOR, "Candidate executed successfully, candidate_data: {}", candidate_data);
                    listener.as_ref().map(|l| l.num_successful.inc());

                    Self::send(&candidate_data, &mut rmq_handle).await?;
                }
                FinalExecutionStatus::Failure(_) => {
                    info!(target: CANDIDATES_VALIDATOR, "Execution failed, not sending to RabbitMQ");
                    listener.as_ref().map(|l| l.num_failed.inc());
                }
            }
        }

        let _ = done_sender.send(());
        Ok(())
    }

    // TODO: JoinHandle or errC
    pub(crate) fn run(&self, candidates_receiver: mpsc::Receiver<CandidateData>) -> mpsc::Receiver<PublishData> {
        let (sender, receiver) = mpsc::channel(1000);
        actix::spawn(
            self.clone()
                .process_candidates(candidates_receiver, RabbitPublisherHandle { sender }),
        );

        receiver
    }
}

impl Metricable for CandidatesValidator {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()> {
        let listener = make_candidates_validator_metrics(registry)?;
        self.listener = Some(listener);

        Ok(())
    }
}
