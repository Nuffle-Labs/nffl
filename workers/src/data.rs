use crate::abi::L0V2EndpointAbi::PacketSent;
use crate::config;
use crate::config::DVNConfig;
use eyre::Result;
//use alloy::primitives::{Address, U256};

pub struct Dvn {
    config: DVNConfig,
    status: DvnStatus,
    packet: Option<PacketSent>,
    //receivelib_address: Option<Address>,
    //num_confirmations: Option<U256>,
}

pub enum DvnStatus {
    Stopped,
    Listening,
    PacketReceived,
    Verifying,
}

impl Dvn {
    pub fn new(config: DVNConfig) -> Self {
        Self {
            config,
            status: DvnStatus::Stopped,
            packet: None,
            //receivelib_address: None,
            //num_confirmations: None,
        }
    }

    pub fn new_from_env() -> Result<Self> {
        Ok(Dvn::new(config::DVNConfig::load_from_env()?))
    }

    pub fn packet(&self) -> Option<PacketSent> {
        self.packet.clone()
    }

    pub fn config(&self) -> &DVNConfig {
        &self.config
    }

    pub fn listening(&mut self) {
        self.status = DvnStatus::Listening;
    }

    pub fn stop(&mut self) {
        self.status = DvnStatus::Stopped;
    }

    pub fn packet_received(&mut self, packet_log: PacketSent) {
        self.packet = Some(packet_log);
        self.status = DvnStatus::PacketReceived;
    }

    pub fn reset_packet(&mut self) {
        self.packet = None;
        self.status = DvnStatus::Stopped;
    }

    pub fn verifying(&mut self) {
        self.status = DvnStatus::Verifying;
    }
}
