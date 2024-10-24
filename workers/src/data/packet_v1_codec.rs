//! Encoding and decoding as PacketV1Codec from LayerZero.

use crate::data::bytes_utils::BytesUtils;

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

/// Packet struct as defined in messagelib.
pub struct Packet {
    pub nonce: u64,
    pub src_eid: u32,
    pub sender: [u8; 32],
    pub dst_eid: u32,
    pub receiver: [u8; 32],
    pub guid: [u8; 32],
    pub message: Vec<u8>,
}

/// Encode a whole [`Packet`] into a byte array.
pub fn encode(packet: &Packet) -> Vec<u8> {
    [
        &PACKET_VERSION.to_be_bytes()[..],
        &packet.nonce.to_be_bytes()[..],
        &packet.src_eid.to_be_bytes()[..],
        &packet.sender.to_vec()[..],
        &packet.dst_eid.to_be_bytes()[..],
        &packet.receiver[..],
        &packet.guid[..],
        &packet.message,
    ]
    .concat()
}

/// Encode only the [`Packet`]'s header into a byte array.
pub fn encode_packet_header(packet: &Packet) -> Vec<u8> {
    [
        &PACKET_VERSION.to_be_bytes()[..],
        &packet.nonce.to_be_bytes()[..],
        &packet.src_eid.to_be_bytes()[..],
        &packet.sender.to_vec()[..],
        &packet.dst_eid.to_be_bytes()[..],
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

#[cfg(test)]
mod tests {
    use alloy::hex;

    use super::*;

    #[test]
    fn test_encode() {
        let packet = Packet {
            nonce: 1,
            src_eid: 101,
            sender: [1; 32],
            dst_eid: 102,
            receiver: [3; 32],
            guid: [2; 32],
            message: vec![1, 2, 3],
        };

        let encoded = encode(&packet);
        assert_eq!(version(&encoded), PACKET_VERSION);
        assert_eq!(nonce(&encoded), packet.nonce);
        assert_eq!(src_eid(&encoded), packet.src_eid);
        assert_eq!(sender(&encoded), packet.sender);
        assert_eq!(dst_eid(&encoded), packet.dst_eid);
        assert_eq!(receiver(&encoded), packet.receiver);
        assert_eq!(guid(&encoded), packet.guid);
        assert_eq!(message(&encoded), packet.message);

        // assert payload, should equal to guid + message
        let payload_bytes = [&packet.guid[..], packet.message.as_slice()].concat();
        assert_eq!(payload(&encoded), payload_bytes.as_slice());

        // assert header, should equal to version + nonce + path
        let header_bytes = [
            // version
            &PACKET_VERSION.to_be_bytes()[..],
            // nonce
            &packet.nonce.to_be_bytes()[..],
            // path
            &packet.src_eid.to_be_bytes()[..],
            &packet.sender[..],
            &packet.dst_eid.to_be_bytes()[..],
            &packet.receiver[..],
        ]
        .concat();
        assert_eq!(header(&encoded), header_bytes.as_slice());

        // assert sender
        assert_eq!(sender(&encoded), packet.sender);

        // assert receiver
        assert_eq!(receiver(&encoded), packet.receiver);
    }

    #[test]
    fn test_encode_packet_header() {
        let packet = Packet {
            nonce: 1,
            src_eid: 101,
            sender: [1; 32],
            dst_eid: 102,
            receiver: [3; 32],
            guid: [2; 32],
            message: vec![1, 2, 3],
        };

        let encoded = encode_packet_header(&packet);
        assert_eq!(version(&encoded), PACKET_VERSION);
        assert_eq!(nonce(&encoded), packet.nonce);
        assert_eq!(src_eid(&encoded), packet.src_eid);
        assert_eq!(sender(&encoded), packet.sender);
        assert_eq!(dst_eid(&encoded), packet.dst_eid);
        assert_eq!(receiver(&encoded), packet.receiver);

        // assert header, should equal to version + nonce + path
        let header_bytes = [
            // version
            &PACKET_VERSION.to_be_bytes()[..],
            // nonce
            &packet.nonce.to_be_bytes()[..],
            // path
            &packet.src_eid.to_be_bytes()[..],
            &packet.sender[..],
            &packet.dst_eid.to_be_bytes()[..],
            &packet.receiver[..],
        ]
        .concat();
        assert_eq!(header(&encoded), header_bytes.as_slice());
    }

    /// Fixture using Typescript package `lz-utilities-v2` from LayerZero.
    #[test]
    fn decode() {
        let packet = hex::decode("0x010000000000012c810000759e00000000000000000000000019cfce47ed54a88614648dc3f19a5980097007dd000075e80000000000000000000000005634c4a5fed09819e3c46d86a965dd9447d86e479527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00").unwrap();

        let version = version(&packet);
        assert_eq!(version, PACKET_VERSION);

        let nonce = nonce(&packet);
        assert_eq!(nonce, 76929);

        let src_eid = src_eid(&packet);
        assert_eq!(src_eid, 30110);

        let sender = sender(&packet);
        assert_eq!(
            sender,
            hex::decode("0x00000000000000000000000019cfce47ed54a88614648dc3f19a5980097007dd")
                .unwrap()
                .as_slice()
        );

        let dst_eid = dst_eid(&packet);
        assert_eq!(dst_eid, 30184);

        let receiver = receiver(&packet);
        assert_eq!(
            receiver,
            hex::decode("0x0000000000000000000000005634c4a5fed09819e3c46d86a965dd9447d86e47")
                .unwrap()
                .as_slice()
        );

        let guid = guid(&packet);
        assert_eq!(
            guid,
            hex::decode("0x9527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b")
                .unwrap()
                .as_slice()
        );

        let payload = payload(&packet);
        assert_eq!(
            payload,
            hex::decode("0x9527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00")
            .unwrap().as_slice()
        );
    }
}
