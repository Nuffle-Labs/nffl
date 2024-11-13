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
use eyre::Result;
use futures::StreamExt;
use std::{collections::VecDeque, time::Duration};
use tokio::sync::mpsc;
use tokio::time::sleep;
use tracing::{debug, error, warn};

#[derive(Debug, Clone, Copy, PartialEq)]
pub enum ExecutionState {
    NotExecutable = 0,
    VerifiedNotExecutable = 1,
    Executable = 2,
    Executed = 3,
}

#[derive(Debug, Clone, PartialEq)]
pub enum ExecutorState {
    /// Initialized but not waiting for anything.
    Created,
    /// Listening for a `PacketSent` event to be executed.
    WaitingPacket,
    /// Listening for a `ExecutorFeePaid` event assigning the executor.
    WaitingAssignation,
    /// Listening for a `PacketVerified` event that triggers execution.
    WaitingVerification,
    /// Finished flow. Will resume listening again.
    Finish,
}

// NOTE: [IMPROVEMENT]: could also rewrite the executor with a BundleCreator, that packs the info from the events and then checks at the streams if there's more info to process the bundles.

pub struct NFFLExecutor {
    config: WorkerConfig,
    finish: bool,
    status: ExecutorState,
}

impl NFFLExecutor {
    pub(crate) const MAX_EXECUTE_ATTEMPTS: usize = 3;

    pub fn new(config: WorkerConfig) -> Self {
        NFFLExecutor { config, finish: false }
    }

    pub fn finish(&mut self) {
        self.finish = true;
    }

    pub async fn listen(&mut self) -> Result<()> {
        let (_sp, _tp, mut ps_stream, mut ef_stream, mut pv_stream) =
            build_executor_subscriptions(&self.config).await?;

        let http_provider = get_http_provider(&self.config.target_http_rpc_url)?;
        let l0_abi = get_abi_from_path("offchain/abi/L0V2Endpoint.json")?;
        // Create a contract instance.

        // Verified packet handler task
        let l0_addr = self.config.target_endpoint;
        let (tx, mut rx) = mpsc::channel::<(PacketSent, PacketVerified)>(4);
        tokio::spawn(async move {
            let contract = create_contract_instance(l0_addr, http_provider, l0_abi).unwrap();
            while let Some((packet_sent, packet_verified)) = rx.recv().await {
                debug!("Handler received PacketSent and PacketVerified");
                let _ = Self::handle_verified_packet(&contract, packet_sent, packet_verified).await;
            }
        });

        let mut packet_queue: VecDeque<PacketSent> = VecDeque::new();

        // Network I/O handler
        self.status = ExecutorState::WaitingPacket;
        loop {
            debug!("Iteration started, queue size {:?}", packet_queue.len());
            tokio::select! {
                Some(log) = ps_stream.next() => {
                    match log.log_decode::<PacketSent>() {
                        Ok(packet_log) => {
                            debug!("PacketSent received");
                            packet_queue.push_back(packet_log.data().clone());
                        },
                        Err(e) => { error!("Failed to decode PacketSent event: {:?}", e);}
                    }
                }
                Some(log) = ef_stream.next() => {
                    match log.log_decode::<ExecutorFeePaid>() {
                        Ok(executor_fee_log) => {
                            debug!("ExecutorFeePaid received");
                            if packet_queue.is_empty() {
                                continue;
                            }

                            debug!("{:?} ~ {:?}", executor_fee_log.data().executor, &self.config.executor);
                            // if !executor_fee_log.data().executor.eq(&self.config.executor)  {
                            // packet_queue.pop_front();
                            ////    self.packet_queue.clear();
                            //     continue;
                            // }
                        },
                        Err(e) => { error!("Failed to decode ExecutorFeePaid event: {:?}", e);}
                    }
                },
                Some(log) = pv_stream.next() => {
                    match log.log_decode::<PacketVerified>() {
                        Ok(inner_log) => {
                            if !packet_queue.is_empty() {
                                debug!("PacketSent and PacketVerified sent to handler");
                                tx.send((packet_queue.pop_front().unwrap(), inner_log.data().clone())).await?;
                            } else {
                                 warn!("PacketVerified event {:?} arrived for non-handled PacketSent", inner_log.log_index);
                            }
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

    pub async fn handle_verified_packet(
        contract: &ContractInst,
        packet_sent: PacketSent,
        packet_verified: PacketVerified,
    ) -> Result<()> {
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

            match call_builder.call().await {
                Ok(call_result) => {
                    debug!("> {:?}", call_result);
                    if call_result.len() != 1 {
                        error!("`executable` function call returned invalid response.");
                        break;
                    }
                    debug!(">> {:?}", call_result[0]);
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
                Err(e) => {
                    warn!("Failed to call `executable` function: {:?}", e);
                    sleep(Duration::from_secs(1)).await;
                    retry_count += 1;
                    continue;
                }
            }
        }
        Ok(())
    }
}
