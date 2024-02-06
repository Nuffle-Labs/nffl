package operator

import (
	"context"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	rmq "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/metrics"
	"github.com/NethermindEth/near-sffl/operator/consumer"
	"github.com/NethermindEth/near-sffl/tests"
)

const MOCK_OPERATOR_BLS_PRIVATE_KEY = "69"

// hash of bls_public_key (hardcoded for sk=69)
var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}

// Name starts with Integration test because we don't want it to run with go test ./...
// since this starts a chain and takes longer to run
// TODO(samlaf): buggy test, fix it
func IntegrationTestOperatorRegistration(t *testing.T) {
	anvilCmd := tests.StartAnvilChainAndDeployContracts()
	defer anvilCmd.Process.Kill()
	operator, _, err := createMockOperator()
	assert.Nil(t, err)
	err = operator.RegisterOperatorWithEigenlayer()
	assert.Nil(t, err)
}

type MockConsumer struct {
	blockReceivedC chan consumer.BlockData
}

func createMockConsumer() *MockConsumer {
	return &MockConsumer{
		blockReceivedC: make(chan consumer.BlockData),
	}
}
func (c *MockConsumer) Reconnect(addr string, ctx context.Context) {}
func (c *MockConsumer) ResetChannel(conn *rmq.Connection, ctx context.Context) bool {
	return true
}
func (c *MockConsumer) Close() error {
	return nil
}
func (c *MockConsumer) GetBlockStream() <-chan consumer.BlockData {
	return c.blockReceivedC
}
func (c *MockConsumer) MockReceiveBlockData(data consumer.BlockData) {
	c.blockReceivedC <- data
}

func createMockOperator() (*Operator, *MockConsumer, error) {
	logger := sdklogging.NewNoopLogger()
	reg := prometheus.NewRegistry()
	noopMetrics := metrics.NewNoopMetrics()

	blsPrivateKey, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	if err != nil {
		return nil, nil, err
	}
	operatorKeypair := bls.NewKeyPair(blsPrivateKey)

	mockConsumer := createMockConsumer()

	operator := &Operator{
		logger:                    logger,
		blsKeypair:                operatorKeypair,
		metricsReg:                reg,
		metrics:                   noopMetrics,
		checkpointTaskCreatedChan: make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorId:                MOCK_OPERATOR_ID,
		consumer:                  mockConsumer,
	}

	return operator, mockConsumer, nil
}
