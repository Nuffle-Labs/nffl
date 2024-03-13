package messages

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/NethermindEth/near-sffl/aggregator/database/models"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type StateRootUpdateMessage struct {
	RollupId    coretypes.RollupId
	BlockHeight coretypes.BlockNumber
	Timestamp   coretypes.Timestamp
	StateRoot   [32]byte
}

type SignedStateRootUpdateMessage struct {
	Message      StateRootUpdateMessage
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

func NewStateRootUpdateMessageFromModel(model models.StateRootUpdateMessage) StateRootUpdateMessage {
	return StateRootUpdateMessage{
		RollupId:    model.RollupId,
		BlockHeight: model.BlockHeight,
		Timestamp:   model.Timestamp,
		StateRoot:   [32]byte(model.StateRoot),
	}
}

func (msg StateRootUpdateMessage) ToModel() models.StateRootUpdateMessage {
	return models.StateRootUpdateMessage{
		RollupId:    msg.RollupId,
		BlockHeight: msg.BlockHeight,
		Timestamp:   msg.Timestamp,
		StateRoot:   msg.StateRoot[:],
	}
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

	digest, err := core.Keccak256(data)
	if err != nil {
		return [32]byte{}, err
	}

	return digest, nil
}
