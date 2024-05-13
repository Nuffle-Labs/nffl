package operator

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

const (
	ResendInterval = 2 * time.Second
	MaxRetries     = 10
)

type AggregatorRpcClienter interface {
	core.Metricable

	SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *messages.SignedCheckpointTaskResponse)
	SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *messages.SignedStateRootUpdateMessage)
	SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage *messages.SignedOperatorSetUpdateMessage)
	GetAggregatedCheckpointMessages(fromTimestamp, toTimestamp uint64) (*messages.CheckpointMessages, error)
}

type unsentRpcMessage struct {
	Message interface{}
	Retries int
}

type AggregatorRpcClient struct {
	rpcClientLock              sync.RWMutex
	rpcClient                  *rpc.Client
	aggregatorIpPortAddr       string
	registryCoordinatorAddress common.Address

	unsentMessagesLock sync.Mutex
	unsentMessages     []unsentRpcMessage
	resendTicker       *time.Ticker

	logger   logging.Logger
	listener RpcClientEventListener
}

var _ core.Metricable = (*AggregatorRpcClient)(nil)

func NewAggregatorRpcClient(aggregatorIpPortAddr string, registryCoordinatorAddress common.Address, logger logging.Logger) (*AggregatorRpcClient, error) {
	resendTicker := time.NewTicker(ResendInterval)

	client := &AggregatorRpcClient{
		// set to nil so that we can create an rpc client even if the aggregator is not running
		rpcClient:                  nil,
		logger:                     logger,
		aggregatorIpPortAddr:       aggregatorIpPortAddr,
		registryCoordinatorAddress: registryCoordinatorAddress,
		unsentMessages:             make([]unsentRpcMessage, 0),
		resendTicker:               resendTicker,
		listener:                   &SelectiveRpcClientListener{},
	}

	go client.onTick()
	return client, nil
}

func (c *AggregatorRpcClient) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := MakeRpcClientMetrics(registry)
	if err != nil {
		return err
	}

	c.listener = listener
	return nil
}

func (c *AggregatorRpcClient) dialAggregatorRpcClient() error {
	c.rpcClientLock.Lock()
	defer c.rpcClientLock.Unlock()

	if c.rpcClient != nil {
		return nil
	}

	c.logger.Info("rpc client is nil. Dialing aggregator rpc client")

	client, err := rpc.DialHTTP("tcp", c.aggregatorIpPortAddr)
	if err != nil {
		c.logger.Error("Error dialing aggregator rpc client", "err", err)
		return err
	}

	var aggregatorRegistryCoordinatorAddress string
	err = client.Call("Aggregator.GetRegistryCoordinatorAddress", struct{}{}, &aggregatorRegistryCoordinatorAddress)
	if err != nil {
		c.logger.Info("Received error when getting registry coordinator address", "err", err)
		return err
	}

	if common.HexToAddress(aggregatorRegistryCoordinatorAddress).Cmp(c.registryCoordinatorAddress) != 0 {
		c.logger.Fatal("Registry coordinator address from aggregator does not match the one in the config", "aggregator", aggregatorRegistryCoordinatorAddress, "config", c.registryCoordinatorAddress.String())
		return errors.New("mismatching registry coordinator address from aggregator")
	}

	c.rpcClient = client

	return nil
}

func (c *AggregatorRpcClient) InitializeClientIfNotExist() error {
	c.rpcClientLock.RLock()
	if c.rpcClient != nil {
		c.rpcClientLock.RUnlock()
		return nil
	}
	c.rpcClientLock.RUnlock()

	return c.dialAggregatorRpcClient()
}

func isShutdownOrNetworkError(err error) bool {
	if err == rpc.ErrShutdown {
		return true
	}

	if _, ok := err.(*net.OpError); ok {
		return true
	}

	return false
}

func (c *AggregatorRpcClient) handleRpcError(err error) error {
	if isShutdownOrNetworkError(err) {
		go c.handleRpcShutdown()
	}

	return nil
}

