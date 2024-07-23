package messages

import (
	"github.com/NethermindEth/near-sffl/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Hasher struct {
	messagingPrefix [32]byte
}

func NewHasher(messagingPrefix [32]byte) *Hasher {
	return &Hasher{messagingPrefix: messagingPrefix}
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

	data, err := abi.Arguments{
		{Type: bytes32Ty},
		{Type: bytes32Ty},
		{Type: bytes32Ty},
	}.Pack(h.messagingPrefix, core.Keccak256([]byte(messageName)), digest)
	if err != nil {
		return [32]byte{}, err
	}

	return core.Keccak256(data), nil
}
