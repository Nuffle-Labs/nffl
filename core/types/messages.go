package types

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

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

type CheckpointMessages struct {
	StateRootUpdateMessages              []servicemanager.StateRootUpdateMessage
	StateRootUpdateMessageAggregations   []aggtypes.MessageBlsAggregationServiceResponse
	OperatorSetUpdateMessages            []registryrollup.OperatorSetUpdateMessage
	OperatorSetUpdateMessageAggregations []aggtypes.MessageBlsAggregationServiceResponse
}
