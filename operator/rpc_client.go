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
	ResendInterval = 5 * time.Second
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

	messageDequeMutex sync.Mutex
	messageDeque      []RpcMessageType

	resendTickerMutex sync.Mutex
	resendTicker      *time.Ticker
}

func NewAggregatorRpcClient(aggregatorIpPortAddr string, logger logging.Logger, metrics metrics.Metrics) (*AggregatorRpcClient, error) {
	resendTicker := time.NewTicker(ResendInterval)

	client := &AggregatorRpcClient{
		// set to nil so that we can create an rpc client even if the aggregator is not running
		rpcClient:            nil,
		metrics:              metrics,
		logger:               logger,
		aggregatorIpPortAddr: aggregatorIpPortAddr,
		messageDeque:         make([]RpcMessageType, 0),
		resendTicker:         resendTicker,
	}

	go client.onTick()
	return client, nil
}

func (c *AggregatorRpcClient) onTick() {
	tickerC := c.resendTicker.C
	for {
		// TODO(edwin): handle closed chan
		<-tickerC

		// Critically ugly section
		{
			c.rpcClientMutex.Lock()
			if c.rpcClient == nil {
				c.rpcClientMutex.Unlock()
				continue
			}

			c.rpcClientMutex.Unlock()
		}

		c.tryResendFromDeque()
	}
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

func (c *AggregatorRpcClient) sendRequest(sendCb func() error, maxRetries int, retryInterval time.Duration) {
	if c.rpcClient == nil {
		c.logger.Info("rpc client is nil. Dialing aggregator rpc client")
		err := c.dialAggregatorRpcClient()
		if err != nil {
			c.logger.Error("Could not dial aggregator rpc client. Not sending signed task response header to aggregator. Is aggregator running?", "err", err)
			return
		}
	}

	c.logger.Info("Sending data to aggregator")
	for i := 0; i < maxRetries; i++ {
		err := sendCb()

		if err == nil {
			c.logger.Info("Data successfully sent to aggregator")
			return
		}

		c.logger.Infof("Retrying in %s", retryInterval.String())
		time.Sleep(retryInterval)
	}

	c.logger.Errorf("Could not send data to aggregator. Tried %d times.", maxRetries)
}

func (c *AggregatorRpcClient) InitializeClientIfNotExist() error {
	c.rpcClientMutex.Lock()
	defer c.rpcClientMutex.Unlock()

	if c.rpcClient != nil {
		return nil
	}

	c.logger.Info("rpc client is nil. Dialing aggregator rpc client")
	return c.dialAggregatorRpcClient()
}

// Expected to be called with initialized client.
func (c *AggregatorRpcClient) tryResendFromDeque() error {
	c.messageDequeMutex.Lock()
	defer c.messageDequeMutex.Unlock()

	if len(c.messageDeque) != 0 {
		c.logger.Info("Resending messages from queue")
	}

	for len(c.messageDeque) != 0 {
		message := c.messageDeque[0]

		// Assumes client exists
		var err error
		var reply bool

		switch message.MessageType {
		case CheckpointType:
			signedCheckpointTaskResponse := message.Message.(*coretypes.SignedCheckpointTaskResponse)
			// TODO(edwin): handle error
			err = c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)

			//c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)

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
			return err
		}

		c.messageDeque = c.messageDeque[1:]
	}

	return nil
}

func (c *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *coretypes.SignedCheckpointTaskResponse) {
	c.logger.Info("Sending signed task response header to aggregator", "signedCheckpointTaskResponse", fmt.Sprintf("%#v", signedCheckpointTaskResponse))

	appendProtected := func() {
		c.messageDequeMutex.Lock()
		c.messageDeque = append(c.messageDeque, RpcMessageType{
			MessageType: CheckpointType,
			Message:     signedCheckpointTaskResponse,
		})
		c.messageDequeMutex.Unlock()
	}

	err := c.InitializeClientIfNotExist()
	if err != nil {
		appendProtected()
		return
	}

	err = c.tryResendFromDeque()
	if err != nil {
		appendProtected()
		return
	}

	var reply bool
	err = c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)
	if err != nil {
		c.logger.Info("Received error from aggregator", "err", err)

		// TODO(edwin): filter which errors to append
		return
		appendProtected()
		return
	}

	c.logger.Info("Signed task response header accepted by aggregator.", "reply", reply)
	c.metrics.IncNumMessagesAcceptedByAggregator()
}

func (c *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *coretypes.SignedStateRootUpdateMessage) {
	c.logger.Info("Sending signed state root update message to aggregator", "signedStateRootUpdateMessage", fmt.Sprintf("%#v", signedStateRootUpdateMessage))

	appendProtected := func() {
		c.messageDequeMutex.Lock()
		c.messageDeque = append(c.messageDeque, RpcMessageType{
			MessageType: StateRootType,
			Message:     signedStateRootUpdateMessage,
		})
		c.messageDequeMutex.Unlock()
	}

	err := c.InitializeClientIfNotExist()
	if err != nil {
		appendProtected()
		return
	}

	err = c.tryResendFromDeque()
	if err != nil {
		appendProtected()
		return
	}

	var reply bool
	err = c.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", signedStateRootUpdateMessage, &reply)
	if err != nil {
		c.logger.Info("Received error from aggregator", "err", err)

		appendProtected()
		return
	}

	c.logger.Info("Signed state root update message accepted by aggregator.", "reply", reply)
	c.metrics.IncNumMessagesAcceptedByAggregator()
}

func (c *AggregatorRpcClient) SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage *coretypes.SignedOperatorSetUpdateMessage) {
	c.logger.Info("Sending operator set update message to aggregator", "signedOperatorSetUpdateMessage", fmt.Sprintf("%#v", signedOperatorSetUpdateMessage))

	appendProtected := func() {
		c.messageDequeMutex.Lock()
		c.messageDeque = append(c.messageDeque, RpcMessageType{
			MessageType: OperatorSetUpdateType,
			Message:     signedOperatorSetUpdateMessage,
		})
		c.messageDequeMutex.Unlock()
	}

	err := c.InitializeClientIfNotExist()
	if err != nil {
		appendProtected()
		return
	}

	err = c.tryResendFromDeque()
	if err != nil {
		appendProtected()
		return
	}

	var reply bool
	err = c.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", signedOperatorSetUpdateMessage, &reply)
	if err != nil {
		c.logger.Info("Received error from aggregator", "err", err)

		appendProtected()
		return
	}

	c.logger.Info("Signed operator set update message accepted by aggregator.", "reply", reply)
	c.metrics.IncNumMessagesAcceptedByAggregator()
}

// Init if not initialized
