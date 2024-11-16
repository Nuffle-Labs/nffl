use crate::{
    abi::{L0V2EndpointAbi::PacketSent, SendLibraryAbi::DVNFeePaid},
    chain::{
        connections::{build_dvn_subscriptions, get_abi_from_path, get_http_provider},
        contracts::{commit, create_contract_instance, query_confirmations, verify},
        ContractInst,
    },
    config::WorkerConfig,
    data::packet_v1_codec::{header, message},
    verifier::NFFLVerifier,
};
use alloy::{
    primitives::{keccak256, B256, U256},
    rpc::types::Log,
};
use eyre::{eyre, Result};
use futures::StreamExt;
use std::collections::VecDeque;
use tracing::{debug, error, info};

const RECEIVELIB_ABI_PATH: &str = "offchain/abi/ReceiveLibUln302.json";

/// Entity in charge of running the DVN workflow.
pub struct Dvn {
    pub config: WorkerConfig,
    pub packet_queue: VecDeque<PacketSent>,
    pub target_receivelib: Option<ContractInst>,
    pub verifier: Option<NFFLVerifier>,
}

impl Dvn {
    /// Create a new DVN instance providing the configuration.
    pub fn new(config: WorkerConfig) -> Self {
        Self {
            config,
            packet_queue: VecDeque::new(),
            target_receivelib: None,
            verifier: None,
        }
    }

    /// Create a new DVN instance from the environment variables.
    pub fn new_from_env() -> Result<Self> {
        let conf = WorkerConfig::load_from_env()?;
        Ok(Dvn::new(conf))
    }

    /// Listen to the corresponding events coming from the LayerZeroV2 contract and the SendLib
    /// contract.
    ///
    /// The worflow pushes the received packets into the queue, which will then be verified if the
    /// DVN fee is paid and thus assigned to its verification (otherwise, they will be dropped).
    /// It is assumed that the packets and the fee-paid events arrive in the correct order (first
    /// the packet to be verified, then the fee-paid event).
    pub async fn listen(&mut self) -> Result<()> {
        // Create the WS subscriptions for listening to the events.
        let (_provider, mut endpoint_stream, mut sendlib_stream) = build_dvn_subscriptions(&self.config).await?;
        self.target_receivelib = Some(self.create_receivelib_contract()?);

        info!("Configured DVN: {:?}", self.config.source_dvn);
        info!("Listening to chain events...");

        loop {
            debug!("Queue with size: {:?}", self.packet_queue.len());
            tokio::select! {
                // From the LayerZeroV2 Endpoint, the event `PacketSent`, which contains the message to be sent.
                Some(log) = endpoint_stream.next() => {
                    if let Err(e) = self.endpoint_log_logic(&log) {
                        error!("Error processing Endpoint log. Error: {:?}", e);
                    }
                },
                // From the SendLib, the event which triggers the verification: `DVNFeePaid`.
                Some(log) = sendlib_stream.next() => {
                    if let Err(e) = self.sendlib_log_logic(&log).await {
                        error!("Error processing SendLib log. Error: {:?}", e);
                    }
                },
            }
        }
    }

    /// Run the workflow when receiving a [`Log`] from the LayerZero endpoint.
    ///
    /// Store the packet at the end of the queue.
    fn endpoint_log_logic(&mut self, log: &Log) -> Result<()> {
        log.log_decode::<PacketSent>().map_or_else(
            |e| {
                error!("Received a `PacketSent` event but failed to decode it: {:?}", e);
                Err(eyre!(e))
            },
            |inner_log| {
                self.packet_queue.push_back(inner_log.data().clone());
                debug!("Received packet added to queue.");
                Ok(())
            },
        )
    }

    /// Run the worflow when receiving a [`Log`] from the SendLib.
    ///
    /// If the DVN is assigned, poco the packet from the queue (they are pushed to the back when
    /// received, and popped from the front when processing them).
    async fn sendlib_log_logic(&mut self, log: &Log) -> Result<()> {
        match log.log_decode::<DVNFeePaid>() {
            Err(e) => {
                error!("Received a `DVNFeePaid` event but failed to decode it: {:?}", e);
                Err(eyre!(e))
            }
            Ok(inner_log) if inner_log.inner.requiredDVNs.contains(&self.config.source_dvn) => {
                info!("DVN assigned. Starting verification process...");

                // Get the stored packet.
                let packet = self.packet_queue.pop_front().ok_or_else(|| {
                    error!("No packet stored to verify.");
                    eyre!("No packet stored to verify.")
                })?;

                // Query the number of confirmations.
                let required_confirmations = self.get_required_confirmations().await?;

                // Check if the info from the payload could have been extracted.
                let header = self.get_header(&packet);
                let _header_hash = self.get_header_hash(&packet);
                let message_hash = self.get_message_hash(&packet);

                // FIXME: some contracts don't have `_verified` method, work around this
                //
                //let already_verified = query_already_verified(
                //    self.receivelib_contract.as_ref().ok_or_eyre("No ReceiveLib contract")?,
                //    self.config.source_dvn,
                //    header_hash.as_ref(),
                //    message_hash.as_ref(),
                //    required_confirmations,
                //)
                //.await?;

                // Verify the the message..
                //if !already_verified {
                self.verify_message(log, message_hash, required_confirmations, &header)
                    .await?;
                //}

                // NOTE: in the DVN docs, it's not said to commit the verification, somewhere
                // else it's mentioned to do it by "any address".
                //
                // Commit the verification to the ReceiveLib contract.
                //self.commit_verification(&header, message_hash).await?;

                info!("Verification process completed successfully.");
                Ok(())
            }
            Ok(_) => {
                // Remove the packet from the queue, since we are not assigned to verify it.
                debug!("Received `DVNFeePaid` event but configured DVN is not included. Dropping stored packet.");
                self.packet_queue.pop_front();
                Ok(())
            }
        }
    }

