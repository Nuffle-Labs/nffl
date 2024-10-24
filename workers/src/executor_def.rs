use crate::abi::L0V2EndpointAbi::PacketVerified;
use crate::abi::ReceiveLibraryAbi::PayloadVerified;
use crate::abi::SendLibraryAbi::ExecutorFeePaid;
use crate::chain::contracts::{executable, lz_receive, prepare_header};
use crate::chain::ContractInst;
use crate::config::DVNConfig;
use crate::config::LayerZeroEvent;
use crate::data::packet_v1_codec::receiver;
use crate::{
    abi::L0V2EndpointAbi::PacketSent,
    abi::SendLibraryAbi::DVNFeePaid,
    chain::{
        connections::{build_executor_subscriptions, get_abi_from_path, get_http_provider},
        contracts::{create_contract_instance, query_already_verified, query_confirmations, verify},
    },
};
use alloy::dyn_abi::DynSolValue;
use alloy::primitives::Address;
use alloy::primitives::{keccak256, I256};
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

        loop {
            let call_result = call_builder.call().await?;
            match call_result[0] {
                DynSolValue::Int(I256::ZERO, 32) => {
                    // state: NotExecutable, continue to await commits
                    sleep(Duration::from_millis(500)).await;
                    continue;
                }
                DynSolValue::Int(I256::ONE, 32) => {
                    // state: Executable, firing lz_receive
                    lz_receive(contract, &packet_sent.encodedPayload[..]).await?;
                    // TODO: handle the result of the lz_receive call
                    break;
                }
                // We may ignore Executed status, it just frees the executor.
                _ => break,
            };
        }
        Ok(())
    }
}
