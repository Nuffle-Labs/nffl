package operator2

import (
	"context"
	"errors"
	"net/rpc"
	"sync"
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

type Listener interface {
	IncError()
	IncSuccess()
}

type RetryPredicate func(action action, err error) bool

func NeverRetry(_ action, _ error) bool {
	return false
}

func AlwaysRetry(_ action, _ error) bool {
	return true
}

func RetryIfRecentEnough(ttl time.Duration) RetryPredicate {
	return func(action action, err error) bool {
		return time.Since(action.submittedAt) < ttl
	}
}

func RetryAtMost(retries int) RetryPredicate {
	return func(action action, err error) bool {
		return action.retryCount < retries
	}
}

type action struct {
	run         func() error
	submittedAt time.Time
	retryCount  int
}

type AggregatorRpcClient struct {
	listener    Listener
	rpcClient   RpcClient
	shouldRetry RetryPredicate
	logger      logging.Logger
	actionCh    chan action

	once    sync.Once
	closeCh chan struct{}
}

func NewAggregatorRpcClient(listener Listener, rpcClient RpcClient, retryPredicate RetryPredicate, logger logging.Logger) AggregatorRpcClient {
	return AggregatorRpcClient{
		listener:    listener,
		rpcClient:   rpcClient,
		shouldRetry: retryPredicate,
		logger:      logger,
		actionCh:    make(chan action, 10),

		once:    sync.Once{},
		closeCh: make(chan struct{}),
	}
}

func (a *AggregatorRpcClient) Start(ctx context.Context) {
	defer func() {
		a.closeCh <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			a.logger.Debug("AggRpcClient: context done")
			return
		case <-a.closeCh:
			a.logger.Debug("AggRpcClient: close message received")
			return
		case action, ok := <-a.actionCh:
			if !ok {
				continue
			}
			a.logger.Debug("AggRpcClient: action message received")
			err := action.run()
			if err != nil {
				a.logger.Error("AggRpcClient: action failed after retrying", "err", err)
				a.listener.IncError()

				if a.shouldRetry(action, err) {
					a.logger.Debug("AggRpcClient: retrying later")
					action.retryCount++
					a.actionCh <- action
				} else {
					a.logger.Debug("AggRpcClient: not retrying later")
				}
			} else {
				a.logger.Debug("AggRpcClient: action executed successfully")
				a.listener.IncSuccess()
			}
		}
	}
}

func (a *AggregatorRpcClient) Close() {
	a.once.Do(func() {
		a.logger.Debug("AggRpcClient: close")

		close(a.actionCh)
		a.closeCh <- struct{}{}

		<-a.closeCh
		close(a.closeCh)
	})
}

func (a *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(message *messages.SignedCheckpointTaskResponse) {
	a.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return a.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", message, &ignore)
		},
	}
}

func (a *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(message *messages.SignedStateRootUpdateMessage) {
	a.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return a.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", message, &ignore)
		},
	}
}

func (a *AggregatorRpcClient) SendSignedOperatorSetUpdateToAggregator(message *messages.SignedOperatorSetUpdateMessage) {
	a.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return a.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", message, &ignore)
		},
	}
}

// Blocking operation since we want to wait for the return value(s).
func (a *AggregatorRpcClient) GetAggregatedCheckpointMessages(fromTimestamp, toTimestamp uint64) (messages.CheckpointMessages, error) {
	type Args struct {
		FromTimestamp, ToTimestamp uint64
	}

	var checkpointMessages messages.CheckpointMessages

	action := action{
		submittedAt: time.Now(),
		run: func() error {
			return a.rpcClient.Call("Aggregator.GetAggregatedCheckpointMessages", &Args{fromTimestamp, toTimestamp}, &checkpointMessages)
		},
	}
	err := action.run()

	for err != nil && a.shouldRetry(action, err) {
		action.retryCount++
		err = action.run()
	}

	return checkpointMessages, err
}
