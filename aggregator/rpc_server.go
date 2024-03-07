package aggregator

import (
	"context"
	"errors"
	"net/http"
	"net/rpc"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	MessageDigestNotFoundError500            = errors.New("500. Failed to get message digest")
	OperatorSetUpdateBlockNotFoundError500   = errors.New("500. Failed to get operator set update block")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer() error {
	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse *coretypes.SignedCheckpointTaskResponse, reply *bool) error {
	agg.logger.Infof("Received signed task response: %#v", signedCheckpointTaskResponse)
	taskIndex := signedCheckpointTaskResponse.TaskResponse.ReferenceTaskIndex
	taskResponseDigest, err := core.GetCheckpointTaskResponseDigest(&signedCheckpointTaskResponse.TaskResponse)
	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	err = agg.taskBlsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedCheckpointTaskResponse.BlsSignature, signedCheckpointTaskResponse.OperatorId,
	)
	if err != nil {
		return err
	}

	agg.taskResponsesLock.Lock()
	if _, ok := agg.taskResponses[taskIndex]; !ok {
		agg.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]taskmanager.CheckpointTaskResponse)
	}
	if _, ok := agg.taskResponses[taskIndex][taskResponseDigest]; !ok {
		agg.taskResponses[taskIndex][taskResponseDigest] = signedCheckpointTaskResponse.TaskResponse
	}
	agg.taskResponsesLock.Unlock()

	return nil
}

func (agg *Aggregator) ProcessSignedStateRootUpdateMessage(signedStateRootUpdateMessage *coretypes.SignedStateRootUpdateMessage, reply *bool) error {
	agg.logger.Infof("Received signed state root update message: %#v", signedStateRootUpdateMessage)
	messageDigest, err := core.GetStateRootUpdateMessageDigest(&signedStateRootUpdateMessage.Message)
	if err != nil {
		agg.logger.Error("Failed to get message digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	agg.stateRootUpdateBlsAggregationService.InitializeMessageIfNotExists(messageDigest, coretypes.QUORUM_NUMBERS, []uint32{types.QUORUM_THRESHOLD_NUMERATOR}, types.MESSAGE_TTL, 0)

	err = agg.stateRootUpdateBlsAggregationService.ProcessNewSignature(
		context.Background(), messageDigest,
		&signedStateRootUpdateMessage.BlsSignature, signedStateRootUpdateMessage.OperatorId,
	)
	if err != nil {
		return err
	}

	agg.stateRootUpdatesLock.Lock()
	agg.stateRootUpdates[messageDigest] = signedStateRootUpdateMessage.Message
	agg.stateRootUpdatesLock.Unlock()

	return nil
}

func (agg *Aggregator) ProcessSignedOperatorSetUpdateMessage(signedOperatorSetUpdateMessage *coretypes.SignedOperatorSetUpdateMessage, reply *bool) error {
	agg.logger.Infof("Received signed operator set update message: %#v", signedOperatorSetUpdateMessage)
	messageDigest, err := core.GetOperatorSetUpdateMessageDigest(&signedOperatorSetUpdateMessage.Message)
	if err != nil {
		agg.logger.Error("Failed to get message digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	blockNumber, err := agg.avsReader.GetOperatorSetUpdateBlock(context.Background(), signedOperatorSetUpdateMessage.Message.Id)
	if err != nil {
		agg.logger.Error("Failed to get operator set update block", "err", err)
		return OperatorSetUpdateBlockNotFoundError500
	}

	agg.operatorSetUpdateBlsAggregationService.InitializeMessageIfNotExists(messageDigest, coretypes.QUORUM_NUMBERS, []uint32{types.QUORUM_THRESHOLD_NUMERATOR}, types.MESSAGE_TTL, uint64(blockNumber)-1)

	err = agg.operatorSetUpdateBlsAggregationService.ProcessNewSignature(
		context.Background(), messageDigest,
		&signedOperatorSetUpdateMessage.BlsSignature, signedOperatorSetUpdateMessage.OperatorId,
	)
	if err != nil {
		return err
	}

	agg.operatorSetUpdatesLock.Lock()
	agg.operatorSetUpdates[messageDigest] = signedOperatorSetUpdateMessage.Message
	agg.operatorSetUpdatesLock.Unlock()

	return nil
}
