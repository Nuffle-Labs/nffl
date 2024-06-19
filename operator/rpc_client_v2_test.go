package operator

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/stretchr/testify/assert"
)

type RpcClient interface {
	Call(method string, args interface{}, reply *bool) error
}

var _ = RpcClient(&MockRpcClient{})

type MockRpcClient struct {
	call func(method string, args interface{}, reply *bool) error
}

func (self *MockRpcClient) Call(method string, args interface{}, reply *bool) error {
	return self.call(method, args, reply)
}

func NoopRpcClient() *MockRpcClient {
	return &MockRpcClient{
		call: func(method string, args interface{}, reply *bool) error { return nil },
	}
}

type Listener interface {
	IncError()
	IncSuccess()
}

type MockListener struct {
	incError   func()
	incSuccess func()
}

var _ = Listener(&MockListener{})

func (self *MockListener) IncError() {
	self.incError()
}

func (self *MockListener) IncSuccess() {
	self.incSuccess()
}

func NoopListener() *MockListener {
	return &MockListener{
		incError:   func() {},
		incSuccess: func() {},
	}
}

type RetryPredicate func(action action, err error) bool

func neverRetry(_ action, _ error) bool {
	return false
}

func alwaysRetry(_ action, _ error) bool {
	return true
}

func retryIfRecentEnough(ttl time.Duration) RetryPredicate {
	return func(action action, err error) bool {
		return time.Since(action.submittedAt) < ttl
	}
}

func retryAtMost(retries int) RetryPredicate {
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
	listener  Listener
	rpcClient RpcClient
	logger    logging.Logger
	actionCh  chan action

	once    sync.Once
	closeCh chan struct{}
}

func NewAggRpcClient(listener Listener, rpcClient RpcClient, logger logging.Logger) AggRpcClient {
	return AggRpcClient{
		listener:  listener,
		rpcClient: rpcClient,
		logger:    logger,
		actionCh:  make(chan action, 10),

		once:    sync.Once{},
		closeCh: make(chan struct{}),
	}
}

func (self *AggRpcClient) Start(ctx context.Context, retryLater RetryPredicate) {
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

				if retryLater(action, err) {
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

func TestSendSuccessfulMessages(t *testing.T) {
	ctx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)

	successCount := 0
	listener := MockListener{
		incSuccess: func() { logger.Debug("IncSuccess"); successCount++ },
		incError:   func() { logger.Debug("IncError") },
	}

	rpcClientCallCount := 0
	rpcClient := MockRpcClient{
		call: func(method string, args interface{}, reply *bool) error {
			logger.Debug("MockRpcClient.Call", "method", method, "args", args)
			rpcClientCallCount++
			return nil
		},
	}

	client := NewAggRpcClient(&listener, &rpcClient, logger)
	go client.Start(ctx, neverRetry)

	client.SendSignedStateRootUpdateMessage(&messages.SignedStateRootUpdateMessage{})
	client.SendSignedOperatorSetUpdateMessage(&messages.SignedOperatorSetUpdateMessage{})

	time.Sleep(500 * time.Millisecond)
	client.Close()

	assert.Equal(t, 2, successCount)
	assert.Equal(t, 2, rpcClientCallCount)
}

func TestCloseWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	logger, _ := logging.NewZapLogger(logging.Development)
	listener := NoopListener()
	rpcClient := NoopRpcClient()

	client := NewAggRpcClient(listener, rpcClient, logger)
	go client.Start(ctx, neverRetry)

	time.Sleep(1 * time.Second)
}

func TestMultipleConcurrentClose(t *testing.T) {
	ctx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)
	listener := NoopListener()
	rpcClient := NoopRpcClient()

	client := NewAggRpcClient(listener, rpcClient, logger)
	go client.Start(ctx, neverRetry)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client.Close()
		}()
	}
	wg.Wait()
}

func TestUnboundedRetry(t *testing.T) {
	ctx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)
	listener := NoopListener()

	rpcFailCount := 0
	rpcSuccess := false

	rpcClient := MockRpcClient{
		call: func(method string, args interface{}, reply *bool) error {
			if rpcFailCount < 2 {
				rpcFailCount++
				return assert.AnError
			}

			rpcSuccess = true
			return nil
		},
	}

	client := NewAggRpcClient(listener, &rpcClient, logger)
	go client.Start(ctx, alwaysRetry)

	client.SendSignedStateRootUpdateMessage(&messages.SignedStateRootUpdateMessage{})

	time.Sleep(500 * time.Millisecond)
	client.Close()

	assert.Equal(t, 2, rpcFailCount)
	assert.True(t, rpcSuccess)
}

func TestRetryAtMost(t *testing.T) {
	ctx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)
	listener := NoopListener()

	rpcFailCount := 0

	rpcClient := MockRpcClient{
		call: func(method string, args interface{}, reply *bool) error {
			rpcFailCount++
			return assert.AnError
		},
	}

	client := NewAggRpcClient(listener, &rpcClient, logger)
	go client.Start(ctx, retryAtMost(4))

	client.SendSignedStateRootUpdateMessage(&messages.SignedStateRootUpdateMessage{})

	time.Sleep(500 * time.Millisecond)
	client.Close()

	assert.Equal(t, 5, rpcFailCount) // 1 run, 4 retries
}

func TestRetryLaterIfRecentEnough(t *testing.T) {
	ctx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)
	listener := NoopListener()

	rpcFailCount := 0

	rpcClient := MockRpcClient{
		call: func(method string, args interface{}, reply *bool) error {
			time.Sleep(100 * time.Millisecond)

			rpcFailCount++
			return assert.AnError
		},
	}

	client := NewAggRpcClient(listener, &rpcClient, logger)
	go client.Start(ctx, retryIfRecentEnough(500*time.Millisecond))

	client.SendSignedStateRootUpdateMessage(&messages.SignedStateRootUpdateMessage{})

	time.Sleep(500 * time.Millisecond)
	client.Close()

	assert.Equal(t, 5, rpcFailCount)
}