func (c *AggregatorRpcClient) handleRpcShutdown() {
	c.rpcClientLock.Lock()
	defer c.rpcClientLock.Unlock()

	if c.rpcClient != nil {
		c.logger.Info("Closing RPC client due to shutdown")

		err := c.rpcClient.Close()
		if err != nil {
			c.logger.Error("Error closing RPC client", "err", err)
		}

		c.rpcClient = nil
	}
}

func (c *AggregatorRpcClient) onTick() {
	for {
		<-c.resendTicker.C

		err := c.InitializeClientIfNotExist()
		if err != nil {
			c.logger.Error("Error initializing client", "err", err)
			continue
		}

		c.unsentMessagesLock.Lock()
		if len(c.unsentMessages) == 0 {
			c.unsentMessagesLock.Unlock()
			continue
		}
		c.unsentMessagesLock.Unlock()

		c.tryResendFromDeque()
	}
}

// Expected to be called with initialized client.
func (c *AggregatorRpcClient) tryResendFromDeque() {
	c.rpcClientLock.RLock()
	defer c.rpcClientLock.RUnlock()

	c.unsentMessagesLock.Lock()
	defer c.unsentMessagesLock.Unlock()

	if len(c.unsentMessages) != 0 {
		c.logger.Info("Resending messages from queue")
	}

	errorPos := 0
	for i := 0; i < len(c.unsentMessages); i++ {
		entry := c.unsentMessages[i]
		message := entry.Message

		// Assumes client exists
		var err error
		var reply bool

		switch message := message.(type) {
		case *messages.SignedCheckpointTaskResponse:
			err = c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", message, &reply)
			if err != nil {
				c.listener.IncErroredCheckpointSubmissions(true)
			} else {
				c.listener.IncCheckpointTaskResponseSubmissions(true)
				c.listener.ObserveLastCheckpointIdResponded(message.TaskResponse.ReferenceTaskIndex)
			}

		case *messages.SignedStateRootUpdateMessage:
			err = c.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", message, &reply)
			if err != nil {
				c.listener.IncErroredStateRootUpdateSubmissions(message.Message.RollupId, true)
			} else {
				c.listener.IncStateRootUpdateSubmissions(message.Message.RollupId, true)
			}

		case *messages.SignedOperatorSetUpdateMessage:
			err = c.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", message, &reply)
			if err != nil {
				c.listener.IncErroredOperatorSetUpdateSubmissions(true)
			} else {
				c.listener.IncOperatorSetUpdateUpdateSubmissions(true)
				c.listener.ObserveLastOperatorSetUpdateIdResponded(message.Message.Id)
			}

		default:
			panic("unreachable")
		}

		if err != nil {
			c.logger.Error("Couldn't resend message", "err", err)

			if isShutdownOrNetworkError(err) {
				c.logger.Error("Couldn't resend message due to shutdown or network error")

				if errorPos == 0 {
					c.unsentMessages = c.unsentMessages[i:]
					return
				}

				for j := i; j < len(c.unsentMessages); j++ {
					rpcMessage := c.unsentMessages[j]
					c.unsentMessages[errorPos] = rpcMessage
					errorPos++
				}

				break
			}

			entry.Retries++
			if entry.Retries >= MaxRetries {
				c.logger.Error("Max retries reached, dropping message", "message", fmt.Sprintf("%#v", message))
				continue
			}

			c.unsentMessages[errorPos] = entry
			errorPos++
		}
	}

	c.unsentMessages = c.unsentMessages[:errorPos]
	c.listener.ObserveResendQueueSize(len(c.unsentMessages))
}

func (c *AggregatorRpcClient) sendOperatorMessage(sendCb func() error, message interface{}) {
	c.rpcClientLock.RLock()
	defer c.rpcClientLock.RUnlock()

	appendProtected := func() {
		c.unsentMessagesLock.Lock()
		c.unsentMessages = append(c.unsentMessages, unsentRpcMessage{Message: message})
		c.listener.ObserveResendQueueSize(len(c.unsentMessages))
		c.unsentMessagesLock.Unlock()
	}

	if c.rpcClient == nil {
		appendProtected()
		return
	}

	c.logger.Info("Sending request to aggregator")
	err := sendCb()
	if err != nil {
		c.handleRpcError(err)
		appendProtected()
		return
	}
}

