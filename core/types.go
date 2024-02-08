package core

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

// we only use a single quorum (quorum 0) for sffl
var QUORUM_NUMBERS = []byte{0}

type BlockNumber = uint32
type TaskIndex = uint32
type RollupId = uint32
type BlockHeight = uint64
type NearBlockHeight = uint64
type MessageDigest = [32]byte

type SignedCheckpointTaskResponse struct {
	TaskResponse taskmanager.CheckpointTaskResponse
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

type SignedStateRootUpdateMessage struct {
	Message      servicemanager.StateRootUpdateMessage
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

type SignedOperatorSetUpdateMessage struct {
	Message      registryrollup.OperatorSetUpdateMessage
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}
