package messages

import (
	"github.com/NethermindEth/near-sffl/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type MessageHasher struct {
	protocolVersion [32]byte
}

func NewMessageHasher(protocolVersion [32]byte) *MessageHasher {
	return &MessageHasher{protocolVersion: protocolVersion}
}

func (h *MessageHasher) Hash(message interface{}) ([32]byte, error) {
	var messageName string
	var digest [32]byte
	var err error

	switch message := message.(type) {
	case *CheckpointTaskResponse:
		messageName = "CheckpointTaskResponse"
		digest, err = message.Digest()
		if err != nil {
			return [32]byte{}, err
		}
	case *StateRootUpdateMessage:
		messageName = "StateRootUpdateMessage"
		digest, err = message.Digest()
		if err != nil {
			return [32]byte{}, err
		}
	case *OperatorSetUpdateMessage:
		messageName = "OperatorSetUpdateMessage"
		digest, err = message.Digest()
		if err != nil {
			return [32]byte{}, err
		}
	}

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return [32]byte{}, err
	}

	domainSeparator, err := h.getDomainSeparator(messageName, h.protocolVersion)
	if err != nil {
		return [32]byte{}, err
	}

	data, err := abi.Arguments{{Type: bytes32Ty}, {Type: bytes32Ty}}.Pack(domainSeparator, digest)
	if err != nil {
		return [32]byte{}, err
	}

	return core.Keccak256(data), nil
}

func (h *MessageHasher) getDomainSeparator(messageName string, protocolVersion [32]byte) ([32]byte, error) {
	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return [32]byte{}, err
	}

	encodedDomainBytes, err := abi.Arguments{{Type: bytes32Ty}, {Type: bytes32Ty}, {Type: bytes32Ty}}.Pack(
		core.Keccak256([]byte("SFFLDomain(bytes32 name,bytes32 protocolVersion)")),
		core.Keccak256([]byte(messageName)),
		protocolVersion,
	)
	if err != nil {
		return [32]byte{}, err
	}

	return core.Keccak256(encodedDomainBytes), nil
}
