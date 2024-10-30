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
use alloy::primitives::U256;
use eyre::Result;
use futures::StreamExt;
use std::collections::VecDeque;
use std::time::Duration;
use tokio::time::sleep;
use tracing::{debug, error};

pub struct NFFLExecutor {
    config: DVNConfig,
    packet_queue: VecDeque<PacketSent>,
    finish: bool,
}

impl NFFLExecutor {
    pub(crate) const MAX_EXECUTE_ATTEMPTS: usize = 10;

    pub fn new(config: DVNConfig) -> Self {
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
        let (mut ps_stream, mut ef_stream, mut pv_stream) = build_executor_subscriptions(&self.config).await?;

        let http_provider = get_http_provider(&self.config)?;
        let l0_abi = get_abi_from_path("./abi/L0V2Endpoint.json")?;
        // Create a contract instance.
        let contract = create_contract_instance(self.config.l0_endpoint_addr, http_provider, l0_abi)?;

        loop {
            tokio::select! {
                Some(log) = ps_stream.next() => {
                    match log.log_decode::<PacketSent>() {
                        Ok(packet_log) => {
                            self.packet_queue.push_back(packet_log.data().clone());
                        },
                        Err(e) => { error!("Failed to decode PacketSent event: {:?}", e);}
                    }
                }
                Some(log) = ef_stream.next() => {
                    match log.log_decode::<ExecutorFeePaid>() {
                        Ok(executor_fee_log) => {
                            if self.packet_queue.is_empty() {
                                continue;
                            }

                            if !executor_fee_log.data().executor.eq(&self.config.dvn_addr)  {
                                self.packet_queue.clear();
                                continue;
                            }
                        },
                        Err(e) => { error!("Failed to decode ExecutorFeePaid event: {:?}", e);}
                    }
                },
                Some(log) = pv_stream.next() => {
                    match log.log_decode::<PacketVerified>() {
                        Ok(inner_log) => {
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
        if queue.is_empty() {
            return Ok(());
        }

        let packet_sent = queue.pop_front().unwrap();
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
        let not_executable = DynSolValue::Uint(U256::from(0), 8);
        let verified_not_executable = DynSolValue::Uint(U256::from(1), 8);
        let executable = DynSolValue::Uint(U256::from(2), 8);
        let executed = DynSolValue::Uint(U256::from(3), 8);
        loop {
            debug!("Attempt #{retry_count} to call 'executable'");
            if retry_count == Self::MAX_EXECUTE_ATTEMPTS {
                error!("Maximum retries reached while waiting for `Executable` state.");
                break;
            }
            let call_result = call_builder.call().await?;
            if call_result.len() != 1 {
                error!("`executable` function call returned invalid response.");
                break;
            }

            debug!("Call value {:?}", call_result);

            // Note: why not pattern matching here? Rust analyzer ranted on `executable` variable
            // in the pattern, so an author decided to make it via conditions.
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