func (c *AggregatorRpcClient) sendRequest(sendCb func() error) error {
	c.rpcClientLock.RLock()
	defer c.rpcClientLock.RUnlock()

	if c.rpcClient == nil {
		return errors.New("rpc client is nil")
	}

	c.logger.Info("Sending request to aggregator")

	err := sendCb()
	if err != nil {
		c.handleRpcError(err)
		return err
	}

	c.logger.Info("Request successfully sent to aggregator")

	return nil
}

func (c *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *messages.SignedCheckpointTaskResponse) {
	c.logger.Info("Sending signed task response header to aggregator", "signedCheckpointTaskResponse", fmt.Sprintf("%#v", signedCheckpointTaskResponse))

	c.sendOperatorMessage(func() error {
		var reply bool
		err := c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)
		if err != nil {
			c.listener.IncErroredCheckpointSubmissions(false)

			c.logger.Info("Received error from aggregator", "err", err)
			return err
		}

		c.listener.IncCheckpointTaskResponseSubmissions(false)
		c.listener.ObserveLastCheckpointIdResponded(signedCheckpointTaskResponse.TaskResponse.ReferenceTaskIndex)

		c.logger.Info("Signed task response header accepted by aggregator.", "reply", reply)
		c.listener.OnMessagesReceived()
		return nil
	}, signedCheckpointTaskResponse)
}

func (c *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *messages.SignedStateRootUpdateMessage) {
	c.logger.Info("Sending signed state root update message to aggregator", "signedStateRootUpdateMessage", fmt.Sprintf("%#v", signedStateRootUpdateMessage))

	c.sendOperatorMessage(func() error {
		var reply bool
		err := c.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", signedStateRootUpdateMessage, &reply)
		if err != nil {
			c.listener.IncErroredStateRootUpdateSubmissions(signedStateRootUpdateMessage.Message.RollupId, false)

			c.logger.Info("Received error from aggregator", "err", err)
			return err
		}

		c.listener.IncStateRootUpdateSubmissions(signedStateRootUpdateMessage.Message.RollupId, false)

		c.logger.Info("Signed state root update message accepted by aggregator.", "reply", reply)
		c.listener.OnMessagesReceived()
		return nil
	}, signedStateRootUpdateMessage)
}

func (c *AggregatorRpcClient) SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage *messages.SignedOperatorSetUpdateMessage) {
	c.logger.Info("Sending operator set update message to aggregator", "signedOperatorSetUpdateMessage", fmt.Sprintf("%#v", signedOperatorSetUpdateMessage))

	c.sendOperatorMessage(func() error {
		var reply bool
		err := c.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", signedOperatorSetUpdateMessage, &reply)
		if err != nil {
			c.listener.IncErroredOperatorSetUpdateSubmissions(false)

			c.logger.Info("Received error from aggregator", "err", err)
			return err
		}

		c.listener.IncOperatorSetUpdateUpdateSubmissions(false)
		c.listener.ObserveLastOperatorSetUpdateIdResponded(signedOperatorSetUpdateMessage.Message.Id)

		c.logger.Info("Signed operator set update message accepted by aggregator.", "reply", reply)
		c.listener.OnMessagesReceived()
		return nil
	}, signedOperatorSetUpdateMessage)
}

func (c *AggregatorRpcClient) GetAggregatedCheckpointMessages(fromTimestamp, toTimestamp uint64) (*messages.CheckpointMessages, error) {
	c.logger.Info("Getting checkpoint messages from aggregator")

	var checkpointMessages messages.CheckpointMessages

	type Args struct {
		FromTimestamp, ToTimestamp uint64
	}

	err := c.sendRequest(func() error {
		err := c.rpcClient.Call("Aggregator.GetAggregatedCheckpointMessages", &Args{fromTimestamp, toTimestamp}, &checkpointMessages)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
			return err
		}

		c.logger.Info("Checkpoint messages fetched from aggregator")
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &checkpointMessages, nil
}
