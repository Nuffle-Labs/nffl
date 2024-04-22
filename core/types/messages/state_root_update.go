package messages

import (
	"encoding/binary"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"

	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type StateRootUpdateMessage struct {
	RollupId            coretypes.RollupId
	BlockHeight         coretypes.BlockNumber
	Timestamp           coretypes.Timestamp
	NearDaTransactionId [32]byte
	NearDaCommitment    [32]byte
	StateRoot           [32]byte
}

type SignedStateRootUpdateMessage struct {
	Message      StateRootUpdateMessage
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

func NewStateRootUpdateMessageFromBinding(binding servicemanager.StateRootUpdateMessage) StateRootUpdateMessage {
	return StateRootUpdateMessage(binding)
}

func (msg StateRootUpdateMessage) ToBinding() servicemanager.StateRootUpdateMessage {
	return servicemanager.StateRootUpdateMessage(msg)
}

func (msg StateRootUpdateMessage) AbiEncode() ([]byte, error) {
	typ, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "rollupId", Type: "uint32"},
		{Name: "blockHeight", Type: "uint64"},
		{Name: "timestamp", Type: "uint64"},
		{Name: "nearDaTransactionId", Type: "bytes32"},
		{Name: "nearDaCommitment", Type: "bytes32"},
		{Name: "stateRoot", Type: "bytes32"},
	})
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{{Type: typ}}

	bytes, err := arguments.Pack(msg.ToBinding())
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (msg StateRootUpdateMessage) Digest() ([32]byte, error) {
	data, err := msg.AbiEncode()
	if err != nil {
		return [32]byte{}, err
	}

	digest, err := core.HashMessageWithPrefix([]byte("SFFL::StateRootUpdateMessage"), data)
	if err != nil {
		return [32]byte{}, err
	}

	return digest, nil
}

func (msg StateRootUpdateMessage) Key() [32]byte {
	key := [32]byte{}

	binary.BigEndian.PutUint32(key[20:24], msg.RollupId)
	binary.BigEndian.PutUint64(key[24:32], msg.BlockHeight)

	return key
}