    #[allow(dead_code)] // not used for now, might be necessary
    /// Commit the verification to the ReceiveLib contract.
    async fn commit_verification(&mut self, header: &[u8], message_hash: B256) -> Result<()> {
        if let Some(contract) = &self.target_receivelib {
            commit(contract, header, message_hash.as_ref()).await
        } else {
            error!("Cannot commit verification. Missing `ReceiveLib` contract");
            Err(eyre!("Cannot commit verification. Missing `ReceiveLib` contract"))
        }
    }

    /// Get the required confirmations from the ReceiveLib contract.
    async fn get_required_confirmations(&mut self) -> Result<U256> {
        // Check if there's a receive lib
        let Some(receive_lib) = &self.target_receivelib else {
            error!("No `ReceiveLib` contract present in worker to query confirmations");
            return Err(eyre!(
                "No `ReceiveLib` contract present in worker to query confirmations"
            ));
        };

        // Query how many confirmations are required.
        let remote_eid = U256::from(self.config.target_network_eid);

        // Query the confirmations from the receive lib contract.
        match query_confirmations(receive_lib, remote_eid).await {
            Ok(confirmations) => Ok(confirmations),
            Err(e) => {
                error!("Failed to query confirmations. Error: {:?}", e);
                Err(eyre!(e))
            }
        }
    }

    /// Verify the message by calling the `verify` function on the target chain's ReceiveLib contract.
    pub(crate) async fn verify_message(
        &mut self,
        log: &Log,
        message_hash: B256,
        required_confirmations: U256,
        header: &[u8],
    ) -> Result<()> {
        debug!("Packet NOT verified. Calling verification.");

        let Some(_block_height) = log.block_number else {
            error!("Block number is `None`, can't verify packet.");
            return Err(eyre!("Block number is `None`, can't verify packet."));
        };

        // FIXME: for now, just verify everything. Uncomment when everything else works
        //
        //if let Some(verifier) = self.verifier.as_ref() {
        //    if let Err(report) = verifier.verify(block_height).await {
        //        error!("Failed to verify the state root. Error: {:?}", report);
        //        return;
        //    } else {
        //        info!("State root verified successfully.");
        //    }
        //} else {
        //    error!("Verifier not present")
        //}

        if let Some(contract) = &self.target_receivelib {
            verify(contract, header, message_hash.as_ref(), required_confirmations).await
        } else {
            error!("Cannot verify packet. Missing `ReceiveLib` contract");
            Err(eyre!("Cannot verify packet. Missing `ReceiveLib` contract"))
        }
    }

    /// Create a handle to interact with the ReceiveLib contract in the target chain.
    pub fn create_receivelib_contract(&self) -> Result<ContractInst> {
        // Create an HTTP provider to call the target ULN contract
        let http_provider = get_http_provider(&self.config.target_http_rpc_url)?;
        // Get the relevant contract ABI, and create contract
        let abi = get_abi_from_path(RECEIVELIB_ABI_PATH)?;
        // Create the contradt instance
        let contract = create_contract_instance(self.config.target_receivelib, http_provider, abi);
        Ok(contract)
    }

    /// Extract the header from the packet.
    pub fn get_header(&self, packet: &PacketSent) -> Vec<u8> {
        header(packet.encodedPayload.as_ref()).to_vec()
    }

    /// Get the hash of the packet's header.
    pub fn get_header_hash(&self, packet: &PacketSent) -> B256 {
        keccak256(header(packet.encodedPayload.as_ref()))
    }

    /// Get the hash of the packet's message.
    pub fn get_message_hash(&self, packet: &PacketSent) -> B256 {
        keccak256(message(packet.encodedPayload.as_ref()))
    }
}
