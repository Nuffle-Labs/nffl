use crate::abi::L0V2EndpointAbi::PacketSent;
use crate::abi::L0V2EndpointAbi::PacketVerified;
use crate::abi::SendLibraryAbi::ExecutorFeePaid;
use crate::chain::connections::build_executor_subscriptions;
use crate::chain::connections::get_abi_from_path;
use crate::chain::connections::get_http_provider;
use crate::chain::contracts::create_contract_instance;
use crate::chain::contracts::{lz_receive, prepare_header};
use crate::chain::ContractInst;
use crate::config::DVNConfig;
use alloy::dyn_abi::DynSolValue;
use alloy::primitives::I256;
use eyre::Result;
use futures::StreamExt;
use std::collections::VecDeque;
use std::time::Duration;
use tokio::time::sleep;
use tracing::error;

pub struct Executor {
    config: DVNConfig,
    finish: bool,
}

impl Executor {
    pub(crate) const NOT_EXECUTABLE: &'static DynSolValue = &DynSolValue::Int(I256::ZERO, 32);
    pub(crate) const VERIFIED_NOT_EXECUTABLE: &'static DynSolValue = &DynSolValue::Int(I256::ONE, 32);

    pub fn new(config: DVNConfig) -> Self {
        Executor { config, finish: false }
    }

    pub fn finish(&mut self) {
        // Note for myself: we are in a single-threaded event loop,
        // may not care about atomicity (yet?).
        self.finish = true;
    }

    #[deny(clippy::while_immutable_condition)]
    pub async fn listen(&self) -> Result<()> {
        let (mut ps_stream, mut ef_stream, mut pv_stream) = build_executor_subscriptions(&self.config).await?;

        let http_provider = get_http_provider(&self.config)?;
        let l0_abi = get_abi_from_path("./abi/L0V2Endpoint.json")?;
        // Create a contract instance.
        let contract = create_contract_instance(self.config.l0_endpoint_addr, http_provider, l0_abi)?;

        let mut packet_sent_queue: VecDeque<PacketSent> = VecDeque::default();

        loop {
            tokio::select! {
                Some(log) = ps_stream.next() => {
                    match log.log_decode::<PacketSent>() {
                        Ok(packet_log) => {
                            packet_sent_queue.push_back(packet_log.data().clone());
                        },
                        Err(e) => { error!("Failed to decode PacketSent event: {:?}", e);}
                    }
                }
                Some(log) = ef_stream.next() => {
                    match log.log_decode::<ExecutorFeePaid>() {
                        Ok(executor_fee_log) => {
                           if packet_sent_queue.is_empty() {
                                continue;
                            }
                            if !executor_fee_log.data().executor.eq(&self.config.dvn_addr)  {
                                packet_sent_queue.clear();
                                continue;
                            }
                        },
                        Err(e) => { error!("Failed to decode ExecutorFeePaid event: {:?}", e);}
                    }
                },
                    Some(log) = pv_stream.next() => {
                    match log.log_decode::<PacketVerified>() {
                        Ok(inner_log) => {
                           Self::handle_verified_packet(&contract, &mut packet_sent_queue, inner_log.data()).await?;
                        },
                        Err(e) => { error!("Failed to decode PacketVerified event: {:?}", e);}
                    }
                },
            }
            if self.finish {
                break;
            }
        }
        Ok(())
    }

    async fn handle_verified_packet(
        contract: &ContractInst,
        queue: &mut VecDeque<PacketSent>,
        packet_verified: &PacketVerified,
    ) -> Result<()> {
        if queue.is_empty() {
            return Ok(());
        }

        let packet_sent = queue.pop_front().unwrap();
        // We don't expect any item to be present. If we have any - it is garbage.
        queue.clear();
        // Despite being described with other arguments, the only real implementation in
        // the contract of the `executable` function located here: https://shorturl.at/4H6Yz
        // function executable(Origin memory _origin, address _receiver) returns (ExecutionState)
        let call_builder = contract.function(
            "executable",
            &[
                prepare_header(&packet_sent.encodedPayload[..]), // Origin (selected header fields)
                DynSolValue::Address(packet_verified.receiver),  // receiver address
            ],
        )?;

        // status `Executable` is represented by the integer 2 in the enum.
        // To read more: https://tinyurl.com/zur3btzs (line 9)
        let executable: &DynSolValue = &DynSolValue::Int(I256::unchecked_from(2), 32);
        loop {
            let call_result = call_builder.call().await?;
            if call_result.len() != 1 {
                error!("Failed to call executable function");
                break;
            }

            // Note: why not pattern matching here? Rust analyzer ranted
            // on `executable` variable in the pattern, so I decided to make it via conditions.
            if call_result[0].eq(Executor::NOT_EXECUTABLE) || call_result[0].eq(Executor::VERIFIED_NOT_EXECUTABLE) {
                // state: NotExecutable or VerifiedNotExecutable, await commits/verifications
                sleep(Duration::from_secs(1)).await;
                continue;
            } else if call_result[0].eq(executable) {
                // state: Executable, firing lz_receive
                lz_receive(contract, &packet_sent.encodedPayload[..]).await?;
                break;
            } else {
                // We may ignore Executed status, it just frees the executor.
                break;
            }
        }
        Ok(())
    }
}
