package operator2

import (
	"errors"
	"net/rpc"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	v1 "github.com/NethermindEth/near-sffl/operator"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
)

type RpcClient interface {
	Call(serviceMethod string, args any, reply any) error
}

var _ RpcClient = (*rpc.Client)(nil)

func NewHTTPAggregatorRpcClient(aggregatorIpPortAddr string, operatorId eigentypes.OperatorId, expectedRegistryCoordinatorAddress common.Address, logger logging.Logger) (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", aggregatorIpPortAddr)
	if err != nil {
		logger.Error("Error dialing aggregator rpc client", "err", err)
		return nil, err
	}

	var aggregatorRegistryCoordinatorAddress string
	err = client.Call("Aggregator.GetRegistryCoordinatorAddress", struct{}{}, &aggregatorRegistryCoordinatorAddress)
	if err != nil {
		logger.Error("Received error when getting registry coordinator address", "err", err)
		return nil, err
	}

	logger.Debug("Notifying aggregator of initialization")

	var reply bool
	err = client.Call("Aggregator.NotifyOperatorInitialization", operatorId, &reply)
	if err != nil {
		logger.Error("Error notifying aggregator of initialization", "err", err)
		return nil, err
	}

	if common.HexToAddress(aggregatorRegistryCoordinatorAddress).Cmp(expectedRegistryCoordinatorAddress) != 0 {
		logger.Fatal("Registry coordinator address from aggregator does not match the one in the config", "aggregator", aggregatorRegistryCoordinatorAddress, "config", expectedRegistryCoordinatorAddress.String())
		return nil, errors.New("mismatching registry coordinator address from aggregator")
	}

	return client, nil
}

type RetryStrategy func(submittedAt time.Time, err error) bool

func NeverRetry(_ time.Time, _ error) bool {
	return false
}

func AlwaysRetry(_ time.Time, _ error) bool {
	return true
}

func RetryIfRecentEnough(ttl time.Duration) RetryStrategy {
	return func(submittedAt time.Time, err error) bool {
		return time.Since(submittedAt) < ttl
	}
}

func RetryAtMost(retries int) RetryStrategy {
	retryCount := 0
	return func(_ time.Time, err error) bool {
		result := retryCount < retries
		retryCount++
		return result
	}
}

type AggregatorRpcClient struct {
	rpcClient   RpcClient
	shouldRetry RetryStrategy
	listener    v1.RpcClientEventListener
	logger      logging.Logger
}

var _ core.Metricable = (*AggregatorRpcClient)(nil)

func NewAggregatorRpcClient(rpcClient RpcClient, retryPredicate RetryStrategy, logger logging.Logger) AggregatorRpcClient {
	return AggregatorRpcClient{
		rpcClient:   rpcClient,
		shouldRetry: retryPredicate,
		listener:    &v1.SelectiveRpcClientListener{},
		logger:      logger,
	}
}

func (a *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *messages.SignedCheckpointTaskResponse) error {
	a.logger.Info("Sending signed task response header to aggregator", "signedCheckpointTaskResponse", signedCheckpointTaskResponse)

	submittedAt := time.Now()
	var reply bool
	action := func() error {
		err := a.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)
		if err != nil {
			a.logger.Error("Received error from aggregator", "err", err)
		}
		return err
	}

	retried := false
	err := action()
	for err != nil && a.shouldRetry(submittedAt, err) {
		a.listener.IncErroredCheckpointSubmissions(retried)
		err = action()
		retried = true
	}

	if err != nil {
		a.logger.Info("Signed task response header accepted by aggregator", "reply", reply)
		a.listener.IncCheckpointTaskResponseSubmissions(retried)
		a.listener.ObserveLastCheckpointIdResponded(signedCheckpointTaskResponse.TaskResponse.ReferenceTaskIndex)
		a.listener.OnMessagesReceived()
	} else {
		a.logger.Error("Dropping message after error", "err", err)
	}

	return err
}

func (a *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *messages.SignedStateRootUpdateMessage) error {
	a.logger.Info("Sending signed state root update message to aggregator", "signedStateRootUpdateMessage", signedStateRootUpdateMessage)

	submittedAt := time.Now()
	var reply bool
	action := func() error {
		err := a.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", signedStateRootUpdateMessage, &reply)
		if err != nil {
			a.logger.Error("Received error from aggregator", "err", err)
		}
		return err
	}

	retried := false
	err := action()
	for err != nil && a.shouldRetry(submittedAt, err) {
		a.listener.IncErroredStateRootUpdateSubmissions(signedStateRootUpdateMessage.Message.RollupId, retried)
		err = action()
		retried = true
	}

	if err != nil {
		a.logger.Info("Signed state root update message accepted by aggregator", "reply", reply)
		a.listener.IncStateRootUpdateSubmissions(signedStateRootUpdateMessage.Message.RollupId, retried)
		a.listener.OnMessagesReceived()
	} else {
		a.logger.Error("Dropping message after error", "err", err)
	}

	return err
}

func (a *AggregatorRpcClient) SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage *messages.SignedOperatorSetUpdateMessage) error {
	a.logger.Info("Sending operator set update message to aggregator", "signedOperatorSetUpdateMessage", signedOperatorSetUpdateMessage)

	submittedAt := time.Now()
	var reply bool
	action := func() error {
		err := a.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", signedOperatorSetUpdateMessage, &reply)
		if err != nil {
			a.logger.Error("Received error from aggregator", "err", err)
		}
		return err
	}

	retried := false
	err := action()
	for err != nil && a.shouldRetry(submittedAt, err) {
		a.listener.IncErroredOperatorSetUpdateSubmissions(retried)
		err = action()
		retried = true
	}

	if err != nil {
		a.logger.Info("Signed operator set update message accepted by aggregator", "reply", reply)
		a.listener.IncOperatorSetUpdateUpdateSubmissions(retried)
		a.listener.ObserveLastOperatorSetUpdateIdResponded(signedOperatorSetUpdateMessage.Message.Id)
		a.listener.OnMessagesReceived()
	} else {
		a.logger.Error("Dropping message after error", "err", err)
	}

	return err
}

func (a *AggregatorRpcClient) GetAggregatedCheckpointMessages(fromTimestamp, toTimestamp uint64) (messages.CheckpointMessages, error) {
	type Args struct {
		FromTimestamp, ToTimestamp uint64
	}

	submittedAt := time.Now()
	var checkpointMessages messages.CheckpointMessages
	action := func() error {
		return a.rpcClient.Call("Aggregator.GetAggregatedCheckpointMessages", &Args{fromTimestamp, toTimestamp}, &checkpointMessages)
	}

	err := action()
	for err != nil && a.shouldRetry(submittedAt, err) {
		err = action()
	}

	return checkpointMessages, err
}

func (c *AggregatorRpcClient) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := v1.MakeRpcClientMetrics(registry)
	if err != nil {
		return err
	}

	c.listener = listener
	return nil
}
