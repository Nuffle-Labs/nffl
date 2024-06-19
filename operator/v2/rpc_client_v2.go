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

type AggRpcClient struct {
	listener    Listener
	rpcClient   RpcClient
	shouldRetry RetryPredicate
	logger      logging.Logger
	actionCh    chan action

	once    sync.Once
	closeCh chan struct{}
}

func NewAggRpcClient(listener Listener, rpcClient RpcClient, retryPredicate RetryPredicate, logger logging.Logger) AggRpcClient {
	return AggRpcClient{
		listener:    listener,
		rpcClient:   rpcClient,
		shouldRetry: retryPredicate,
		logger:      logger,
		actionCh:    make(chan action, 10),

		once:    sync.Once{},
		closeCh: make(chan struct{}),
	}
}

func (self *AggRpcClient) Start(ctx context.Context) {
	defer func() {
		self.closeCh <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			self.logger.Debug("AggRpcClient: context done")
			return
		case <-self.closeCh:
			self.logger.Debug("AggRpcClient: close message received")
			return
		case action, ok := <-self.actionCh:
			if !ok {
				continue
			}
			self.logger.Debug("AggRpcClient: action message received")
			err := action.run()
			if err != nil {
				self.logger.Error("AggRpcClient: action failed after retrying", "err", err)
				self.listener.IncError()

				if self.shouldRetry(action, err) {
					self.logger.Debug("AggRpcClient: retrying later")
					action.retryCount++
					self.actionCh <- action
				} else {
					self.logger.Debug("AggRpcClient: not retrying later")
				}
			} else {
				self.logger.Debug("AggRpcClient: action executed successfully")
				self.listener.IncSuccess()
			}
		}
	}
}

func (self *AggRpcClient) Close() {
	self.once.Do(func() {
		self.logger.Debug("AggRpcClient: close")

		close(self.actionCh)
		self.closeCh <- struct{}{}

		<-self.closeCh
		close(self.closeCh)
	})
}

func (self *AggRpcClient) SendProcessSignedCheckpointTaskResponse(message *messages.SignedCheckpointTaskResponse) {
	self.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return self.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", message, &ignore)
		},
	}
}

func (self *AggRpcClient) SendSignedStateRootUpdateMessage(message *messages.SignedStateRootUpdateMessage) {
	self.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return self.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", message, &ignore)
		},
	}
}

func (self *AggRpcClient) SendSignedOperatorSetUpdateMessage(message *messages.SignedOperatorSetUpdateMessage) {
	self.actionCh <- action{
		submittedAt: time.Now(),
		run: func() error {
			var ignore bool
			return self.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", message, &ignore)
		},
	}
}
