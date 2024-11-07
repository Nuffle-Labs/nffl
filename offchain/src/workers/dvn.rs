use crate::{
    abi::{L0V2EndpointAbi::PacketSent, SendLibraryAbi::DVNFeePaid},
    chain::{
        connections::{build_dvn_subscriptions, get_abi_from_path, get_http_provider},
        contracts::{
            create_contract_instance,
            // query_already_verified,
            query_confirmations,
            verify,
        },
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
use futures::stream::StreamExt;
use tracing::{debug, error, info};

const RECEIVELIB_ABI_PATH: &str = "offchain/abi/ReceiveLibUln302.json";

pub enum DvnStatus {
    Stopped,
    Listening,
    PacketReceived,
    Verifying,
}

pub struct Dvn {
    pub config: WorkerConfig,
    pub status: DvnStatus,
    pub packet: Option<PacketSent>,
    pub receive_lib: Option<ContractInst>,
    pub verifier: Option<NFFLVerifier>,
}

impl Dvn {
    pub fn new(config: WorkerConfig) -> Self {
        Self {
            config,
            status: DvnStatus::Stopped,
            packet: None,
            receive_lib: None,
            verifier: None,
        }
    }

    pub fn new_from_env() -> Result<Self> {
        let conf = WorkerConfig::load_from_env()?;
        Ok(Dvn::new(conf))
    }

    pub fn listening(&mut self) {
        self.status = DvnStatus::Listening;
    }

    pub fn packet_received(&mut self, packet: PacketSent) {
        self.packet = Some(packet);
        self.status = DvnStatus::PacketReceived;
    }

    pub fn reset_packet_storage(&mut self) {
        self.packet = None;
        self.status = DvnStatus::Listening;
    }

    pub fn verifying(&mut self) {
        self.status = DvnStatus::Verifying;
    }

    pub fn get_header(&self) -> Option<&[u8]> {
        if let Some(packet) = self.packet.as_ref() {
            Some(header(packet.encodedPayload.as_ref()))
        } else {
            None
        }
    }

    pub fn get_header_hash(&self) -> Option<B256> {
        self.packet
            .as_ref()
            .map(|packet| keccak256(header(packet.encodedPayload.as_ref())))
    }

    pub fn get_message_hash(&self) -> Option<B256> {
        self.packet
            .as_ref()
            .map(|packet| keccak256(message(packet.encodedPayload.as_ref())))
    }

    pub(crate) async fn verify_message(
        &mut self,
        log: &Log,
        message_hash: B256,
        required_confirmations: U256,
    ) -> Result<()> {
        debug!("Packet NOT verified. Calling verification.");
        self.verifying();

        let Some(_block_height) = log.block_number else {
            error!("Block number is `None`, can't verify packet.");
            return Err(eyre!("Block number is `None`, can't verify packet."));
        };

        // FIXME: for now, just verify everything.
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

        if let (Some(receive_lib), Some(header)) = (self.receive_lib.as_ref(), self.get_header()) {
            verify(receive_lib, header, message_hash.as_ref(), required_confirmations).await
        } else {
            error!("Cannot verify packet. Missing `ReceiveLib` contract or `Packet`");
            Err(eyre!("Cannot verify packet. Missing `ReceiveLib` contract or `Packet`"))
        }
    }

    pub async fn listen(&mut self) -> Result<()> {
        // Create the WS subscriptions for listening to the events.
        let (_provider, mut endpoint_stream, mut sendlib_stream) = build_dvn_subscriptions(&self.config).await?;

        // Create an HTTP provider to call contract functions on the target chain.
        let http_provider = get_http_provider(&self.config.target_http_rpc_url)?;

        // Get the relevant contract ABI, and create contract.
        let abi = get_abi_from_path(RECEIVELIB_ABI_PATH)?;
        self.receive_lib = Some(create_contract_instance(
            self.config.target_receivelib,
            http_provider,
            abi,
        )?);

        // Start listening for events
        info!("Listening to chain events...");
        self.listening();

        loop {
            tokio::select! {
                // From the LayerZeroV2 Endpoint, we need the event `PacketSent`, which contains information about the message to be sent.
                Some(log) = endpoint_stream.next() => {
                    if let Err(e) = self.endpoint_log_logic(&log) {
                        error!("Error processing endpoint log. Error: {:?}", e);
                    }
                },
                // From the SendLib, we need the event which triggers the verification: `DVNFeePaid`.
                Some(log) = sendlib_stream.next() => {
                    if let Err(e) = self.sendlib_log_logic(&log).await {
                        error!("Error processing sendlib log. Error: {:?}", e);
                    }
                },
            }
        }
    }

    /// Run the corresponding logic when receiving a [`Log`] from the LayerZero endpoint.
    fn endpoint_log_logic(&mut self, log: &Log) -> Result<()> {
        log.log_decode::<PacketSent>().map_or_else(
            |e| {
                error!("Received a `PacketSent` event but failed to decode it: {:?}", e);
                Err(eyre!(e))
            },
            |p| {
                self.packet_received(p.data().clone());
                Ok(())
            },
        )
    }

    /// Run the corresponding logic when receiving a [`Log`] from the SendLib.
    async fn sendlib_log_logic(&mut self, log: &Log) -> Result<()> {
        match log.log_decode::<DVNFeePaid>() {
            Err(e) => {
                error!("Received a `DVNFeePaid` event but failed to decode it: {:?}", e);
                Err(eyre!(e))
            }
            Ok(inner_log)
                if self.packet.is_some() && inner_log.inner.requiredDVNs.contains(&self.config.source_dvn) =>
            {
                info!("Required DVN in event. Starting verification process.");

                let required_confirmations = match self.get_required_confirmations().await {
                    Ok(confirmations) => confirmations,
                    Err(e) => {
                        error!("Failed to get required confirmations. Error: {:?}", e);
                        return Err(eyre!(e));
                    }
                };

                // Check if the info from the payload could have been extracted.
                match (self.get_header_hash(), self.get_message_hash()) {
                    (_, None) => {
                        error!("Cannot hash payload");
                        Err(eyre!("Cannot hash payload"))
                    }
                    (None, _) => {
                        error!("Cannot hash message");
                        Err(eyre!("Cannot hash message"))
                    }
                    (Some(_header_hash), Some(message_hash)) => {
                        // FIXME: some contracts don't have `_verified` method, work around
                        // this
                        //
                        //let already_verified = query_already_verified(
                        //    receive_lib,
                        //    own_dvn_addr,
                        //    header_hash.as_ref(),
                        //    message_hash.as_ref(),
                        //    required_confirmations,
                        //)
                        //.await;

                        //if !already_verified {
                        self.verify_message(log, message_hash, required_confirmations).await?;
                        //}
                        info!("Verification process completed successfully.");
                        Ok(())
                    }
                }
            }
            Ok(_) => {
                debug!("Received a `DVNFeePaid` event but DVN is not included to verify.");
                self.reset_packet_storage();
                Ok(())
            }
        }
    }

    async fn get_required_confirmations(&self) -> Result<U256> {
        // If there's a receive lib from the target
        let Some(receive_lib) = &self.receive_lib else {
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
}
