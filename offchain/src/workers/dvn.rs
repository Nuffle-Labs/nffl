use crate::chain::ContractInst;
use crate::{
    abi::{L0V2EndpointAbi::PacketSent, SendLibraryAbi::DVNFeePaid},
    chain::{
        connections::{build_dvn_subscriptions, get_abi_from_path, get_http_provider},
        contracts::{create_contract_instance, query_already_verified, query_confirmations, verify},
    },
    config::WorkerConfig,
    data::packet_v1_codec::{header, message},
    verifier::NFFLVerifier,
};
use alloy::primitives::{keccak256, B256, U256};
use alloy::rpc::types::Log;
use eyre::Result;
use futures::stream::StreamExt;
use tracing::{debug, error, info, warn};

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
        Ok(Dvn::new(crate::config::WorkerConfig::load_from_env()?))
    }

    pub fn listening(&mut self) {
        self.status = DvnStatus::Listening;
    }

    pub fn packet_received(&mut self, packet: PacketSent) {
        self.packet = Some(packet);
        self.status = DvnStatus::PacketReceived;
    }

    pub fn reset_packet(&mut self) {
        self.packet = None;
        self.status = DvnStatus::Listening;
        debug!("DVN not required, stored packet dropped")
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

    pub(crate) async fn verify_message(&mut self, log: &Log, message_hash: B256, required_confirmations: U256) {
        debug!("Packet NOT verified. Calling verification.");
        self.verifying();

        if log.block_number.is_none() {
            error!("Block number is None, can't verify Packet.");
            return;
        }

        if let Some(verifier) = self.verifier.as_ref() {
            if let Err(report) = verifier.verify(log.block_number.unwrap()).await {
                error!("Failed to verify the state root. Error: {:?}", report);
                return;
            }
        } else {
            error!("Verifier not present")
        }

        if let (Some(receive_lib), Some(header)) = (self.receive_lib.as_ref(), self.get_header()) {
            verify(receive_lib, header, message_hash.as_ref(), required_confirmations).await;
        }
    }

    pub async fn listen(&mut self) -> Result<()> {
        // Create the WS subscriptions for listening to the events.
        let (_provider, mut endpoint_stream, mut sendlib_stream) = build_dvn_subscriptions(&self.config).await?;

        // Create an HTTP provider to call contract functions.
        let http_provider = get_http_provider(&self.config)?;

        // Get the relevant contract ABI, and create contract.
        let abi = get_abi_from_path("./abi/ReceiveLibUln302.json")?;
        self.receive_lib = Some(create_contract_instance(
            self.config.receivelib_uln302_addr,
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
                    self.endpoint_log_logic(&log);
                },
                // From the SendLib, we need the event which triggers the verification: `DVNFeePaid`.
                Some(log) = sendlib_stream.next() => {
                    self.sendlib_log_logic(&log).await;
                },
            }
        }
    }

    /// Run the corresponding logic when receiving a [`Log`] from the LayerZero endpoint.
    fn endpoint_log_logic(&mut self, log: &Log) {
        match log.log_decode::<PacketSent>() {
            Err(e) => {
                error!("Received a `PacketSent` event but failed to decode it: {:?}", e);
            }
            Ok(inner_log) => {
                debug!("PacketSent event found and decoded.");
                self.packet_received(inner_log.data().clone());
            }
        }
    }

    /// Run the corresponding logic when receiving a [`Log`] from the SendLib.
    async fn sendlib_log_logic(&mut self, log: &Log) {
        match log.log_decode::<DVNFeePaid>() {
            Err(e) => {
                error!("Received a `DVNFeePaid` event but failed to decode it: {:?}", e);
            }
            Ok(inner_log) if self.packet.is_some() => {
                info!("`DVNFeePaid` event decoded, `Packet` present.");

                let required_dvns = &inner_log.inner.requiredDVNs;
                let own_dvn_addr = self.config.dvn_addr;

                if required_dvns.contains(&own_dvn_addr) {
                    debug!("Found DVN in required DVNs.");

                    // Query how many confirmations are required.
                    let remote_eid = U256::from(self.config.target_network_eid);

                    let Some(receive_lib) = &self.receive_lib else {
                        error!("No `ReceiveLib` contract present in worker to query confirmations");
                        return;
                    };

                    let Ok(required_confirmations) = query_confirmations(&receive_lib, remote_eid).await else {
                        error!("Cannot query `requiredConfirmations` from `ReceiveLib` contract");
                        return;
                    };

                    // Prepare the header hash.
                    let header_hash = self.get_header_hash();
                    // Prepare the payload hash.
                    let message_hash = self.get_message_hash();

                    // Check if the info from the payload could have been extracted.
                    match (header_hash, message_hash) {
                        (_, None) => {
                            error!("Cannot hash payload");
                        }
                        (None, _) => {
                            error!("Cannot hash message");
                        }
                        (Some(header_hash), Some(message_hash)) => {
                            let already_verified = query_already_verified(
                                receive_lib,
                                own_dvn_addr,
                                header_hash.as_ref(),
                                message_hash.as_ref(),
                                required_confirmations,
                            )
                            .await;

                            if !already_verified {
                                let _ = self.verify_message(&log, message_hash, required_confirmations).await;
                            }
                        }
                    }
                } else {
                    debug!("DVN not required");
                    self.reset_packet();
                }
            }
            Ok(_) => {
                warn!("Received a `DVNFeePaid` event but don't have information about the `Packet` to be verified");
            }
        }
    }
}
