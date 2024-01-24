use near_indexer::near_primitives::views::ExecutionStatusView;
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;
use tracing::info;

use crate::block_listener::CandidateData;
use crate::rabbit_publisher::{PublishData, PublishOptions, PublisherContext};
use crate::{errors::Result, rabbit_publisher::RabbitPublisher};

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

    pub(crate) async fn start(self) -> Result<()> {
        let Self {
            mut receiver,
            mut rabbit_publisher,
            view_client,
        } = self;

        let res = 'main: loop {
            let candidate_data = if let Some(candidate_data) = receiver.recv().await {
                candidate_data
            } else {
                break Ok(());
            };

            let execution_outcome = match view_client
                .send(
                    near_client::GetExecutionOutcome {
                        id: candidate_data.transaction_or_receipt_id.clone(),
                    }
                    .with_span_context(),
                )
                .await
            {
                Ok(Ok(response)) => response,
                Ok(Err(err)) => break Err(err.into()),
                Err(err) => break Err(err.into()),
            };

            if !matches!(
                execution_outcome.outcome_proof.outcome.status,
                ExecutionStatusView::SuccessValue(_)
            ) {
                info!(target: "candidates_validator", "unsuccessful value");
                continue;
            };

            // TODO: is sequential order important here?
            for payload in candidate_data.payloads {
                if let Err(err) = rabbit_publisher
                    .publish(PublishData {
                        publish_options: PublishOptions::default(),
                        cx: PublisherContext {
                            block_hash: execution_outcome.outcome_proof.block_hash,
                            id: candidate_data.transaction_or_receipt_id.clone(),
                        },
                        payload,
                    })
                    .await
                {
                    break 'main Err(err);
                };
            }
        };

        receiver.close();

        // Drain messages
        while let Some(_) = receiver.recv().await {}

        res
    }
}
