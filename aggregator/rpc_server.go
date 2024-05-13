package aggregator

import (
	"context"
	"errors"
	"net/http"
	"net/rpc"
	"strings"

	eigentypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
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
func (agg *Aggregator) ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse *messages.SignedCheckpointTaskResponse, reply *bool) error {
	agg.logger.Infof("Received signed task response: %#v", signedCheckpointTaskResponse)

	taskIndex := signedCheckpointTaskResponse.TaskResponse.ReferenceTaskIndex
	taskResponseDigest, err := signedCheckpointTaskResponse.TaskResponse.Digest()
	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	agg.rpcListener.IncTotalSignedCheckpointTaskResponse()
	agg.rpcListener.ObserveLastMessageReceivedTime(signedCheckpointTaskResponse.OperatorId, CheckpointTaskResponseLabel)

	err = agg.taskBlsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedCheckpointTaskResponse.BlsSignature, signedCheckpointTaskResponse.OperatorId,
	)
	if err != nil {
		agg.rpcListener.IncSignedCheckpointTaskResponse(
			signedCheckpointTaskResponse.OperatorId,
			true,
			strings.Contains(err.Error(), "not initialized"),
		)
		return err
	}

	agg.taskResponsesLock.Lock()
	if _, ok := agg.taskResponses[taskIndex]; !ok {
		agg.taskResponses[taskIndex] = make(map[eigentypes.TaskResponseDigest]messages.CheckpointTaskResponse)
	}
	if _, ok := agg.taskResponses[taskIndex][taskResponseDigest]; !ok {
		agg.taskResponses[taskIndex][taskResponseDigest] = signedCheckpointTaskResponse.TaskResponse
	}
	agg.taskResponsesLock.Unlock()

	agg.rpcListener.IncSignedCheckpointTaskResponse(signedCheckpointTaskResponse.OperatorId, false, false)

	return nil
}

func (agg *Aggregator) ProcessSignedStateRootUpdateMessage(signedStateRootUpdateMessage *messages.SignedStateRootUpdateMessage, reply *bool) error {
	messageDigest, err := signedStateRootUpdateMessage.Message.Digest()
	if err != nil {
		agg.logger.Error("Failed to get message digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	hasNearDaCommitment := signedStateRootUpdateMessage.Message.HasNearDaCommitment()
	operatorId := signedStateRootUpdateMessage.OperatorId

	agg.logger.Infof("Received signed state root update message: %#v %#v", signedStateRootUpdateMessage, messageDigest)

	agg.rpcListener.IncTotalSignedCheckpointTaskResponse()
	agg.rpcListener.ObserveLastMessageReceivedTime(operatorId, StateRootUpdateMessageLabel)

	err = agg.stateRootUpdateBlsAggregationService.InitializeMessageIfNotExists(
		messageDigest,
		coretypes.QUORUM_NUMBERS,
		[]eigentypes.QuorumThresholdPercentage{types.MESSAGE_AGGREGATION_QUORUM_THRESHOLD},
		types.MESSAGE_TTL,
		types.MESSAGE_BLS_AGGREGATION_TIMEOUT,
		0,
	)
	if err != nil {
		agg.rpcListener.IncSignedStateRootUpdateMessage(operatorId, true, hasNearDaCommitment)
		return err
	}

	agg.stateRootUpdatesLock.Lock()
	agg.stateRootUpdates[messageDigest] = signedStateRootUpdateMessage.Message
	agg.stateRootUpdatesLock.Unlock()

	err = agg.stateRootUpdateBlsAggregationService.ProcessNewSignature(
		context.Background(), messageDigest,
		&signedStateRootUpdateMessage.BlsSignature, signedStateRootUpdateMessage.OperatorId,
	)
	if err != nil {
		agg.rpcListener.IncSignedStateRootUpdateMessage(operatorId, true, hasNearDaCommitment)
		return err
	}

	agg.rpcListener.IncSignedStateRootUpdateMessage(operatorId, false, hasNearDaCommitment)

	return nil
}

func (agg *Aggregator) ProcessSignedOperatorSetUpdateMessage(signedOperatorSetUpdateMessage *messages.SignedOperatorSetUpdateMessage, reply *bool) error {
	messageDigest, err := signedOperatorSetUpdateMessage.Message.Digest()
	if err != nil {
		agg.logger.Error("Failed to get message digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	operatorId := signedOperatorSetUpdateMessage.OperatorId

	agg.logger.Infof("Received signed operator set update message: %#v", signedOperatorSetUpdateMessage)

	agg.rpcListener.IncTotalSignedOperatorSetUpdateMessage()
	agg.rpcListener.ObserveLastMessageReceivedTime(operatorId, OperatorSetUpdateMessageLabel)

	blockNumber, err := agg.avsReader.GetOperatorSetUpdateBlock(context.Background(), signedOperatorSetUpdateMessage.Message.Id)
	if err != nil {
		agg.rpcListener.IncSignedOperatorSetUpdateMessage(operatorId, true)

		agg.logger.Error("Failed to get operator set update block", "err", err)
		return OperatorSetUpdateBlockNotFoundError500
	}

	err = agg.operatorSetUpdateBlsAggregationService.InitializeMessageIfNotExists(
		messageDigest,
		coretypes.QUORUM_NUMBERS,
		[]eigentypes.QuorumThresholdPercentage{types.MESSAGE_AGGREGATION_QUORUM_THRESHOLD},
		types.MESSAGE_TTL,
		types.MESSAGE_BLS_AGGREGATION_TIMEOUT,
		uint64(blockNumber)-1,
	)
	if err != nil {
		agg.rpcListener.IncSignedOperatorSetUpdateMessage(operatorId, true)
		return err
	}

	agg.operatorSetUpdatesLock.Lock()
	agg.operatorSetUpdates[messageDigest] = signedOperatorSetUpdateMessage.Message
	agg.operatorSetUpdatesLock.Unlock()

	err = agg.operatorSetUpdateBlsAggregationService.ProcessNewSignature(
		context.Background(), messageDigest,
		&signedOperatorSetUpdateMessage.BlsSignature, signedOperatorSetUpdateMessage.OperatorId,
	)
	if err != nil {
		agg.rpcListener.IncSignedOperatorSetUpdateMessage(operatorId, true)
		return err
	}

	agg.rpcListener.IncSignedOperatorSetUpdateMessage(operatorId, false)

	return nil
}

type GetAggregatedCheckpointMessagesArgs struct {
	FromTimestamp, ToTimestamp uint64
}

func (agg *Aggregator) GetAggregatedCheckpointMessages(args *GetAggregatedCheckpointMessagesArgs, reply *messages.CheckpointMessages) error {
	checkpointMessages, err := agg.msgDb.FetchCheckpointMessages(args.FromTimestamp, args.ToTimestamp)
	if err != nil {
		return err
	}

	*reply = *checkpointMessages

	return nil
}

func (agg *Aggregator) GetRegistryCoordinatorAddress(_ *struct{}, reply *string) error {
	*reply = agg.config.SFFLRegistryCoordinatorAddr.String()
	return nil
}
