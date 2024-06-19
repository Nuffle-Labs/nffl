package operator2

import (
	"context"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

type RpcClient interface {
	Call(method string, args interface{}, reply *bool) error
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

func (a *AggregatorRpcClient) SendProcessSignedCheckpointTaskResponse(message *messages.SignedCheckpointTaskResponse) {
	a.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return a.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", message, &ignore)
		},
	}
}

func (a *AggregatorRpcClient) SendSignedStateRootUpdateMessage(message *messages.SignedStateRootUpdateMessage) {
	a.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return a.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", message, &ignore)
		},
	}
}

func (a *AggregatorRpcClient) SendSignedOperatorSetUpdateMessage(message *messages.SignedOperatorSetUpdateMessage) {
	a.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return a.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", message, &ignore)
		},
	}
}
