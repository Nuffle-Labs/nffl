use crate::abi::PacketStruct;
use crate::codec::bytes_utils::BytesUtils;

pub const PACKET_VERSION: u8 = 1;
pub const PACKET_HEADER_SIZE: usize = 81;

// header (version + nonce + path)
// version
const PACKET_VERSION_OFFSET: usize = 0;
// nonce
const NONCE_OFFSET: usize = 1;
// path
const SRC_EID_OFFSET: usize = 9;
const SENDER_OFFSET: usize = 13;
const DST_EID_OFFSET: usize = 45;
const RECEIVER_OFFSET: usize = 49;
// payload (guid + message)
const GUID_OFFSET: usize = 81;
const MESSAGE_OFFSET: usize = 113;

pub fn encode(packet: &PacketStruct) -> Vec<u8> {
    [
        &PACKET_VERSION.to_be_bytes()[..],
        &packet.nonce.to_be_bytes()[..],
        &packet.srcEid.to_be_bytes()[..],
        &packet.sender.to_vec()[..],
        &packet.dstEid.to_be_bytes()[..],
        &packet.receiver[..],
        &packet.guid[..],
        &packet.message,
    ]
    .concat()
}

pub fn encode_packet_header(packet: &PacketStruct) -> Vec<u8> {
    [
        &PACKET_VERSION.to_be_bytes()[..],
        &packet.nonce.to_be_bytes()[..],
        &packet.srcEid.to_be_bytes()[..],
        &packet.sender.to_vec()[..],
        &packet.dstEid.to_be_bytes()[..],
        &packet.receiver[..],
    ]
    .concat()
}

pub fn header(packet: &[u8]) -> &[u8] {
    &packet[0..GUID_OFFSET]
}

pub fn version(packet: &[u8]) -> u8 {
    packet.to_u8(PACKET_VERSION_OFFSET)
}

pub fn nonce(packet: &[u8]) -> u64 {
    packet.to_u64(NONCE_OFFSET)
}

pub fn src_eid(packet: &[u8]) -> u32 {
    packet.to_u32(SRC_EID_OFFSET)
}

pub fn sender(packet: &[u8]) -> [u8; 32] {
    packet.to_byte_array(SENDER_OFFSET)
}

pub fn dst_eid(packet: &[u8]) -> u32 {
    packet.to_u32(DST_EID_OFFSET)
}

pub fn receiver(packet: &[u8]) -> [u8; 32] {
    packet.to_byte_array(RECEIVER_OFFSET)
}

pub fn guid(packet: &[u8]) -> [u8; 32] {
    packet.to_byte_array(GUID_OFFSET)
}

pub fn message(packet: &[u8]) -> &[u8] {
    &packet[MESSAGE_OFFSET..]
}

pub fn payload(packet: &[u8]) -> &[u8] {
    &packet[GUID_OFFSET..]
}
