package operator2

import (
	"errors"
	"net/rpc"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/ethereum/go-ethereum/common"
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
	logger      logging.Logger
}

func NewAggregatorRpcClient(rpcClient RpcClient, retryPredicate RetryStrategy, logger logging.Logger) AggregatorRpcClient {
	return AggregatorRpcClient{
		rpcClient:   rpcClient,
		shouldRetry: retryPredicate,
		logger:      logger,
	}
}

func (a *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(message *messages.SignedCheckpointTaskResponse) error {
	submittedAt := time.Now()
	var ignore bool
	action := func() error {
		return a.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", message, &ignore)
	}

	err := action()
	for err != nil && a.shouldRetry(submittedAt, err) {
		err = action()
	}

	return err
}

func (a *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(message *messages.SignedStateRootUpdateMessage) error {
	submittedAt := time.Now()
	var ignore bool
	action := func() error {
		return a.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", message, &ignore)
	}

	err := action()
	for err != nil && a.shouldRetry(submittedAt, err) {
		err = action()
	}

	return err
}

func (a *AggregatorRpcClient) SendSignedOperatorSetUpdateToAggregator(message *messages.SignedOperatorSetUpdateMessage) error {
	submittedAt := time.Now()
	var ignore bool
	action := func() error {
		return a.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", message, &ignore)
	}

	err := action()
	for err != nil && a.shouldRetry(submittedAt, err) {
		err = action()
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
