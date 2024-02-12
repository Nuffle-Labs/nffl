package mocks

import (
	"context"

	"github.com/NethermindEth/near-sffl/operator/consumer"
	rmq "github.com/rabbitmq/amqp091-go"
)

type MockConsumer struct {
	blockReceivedC chan consumer.BlockData
}

func NewMockConsumer() *MockConsumer {
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
