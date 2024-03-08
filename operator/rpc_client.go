package operator

import (
	"fmt"
	"net/rpc"
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
type AggregatorRpcClient struct {
	rpcClient            *rpc.Client
	metrics              metrics.Metrics
	logger               logging.Logger
	aggregatorIpPortAddr string
}

func NewAggregatorRpcClient(aggregatorIpPortAddr string, logger logging.Logger, metrics metrics.Metrics) (*AggregatorRpcClient, error) {
	return &AggregatorRpcClient{
		// set to nil so that we can create an rpc client even if the aggregator is not running
		rpcClient:            nil,
		metrics:              metrics,
		logger:               logger,
		aggregatorIpPortAddr: aggregatorIpPortAddr,
	}, nil
}

func (c *AggregatorRpcClient) dialAggregatorRpcClient() error {
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

func (c *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *coretypes.SignedCheckpointTaskResponse) {
	c.logger.Info("Sending signed task response header to aggregator", "signedCheckpointTaskResponse", fmt.Sprintf("%#v", signedCheckpointTaskResponse))

	c.sendRequest(func() error {
		var reply bool

		err := c.rpcClient.Call("Aggregator.ProcessSignedCheckpointTaskResponse", signedCheckpointTaskResponse, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
		} else {
			c.logger.Info("Signed task response header accepted by aggregator.", "reply", reply)
			c.metrics.IncNumTasksAcceptedByAggregator()
		}

		return err
	}, 5, 2*time.Second)
}

func (c *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *coretypes.SignedStateRootUpdateMessage) {
	c.logger.Info("Sending signed state root update message to aggregator", "signedStateRootUpdateMessage", fmt.Sprintf("%#v", signedStateRootUpdateMessage))

	c.sendRequest(func() error {
		var reply bool

		err := c.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", signedStateRootUpdateMessage, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
		} else {
			c.logger.Info("Signed state root update message accepted by aggregator.", "reply", reply)
			c.metrics.IncNumMessagesAcceptedByAggregator()
		}

		return err
	}, 5, 2*time.Second)
}

func (c *AggregatorRpcClient) SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage *coretypes.SignedOperatorSetUpdateMessage) {
	c.logger.Info("Sending operator set update message to aggregator", "signedOperatorSetUpdateMessage", fmt.Sprintf("%#v", signedOperatorSetUpdateMessage))

	c.sendRequest(func() error {
		var reply bool

		err := c.rpcClient.Call("Aggregator.ProcessSignedOperatorSetUpdateMessage", signedOperatorSetUpdateMessage, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
		} else {
			c.logger.Info("Signed operator set update message accepted by aggregator.", "reply", reply)
			c.metrics.IncNumMessagesAcceptedByAggregator()
		}

		return err
	}, 5, 2*time.Second)
}
