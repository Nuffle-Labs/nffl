package operator

import (
	"fmt"
	"net/rpc"
	"sync"
	"time"

	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/metrics"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

type AggregatorRpcClienter interface {
	SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *coretypes.SignedCheckpointTaskResponse)
	SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *coretypes.SignedStateRootUpdateMessage)
	SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage *coretypes.SignedOperatorSetUpdateMessage)
}

const (
	CheckpointType = iota
	StateRootType
	OperatorSetUpdateType
)

const (
	ResendInterval = 2 * time.Second
)

type RpcMessageType struct {
	MessageType int
	Message     interface{}
}

type AggregatorRpcClient struct {
	rpcClientMutex sync.Mutex
	rpcClient      *rpc.Client

	metrics              metrics.Metrics
	logger               logging.Logger
	aggregatorIpPortAddr string

	unsentMessagesMutex sync.Mutex
	unsentMessages      []RpcMessageType

	resendTicker *time.Ticker
}

func NewAggregatorRpcClient(aggregatorIpPortAddr string, logger logging.Logger, metrics metrics.Metrics) (*AggregatorRpcClient, error) {
	resendTicker := time.NewTicker(ResendInterval)

	client := &AggregatorRpcClient{
		// set to nil so that we can create an rpc client even if the aggregator is not running
		rpcClient:            nil,
		metrics:              metrics,
		logger:               logger,
		aggregatorIpPortAddr: aggregatorIpPortAddr,
		unsentMessages:       make([]RpcMessageType, 0),
		resendTicker:         resendTicker,
	}

	go client.onTick()
	return client, nil
}

func (c *AggregatorRpcClient) dialAggregatorRpcClient() error {
	c.logger.Info("rpc client is nil. Dialing aggregator rpc client")

	client, err := rpc.DialHTTP("tcp", c.aggregatorIpPortAddr)
	if err != nil {
		return err
	}
	c.rpcClient = client
	return nil
}

func (c *AggregatorRpcClient) InitializeClientIfNotExist() error {
	c.rpcClientMutex.Lock()
	defer c.rpcClientMutex.Unlock()

	if c.rpcClient != nil {
		return nil
	}

	return c.dialAggregatorRpcClient()
}
func (c *AggregatorRpcClient) onTick() {
	tickerC := c.resendTicker.C
	for {
		// TODO(edwin): handle closed chan
		<-tickerC

		// Critically ugly section
		{
			c.unsentMessagesMutex.Lock()
			if len(c.unsentMessages) == 0 {
				c.unsentMessagesMutex.Unlock()
				continue
			}
			c.unsentMessagesMutex.Unlock()
		}

		err := c.InitializeClientIfNotExist()
		if err != nil {
			continue
		}

		c.tryResendFromDeque()
	}
}

// Expected to be called with initialized client.
func (c *AggregatorRpcClient) tryResendFromDeque() {
	c.unsentMessagesMutex.Lock()
	defer c.unsentMessagesMutex.Unlock()

	if len(c.unsentMessages) != 0 {
		c.logger.Info("Resending messages from queue")
	}

	errorPos := 0
	for i := 0; i < len(c.unsentMessages); i++ {
		message := c.unsentMessages[i]

		// Assumes client exists
		var err error
		var reply bool

		switch message.MessageType {
		case CheckpointType:
			signedCheckpointTaskResponse := message.Message.(*coretypes.SignedCheckpointTaskResponse)
			// TODO(edwin): handle error
			err = c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)

		case StateRootType:
			signedStateRootUpdateMessage := message.Message.(*coretypes.SignedStateRootUpdateMessage)
			err = c.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", signedStateRootUpdateMessage, &reply)

		case OperatorSetUpdateType:
			signedOperatorSetUpdateMessage := message.Message.(*coretypes.SignedOperatorSetUpdateMessage)
			err = c.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", signedOperatorSetUpdateMessage, &reply)

		default:
			panic("unreachable")
		}

		if err != nil {
			c.logger.Error("Couldn't resend message", "err", err)

			c.unsentMessages[errorPos] = message
			errorPos++
		}
	}

	c.unsentMessages = c.unsentMessages[:errorPos]
}

func (c *AggregatorRpcClient) sendRequest(sendCb func() error, message RpcMessageType) {
	appendProtected := func() {
		c.unsentMessagesMutex.Lock()
		c.unsentMessages = append(c.unsentMessages, message)
		c.unsentMessagesMutex.Unlock()
	}

	err := c.InitializeClientIfNotExist()
	if err != nil {
		appendProtected()
		return
	}

	c.tryResendFromDeque()

	c.logger.Info("Sending data to aggregator")
	err = sendCb()
	if err != nil {
		appendProtected()
		return
	}
}

func (c *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *coretypes.SignedCheckpointTaskResponse) {
	c.logger.Info("Sending signed task response header to aggregator", "signedCheckpointTaskResponse", fmt.Sprintf("%#v", signedCheckpointTaskResponse))

	message := RpcMessageType{
		MessageType: CheckpointType,
		Message:     signedCheckpointTaskResponse,
	}

	c.sendRequest(func() error {
		var reply bool
		err := c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
			return err
		}

		c.logger.Info("Signed task response header accepted by aggregator.", "reply", reply)
		c.metrics.IncNumMessagesAcceptedByAggregator()
		return nil
	}, message)
}

func (c *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *coretypes.SignedStateRootUpdateMessage) {
	c.logger.Info("Sending signed state root update message to aggregator", "signedStateRootUpdateMessage", fmt.Sprintf("%#v", signedStateRootUpdateMessage))

	message := RpcMessageType{
		MessageType: StateRootType,
		Message:     signedStateRootUpdateMessage,
	}

	c.sendRequest(func() error {
		var reply bool
		err := c.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", signedStateRootUpdateMessage, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
			return err
		}

		c.logger.Info("Signed state root update message accepted by aggregator.", "reply", reply)
		c.metrics.IncNumMessagesAcceptedByAggregator()
		return nil
	}, message)
}

func (c *AggregatorRpcClient) SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage *coretypes.SignedOperatorSetUpdateMessage) {
	c.logger.Info("Sending operator set update message to aggregator", "signedOperatorSetUpdateMessage", fmt.Sprintf("%#v", signedOperatorSetUpdateMessage))

	message := RpcMessageType{
		MessageType: OperatorSetUpdateType,
		Message:     signedOperatorSetUpdateMessage,
	}

	c.sendRequest(func() error {
		var reply bool
		err := c.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", signedOperatorSetUpdateMessage, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
			return err
		}

		c.logger.Info("Signed operator set update message accepted by aggregator.", "reply", reply)
		c.metrics.IncNumMessagesAcceptedByAggregator()
		return nil
	}, message)
}
