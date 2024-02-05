package operator

import (
	"fmt"
	"net/rpc"
	"time"

	"github.com/NethermindEth/near-sffl/aggregator"
	"github.com/NethermindEth/near-sffl/metrics"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

type AggregatorRpcClienter interface {
	SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *aggregator.SignedCheckpointTaskResponse)
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

func (c *AggregatorRpcClient) sendRequest(sendCb func(retryCount int) error, maxRetries int, retryInterval time.Duration) {
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
		err := sendCb(i)

		if err == nil {
			c.logger.Info("Data successfully sent to aggregator")
			return
		}

		c.logger.Infof("Retrying in %s", retryInterval.String())
		time.Sleep(retryInterval)
	}

	c.logger.Error("Could not send data to aggregator. Tried 5 times.")
}

func (c *AggregatorRpcClient) SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse *aggregator.SignedCheckpointTaskResponse) {
	c.sendRequest(func(retryCount int) error {
		var reply bool

		if retryCount == 0 {
			c.logger.Info("Sending signed task response header to aggregator", "signedCheckpointTaskResponse", fmt.Sprintf("%#v", signedCheckpointTaskResponse))
		}

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

func (c *AggregatorRpcClient) SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage *aggregator.SignedStateRootUpdateMessage) {
	c.sendRequest(func(retryCount int) error {
		var reply bool

		if retryCount == 0 {
			c.logger.Info("Sending signed state root update message to aggregator", "signedStateRootUpdateMessage", fmt.Sprintf("%#v", signedStateRootUpdateMessage))
		}

		err := c.rpcClient.Call("Aggregator.ProcessSignedStateRootUpdateMessage", signedStateRootUpdateMessage, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
		} else {
			c.logger.Info("Signed state root update message accepted by aggregator.", "reply", reply)
		}

		return err
	}, 5, 2*time.Second)
}
