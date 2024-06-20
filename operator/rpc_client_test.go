package operator_test

import (
	"sync"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/NethermindEth/near-sffl/operator"
	"github.com/stretchr/testify/assert"
)

var _ = operator.RpcClient(&MockRpcClient{})

type MockRpcClient struct {
	lock sync.Mutex
	call func(serviceMethod string, args any, reply any) error
}

func (self *MockRpcClient) Call(serviceMethod string, args any, reply any) error {
	self.lock.Lock()
	defer self.lock.Unlock()

	return self.call(serviceMethod, args, reply)
}

func NoopRpcClient() *MockRpcClient {
	return &MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error { return nil },
	}
}

func TestSendSuccessfulMessages(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	rpcClientCallCount := 0
	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			logger.Debug("MockRpcClient.Call", "method", serviceMethod, "args", args)
			rpcClientCallCount++

			return nil
		},
	}

	client := operator.NewAggregatorRpcClient(&rpcClient, func() operator.RetryStrategy { return operator.NeverRetry }, logger)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		client.SendSignedStateRootUpdateToAggregator(&messages.SignedStateRootUpdateMessage{})
	}()

	go func() {
		defer wg.Done()
		client.SendSignedOperatorSetUpdateToAggregator(&messages.SignedOperatorSetUpdateMessage{})
	}()

	wg.Wait()

	assert.Equal(t, 2, rpcClientCallCount)
}

func TestUnboundedRetry(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	rpcSuccess := false
	rpcFailCount := 0
	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			if rpcFailCount < 2 {
				rpcFailCount++
				return assert.AnError
			}

			rpcSuccess = true
			return nil
		},
	}

	client := operator.NewAggregatorRpcClient(&rpcClient, func() operator.RetryStrategy { return operator.AlwaysRetry }, logger)

	client.SendSignedStateRootUpdateToAggregator(&messages.SignedStateRootUpdateMessage{})

	assert.Equal(t, 2, rpcFailCount)
	assert.True(t, rpcSuccess)
}

func TestUnboundedRetry_Delayed(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	rpcSuccess := false
	rpcFailCount := 0
	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			if rpcFailCount < 2 {
				rpcFailCount++
				return assert.AnError
			}

			rpcSuccess = true
			return nil
		},
	}

	retryFactory := func() operator.RetryStrategy {
		return operator.RetryAnd(operator.RetryWithDelay(100*time.Millisecond), operator.AlwaysRetry)
	}
	client := operator.NewAggregatorRpcClient(&rpcClient, retryFactory, logger)

	startedAt := time.Now()
	client.SendSignedCheckpointTaskResponseToAggregator(&messages.SignedCheckpointTaskResponse{})
	execTime := time.Since(startedAt)

	assert.True(t, execTime > 180*time.Millisecond)
	assert.True(t, execTime < 220*time.Millisecond)
	assert.Equal(t, 2, rpcFailCount)
	assert.True(t, rpcSuccess)
}

func TestRetryAtMost(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	rpcFailCount := 0
	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			rpcFailCount++
			return assert.AnError
		},
	}

	client := operator.NewAggregatorRpcClient(&rpcClient, func() operator.RetryStrategy { return operator.RetryAtMost(4) }, logger)

	client.SendSignedOperatorSetUpdateToAggregator(&messages.SignedOperatorSetUpdateMessage{})

	assert.Equal(t, 5, rpcFailCount) // 1 run, 4 retries
}

func TestRetryAtMost_Concurrent(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	rpcFailCount := 0
	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			rpcFailCount++
			return assert.AnError
		},
	}

	client := operator.NewAggregatorRpcClient(&rpcClient, func() operator.RetryStrategy { return operator.RetryAtMost(4) }, logger)

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		client.SendSignedCheckpointTaskResponseToAggregator(&messages.SignedCheckpointTaskResponse{})
	}()
	go func() {
		defer wg.Done()
		client.SendSignedStateRootUpdateToAggregator(&messages.SignedStateRootUpdateMessage{})
	}()
	go func() {
		defer wg.Done()
		client.SendSignedOperatorSetUpdateToAggregator(&messages.SignedOperatorSetUpdateMessage{})
	}()
	wg.Wait()

	assert.Equal(t, 15, rpcFailCount) // 1 run, 4 retries for each of the 3 calls = 15 calls in total
}

func TestRetryLaterIfRecentEnough(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	rpcFailCount := 0
	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			time.Sleep(100 * time.Millisecond)

			rpcFailCount++
			return assert.AnError
		},
	}

	client := operator.NewAggregatorRpcClient(&rpcClient, func() operator.RetryStrategy { return operator.RetryIfRecentEnough(500 * time.Millisecond) }, logger)

	client.SendSignedStateRootUpdateToAggregator(&messages.SignedStateRootUpdateMessage{})

	assert.Equal(t, 5, rpcFailCount)
}

func TestGetAggregatedCheckpointMessages(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	expected := messages.CheckpointMessages{
		StateRootUpdateMessages: []messages.StateRootUpdateMessage{{BlockHeight: 100}},
	}

	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			switch v := reply.(type) {
			case *messages.CheckpointMessages:
				*v = expected
			}
			return nil
		},
	}

	client := operator.NewAggregatorRpcClient(&rpcClient, func() operator.RetryStrategy { return operator.NeverRetry }, logger)

	response, err := client.GetAggregatedCheckpointMessages(0, 0)

	assert.NoError(t, err)
	assert.Equal(t, expected, *response)
}

func TestGetAggregatedCheckpointMessagesRetry(t *testing.T) {
	logger, _ := logging.NewZapLogger(logging.Development)

	expected := messages.CheckpointMessages{
		StateRootUpdateMessages: []messages.StateRootUpdateMessage{{BlockHeight: 100}},
	}

	rpcFailCount := 0
	rpcClient := MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error {
			if rpcFailCount < 2 {
				rpcFailCount++
				return assert.AnError
			}

			switch v := reply.(type) {
			case *messages.CheckpointMessages:
				*v = expected
			}
			return nil
		},
	}

	client := operator.NewAggregatorRpcClient(&rpcClient, func() operator.RetryStrategy { return operator.AlwaysRetry }, logger)

	response, err := client.GetAggregatedCheckpointMessages(0, 0)

	assert.NoError(t, err)
	assert.Equal(t, expected, *response)
	assert.Equal(t, 2, rpcFailCount)
}
