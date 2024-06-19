package operator2_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	operator "github.com/NethermindEth/near-sffl/operator/v2"
	"github.com/stretchr/testify/assert"
)

var _ = operator.RpcClient(&MockRpcClient{})

type MockRpcClient struct {
	call func(serviceMethod string, args any, reply any) error
}

func (self *MockRpcClient) Call(serviceMethod string, args any, reply any) error {
	return self.call(serviceMethod, args, reply)
}

func NoopRpcClient() *MockRpcClient {
	return &MockRpcClient{
		call: func(serviceMethod string, args any, reply any) error { return nil },
	}
}

type MockListener struct {
	incError   func()
	incSuccess func()
}

var _ = operator.Listener(&MockListener{})

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
		call: func(serviceMethod string, args any, reply any) error {
			logger.Debug("MockRpcClient.Call", "method", serviceMethod, "args", args)
			rpcClientCallCount++
			return nil
		},
	}

	client := operator.NewAggregatorRpcClient(&listener, &rpcClient, operator.NeverRetry, logger)
	go client.Start(ctx)

	client.SendSignedStateRootUpdateToAggregator(&messages.SignedStateRootUpdateMessage{})
	client.SendSignedOperatorSetUpdateToAggregator(&messages.SignedOperatorSetUpdateMessage{})

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

	client := operator.NewAggregatorRpcClient(listener, rpcClient, operator.NeverRetry, logger)
	go client.Start(ctx)

	time.Sleep(1 * time.Second)
}

func TestMultipleConcurrentClose(t *testing.T) {
	ctx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)
	listener := NoopListener()
	rpcClient := NoopRpcClient()

	client := operator.NewAggregatorRpcClient(listener, rpcClient, operator.NeverRetry, logger)
	go client.Start(ctx)

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

	client := operator.NewAggregatorRpcClient(listener, &rpcClient, operator.AlwaysRetry, logger)
	go client.Start(ctx)

	client.SendSignedStateRootUpdateToAggregator(&messages.SignedStateRootUpdateMessage{})

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
		call: func(serviceMethod string, args any, reply any) error {
			rpcFailCount++
			return assert.AnError
		},
	}

	client := operator.NewAggregatorRpcClient(listener, &rpcClient, operator.RetryAtMost(4), logger)
	go client.Start(ctx)

	client.SendSignedStateRootUpdateToAggregator(&messages.SignedStateRootUpdateMessage{})

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
		call: func(serviceMethod string, args any, reply any) error {
			time.Sleep(100 * time.Millisecond)

			rpcFailCount++
			return assert.AnError
		},
	}

	client := operator.NewAggregatorRpcClient(listener, &rpcClient, operator.RetryIfRecentEnough(500*time.Millisecond), logger)
	go client.Start(ctx)

	client.SendSignedCheckpointTaskResponseToAggregator(&messages.SignedCheckpointTaskResponse{})

	time.Sleep(500 * time.Millisecond)
	client.Close()

	assert.Equal(t, 5, rpcFailCount)
}

func TestGetAggregatedCheckpointMessages(t *testing.T) {
	ctx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)
	listener := NoopListener()

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

	client := operator.NewAggregatorRpcClient(listener, &rpcClient, operator.NeverRetry, logger)
	go client.Start(ctx)

	response, err := client.GetAggregatedCheckpointMessages(0, 0)
	time.Sleep(200 * time.Millisecond)

	assert.NoError(t, err)
	assert.Equal(t, expected, *response)

	client.Close()
}
