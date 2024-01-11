use near_indexer::near_primitives::views::ExecutionStatusView;
use near_o11y::WithSpanContextExt;
use tokio::sync::mpsc;

use crate::block_listener::CandidateData;
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
}
