use crate::{
    abi::{
        L0V2EndpointAbi::{PacketSent, PacketVerified},
        SendLibraryAbi::ExecutorFeePaid,
    },
    chain::{
        connections::{build_executor_subscriptions, get_abi_from_path, get_http_provider},
        contracts::{create_contract_instance, lz_receive, prepare_header},
        ContractInst,
    },
    config::WorkerConfig,
};
use alloy::{dyn_abi::DynSolValue, primitives::U256};
use eyre::{eyre, Result};
use futures::StreamExt;
use std::{collections::VecDeque, time::Duration};
use tokio::time::sleep;
use tracing::{debug, error};

#[derive(Debug, Clone, Copy, PartialEq)]
pub enum ExecutionState {
    NotExecutable = 0,
    VerifiedNotExecutable = 1,
    Executable = 2,
    Executed = 3,
}

pub struct NFFLExecutor {
    config: WorkerConfig,
    packet_queue: VecDeque<PacketSent>,
    finish: bool,
}

impl NFFLExecutor {
    pub(crate) const MAX_EXECUTE_ATTEMPTS: usize = 10;

    pub fn new(config: WorkerConfig) -> Self {
        NFFLExecutor {
            config,
            packet_queue: VecDeque::new(),
            finish: false,
        }
    }

    pub fn finish(&mut self) {
        // Note: we are in a single-threaded event loop, and we do not care about atomicity (yet).
        self.finish = true;
    }

    pub async fn listen(&mut self) -> Result<()> {
        let (_source_source_provider, _target_provider, mut ps_stream, mut ef_stream, mut pv_stream) =
            build_executor_subscriptions(&self.config).await?;

        // FIXME: should this be source or target?
        let http_provider = get_http_provider(&self.config.target_http_rpc_url)?;
        let l0_abi = get_abi_from_path("offchain/abi/L0V2Endpoint.json")?;
        // Create a contract instance.
        let contract = create_contract_instance(self.config.target_endpoint, http_provider, l0_abi)?;

        loop {
            tokio::select! {
                Some(log) = ps_stream.next() => {
                    match log.log_decode::<PacketSent>() {
                        Ok(packet_log) => {
                            debug!("Packet seemed to arrive");
                            self.packet_queue.push_back(packet_log.data().clone());
                        },
                        Err(e) => { error!("Failed to decode PacketSent event: {:?}", e);}
                    }
                }
                Some(log) = ef_stream.next() => {
                    match log.log_decode::<ExecutorFeePaid>() {
                        Ok(executor_fee_log) => {
                            debug!("Executor fee paid seemed to arrive");
                            if self.packet_queue.is_empty() {
                                continue;
                            }

                            //if !executor_fee_log.data().executor.eq(&self.config.source_dvn)  {
                            //    self.packet_queue.clear();
                            //    continue;
                            //}
                        },
                        Err(e) => { error!("Failed to decode ExecutorFeePaid event: {:?}", e);}
                    }
                },
                Some(log) = pv_stream.next() => {
                    // FIXME: we should not process everything, check for some condition to filter
                    match log.log_decode::<PacketVerified>() {
                        Ok(inner_log) => {
                            debug!("->> Packet verified seemd to arrived");
                           Self::handle_verified_packet(&contract, &mut self.packet_queue, inner_log.data()).await?;
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

    #[cfg(test)]
    pub fn is_queue_empty(&self) -> bool {
        self.packet_queue.is_empty()
    }

    pub async fn handle_verified_packet(
        contract: &ContractInst,
        queue: &mut VecDeque<PacketSent>,
        packet_verified: &PacketVerified,
    ) -> Result<()> {
        //if queue.is_empty() {
        //    return Ok(());
        //}
        //
        //let packet_sent = queue.pop_front().unwrap();
        let Some(packet_sent) = queue.pop_front() else {
            return Err(eyre!("Queue was empty, not handling packet verified"));
        };
        // We don't expect any more items to be present. If we have any - they are garbage/leftovers.
        queue.clear();
        // Despite being described with other arguments, the only real implementation of
        //  `executable` function is in the contract located here: https://shorturl.at/4H6Yz
        // function executable(Origin memory _origin, address _receiver) returns (ExecutionState)
        let call_builder = contract.function(
            "executable",
            &[
                prepare_header(&packet_sent.encodedPayload[..]),
                DynSolValue::Address(packet_verified.receiver),
            ],
        )?;

        let mut retry_count = 0;
        // status `Executable` is represented by the integer 2 in the enum.
        // To read more: https://tinyurl.com/zur3btzs (line 9)
        let not_executable = DynSolValue::Uint(U256::from(ExecutionState::NotExecutable as u8), 8);
        let verified_not_executable = DynSolValue::Uint(U256::from(ExecutionState::VerifiedNotExecutable as u8), 8);
        let executable = DynSolValue::Uint(U256::from(ExecutionState::Executable as u8), 8);
        let executed = DynSolValue::Uint(U256::from(ExecutionState::Executed as u8), 8);
        loop {
            debug!("Attempt #{retry_count} to call 'executable'");
            if retry_count == Self::MAX_EXECUTE_ATTEMPTS {
                error!("Maximum retries reached while waiting for `Executable` state.");
                break;
            }
            let call_result = call_builder.call().await?;
            debug!("Result of `executable` function call: {:?}", call_result);
            if call_result.len() != 1 {
                error!("`executable` function call returned invalid response.");
                break;
            }

            // Note: why not pattern matching here? fn calls are not allowed in patterns
            if call_result[0].eq(&not_executable) || call_result[0].eq(&verified_not_executable) {
                debug!("State: NotExecutable or VerifiedNotExecutable, await commits/verifications");
                sleep(Duration::from_secs(1)).await;
                retry_count += 1;
                continue;
            } else if call_result[0].eq(&executable) {
                debug!("State: Executable, fire and forget `lzReceive`");
                lz_receive(contract, &packet_sent.encodedPayload[..]).await?;
                break;
            } else if call_result[0].eq(&executed) {
                debug!("State: Executed, free the executor");
                break;
            } else {
                error!("Unknown state for `executable` call");
                break;
            }
        }
        Ok(())
    }
}
