package rpc_server

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
	"strings"

	"github.com/NethermindEth/near-sffl/core"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Layr-Labs/eigensdk-go/logging"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/NethermindEth/near-sffl/aggregator"
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
	UnknownError400                          = errors.New("400. Unknown error")

	errorsMap = map[error]error{
		aggregator.DigestError:                    MessageDigestNotFoundError500,
		aggregator.TaskResponseDigestError:        TaskResponseDigestNotFoundError500,
		aggregator.GetOperatorSetUpdateBlockError: OperatorSetUpdateBlockNotFoundError500,
	}
)

type RpcServer struct {
	serverIpPortAddr string
	app              aggregator.RpcAggregatorer

	logger   logging.Logger
	listener EventListener
}

var _ core.Metricable = (*RpcServer)(nil)

func NewRpcServer(serverIpPortAddr string, app aggregator.RpcAggregatorer, logger logging.Logger) *RpcServer {
	return &RpcServer{
		serverIpPortAddr: serverIpPortAddr,
		app:              app,
		logger:           logger,
		listener:         &SelectiveRpcListener{},
	}
}

func (s *RpcServer) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := MakeRpcServerMetrics(registry)
	if err != nil {
		return err
	}

	s.listener = listener
	return nil
}

func (s *RpcServer) Start() error {
	s.logger.Info("Starting aggregator rpc server.")

	err := rpc.RegisterName("Aggregator", s)
	if err != nil {
		s.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(s.serverIpPortAddr, nil)
	if err != nil {
		s.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

func mapErrors(err error) error {
	mappedErr, ok := errorsMap[err]
	if !ok {
		return err
	}

	return mappedErr
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (s *RpcServer) ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse *messages.SignedCheckpointTaskResponse, reply *bool) error {
	s.logger.Info("Received signed task response", "response", fmt.Sprintf("%#v", signedCheckpointTaskResponse))
	s.listener.IncTotalSignedCheckpointTaskResponse()
	s.listener.ObserveLastMessageReceivedTime(signedCheckpointTaskResponse.OperatorId, CheckpointTaskResponseLabel)

	err := s.app.ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse)
	if err != nil {
		s.listener.IncSignedCheckpointTaskResponse(
			signedCheckpointTaskResponse.OperatorId,
			true,
			strings.Contains(err.Error(), "not initialized"),
		)
		return err
	}

	s.listener.IncSignedCheckpointTaskResponse(signedCheckpointTaskResponse.OperatorId, false, false)

	return nil
}

func (s *RpcServer) ProcessSignedStateRootUpdateMessage(signedStateRootUpdateMessage *messages.SignedStateRootUpdateMessage, reply *bool) error {
	s.logger.Info("Received signed state root update message", "updateMessage", fmt.Sprintf("%#v", signedStateRootUpdateMessage))
	s.listener.IncTotalSignedCheckpointTaskResponse()
	s.listener.ObserveLastMessageReceivedTime(signedStateRootUpdateMessage.OperatorId, StateRootUpdateMessageLabel)

	hasNearDaCommitment := signedStateRootUpdateMessage.Message.HasNearDaCommitment()
	operatorId := signedStateRootUpdateMessage.OperatorId
	rollupId := signedStateRootUpdateMessage.Message.RollupId

	err := s.app.ProcessSignedStateRootUpdateMessage(signedStateRootUpdateMessage)
	s.listener.IncSignedStateRootUpdateMessage(operatorId, rollupId, err != nil, hasNearDaCommitment)
	if err != nil {
		return mapErrors(err)
	}

	return nil
}

func (s *RpcServer) ProcessSignedOperatorSetUpdateMessage(signedOperatorSetUpdateMessage *messages.SignedOperatorSetUpdateMessage, reply *bool) error {
	s.logger.Info("Received signed operator set update message", "message", fmt.Sprintf("%#v", signedOperatorSetUpdateMessage))

	operatorId := signedOperatorSetUpdateMessage.OperatorId
	s.listener.ObserveLastMessageReceivedTime(operatorId, OperatorSetUpdateMessageLabel)
	s.listener.IncTotalSignedOperatorSetUpdateMessage()

	err := s.app.ProcessSignedOperatorSetUpdateMessage(signedOperatorSetUpdateMessage)
	s.listener.IncSignedOperatorSetUpdateMessage(operatorId, err != nil)
	if err != nil {
		return mapErrors(err)
	}

	return nil
}

type GetAggregatedCheckpointMessagesArgs struct {
	FromTimestamp, ToTimestamp uint64
}

func (s *RpcServer) GetAggregatedCheckpointMessages(args *GetAggregatedCheckpointMessagesArgs, reply *messages.CheckpointMessages) error {
	result, err := s.app.GetAggregatedCheckpointMessages(args.FromTimestamp, args.ToTimestamp)
	if err != nil {
		return mapErrors(err)
	}

	*reply = *result

	return nil
}

func (s *RpcServer) GetRegistryCoordinatorAddress(_ *struct{}, reply *string) error {
	return s.app.GetRegistryCoordinatorAddress(reply)
}

func (s *RpcServer) NotifyOperatorInitialization(operatorId eigentypes.OperatorId, reply *bool) error {
	s.listener.IncOperatorInitializations(operatorId)
	return nil
}
