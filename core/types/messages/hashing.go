package messages

import (
	"github.com/NethermindEth/near-sffl/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Hasher struct {
	protocolVersion [32]byte
}

func NewHasher(protocolVersion [32]byte) *Hasher {
	return &Hasher{protocolVersion: protocolVersion}
}

type HasherMessage interface {
	Digest() ([32]byte, error)
	Name() string
}

func (h *Hasher) Hash(message HasherMessage) ([32]byte, error) {
	messageName := message.Name()
	digest, err := message.Digest()
	if err != nil {
		return [32]byte{}, err
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

func (h *Hasher) legacyHash(message interface{}) ([32]byte, error) {
	var messagePrefix string
	var digest [32]byte
	var err error

	switch message := message.(type) {
	case *CheckpointTaskResponse:
		messagePrefix = "SFFL::CheckpointTaskResponse"
		digest, err = message.Digest()
		if err != nil {
			return [32]byte{}, err
		}
	case *StateRootUpdateMessage:
		messagePrefix = "SFFL::StateRootUpdateMessage"
		digest, err = message.Digest()
		if err != nil {
			return [32]byte{}, err
		}
	case *OperatorSetUpdateMessage:
		messagePrefix = "SFFL::OperatorSetUpdateMessage"
		digest, err = message.Digest()
		if err != nil {
			return [32]byte{}, err
		}
	}

	prefixHash := core.Keccak256([]byte(messagePrefix))

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return [32]byte{}, err
	}

	arguments := abi.Arguments{{Type: bytes32Ty}, {Type: bytes32Ty}}

	bytes, err := arguments.Pack(prefixHash, digest)
	if err != nil {
		return [32]byte{}, err
	}

	return core.Keccak256(bytes), nil
}

func (h *Hasher) getDomainSeparator(messageName string, protocolVersion [32]byte) ([32]byte, error) {
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
