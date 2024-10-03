package messages

import (
	"encoding/binary"
	"errors"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"

	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/smt"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type CheckpointTaskResponse struct {
	ReferenceTaskIndex     coretypes.TaskIndex
	StateRootUpdatesRoot   [32]byte
	OperatorSetUpdatesRoot [32]byte
}

type SignedCheckpointTaskResponse struct {
	TaskResponse CheckpointTaskResponse
	BlsSignature bls.Signature
	OperatorId   eigentypes.OperatorId
}

func (s *SignedCheckpointTaskResponse) IsValid() error {
	if s == nil {
		return errors.New("SignedCheckpointTaskResponse is nil")
	}

	if s.BlsSignature.G1Point == nil {
		return errors.New("BlsSignature.G1Point is nil")
	}

	return nil
}

type CheckpointMessages struct {
	StateRootUpdateMessages              []StateRootUpdateMessage
	StateRootUpdateMessageAggregations   []MessageBlsAggregation
	OperatorSetUpdateMessages            []OperatorSetUpdateMessage
	OperatorSetUpdateMessageAggregations []MessageBlsAggregation
}

func NewCheckpointTaskResponseFromMessages(taskIndex coretypes.TaskIndex, checkpointMessages *CheckpointMessages) (CheckpointTaskResponse, error) {
	stateRootUpdatesSmt := smt.NewSMT()
	operatorSetUpdatesSmt := smt.NewSMT()

	for _, msg := range checkpointMessages.StateRootUpdateMessages {
		err := stateRootUpdatesSmt.AddMessage(msg)
		if err != nil {
			return CheckpointTaskResponse{}, err
		}
	}

	err := stateRootUpdatesSmt.Commit()
	if err != nil {
		return CheckpointTaskResponse{}, err
	}

	for _, msg := range checkpointMessages.OperatorSetUpdateMessages {
		err := operatorSetUpdatesSmt.AddMessage(msg)
		if err != nil {
			return CheckpointTaskResponse{}, err
		}
	}

	err = operatorSetUpdatesSmt.Commit()
	if err != nil {
		return CheckpointTaskResponse{}, err
	}

	return CheckpointTaskResponse{
		ReferenceTaskIndex:     taskIndex,
		StateRootUpdatesRoot:   [32]byte(stateRootUpdatesSmt.Root()),
		OperatorSetUpdatesRoot: [32]byte(operatorSetUpdatesSmt.Root()),
	}, nil
}

func NewCheckpointTaskResponseFromBinding(binding taskmanager.CheckpointTaskResponse) CheckpointTaskResponse {
	return CheckpointTaskResponse(binding)
}

func (msg CheckpointTaskResponse) ToBinding() taskmanager.CheckpointTaskResponse {
	return taskmanager.CheckpointTaskResponse(msg)
}

func (msg CheckpointTaskResponse) AbiEncode() ([]byte, error) {
	typ, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "referenceTaskIndex", Type: "uint32"},
		{Name: "stateRootUpdatesRoot", Type: "bytes32"},
		{Name: "operatorSetUpdatesRoot", Type: "bytes32"},
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

func (msg CheckpointTaskResponse) Digest() (coretypes.MessageDigest, error) {
	data, err := msg.AbiEncode()
	if err != nil {
		return [32]byte{}, err
	}

	digest, err := core.HashMessageWithPrefix([]byte("SFFL::CheckpointTaskResponse"), data)
	if err != nil {
		return [32]byte{}, err
	}

	return digest, nil
}

func (msg CheckpointTaskResponse) Key() coretypes.MessageKey {
	key := [32]byte{}

	binary.BigEndian.PutUint32(key[28:32], msg.ReferenceTaskIndex)

	return coretypes.MessageKey(key)
}

func CheckpointTaskResponseKeyToTaskIndex(key coretypes.MessageKey) coretypes.TaskIndex {
	return coretypes.TaskIndex(binary.BigEndian.Uint32(key[28:32]))
}
