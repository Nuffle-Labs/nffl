package consumer

import (
	"context"
	"errors"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	rmq "github.com/rabbitmq/amqp091-go"
)

var (
	QueueExistsError = errors.New("Queue already exists")
)

type QueuesListener struct {
	receivedBlocksC chan<- BlockData
	queueDeliveryCs map[uint32]<-chan rmq.Delivery

	logger logging.Logger
}

func NewQueuesListener(receivedBlocksC chan<- BlockData, logger logging.Logger) QueuesListener {
	listener := QueuesListener{
		receivedBlocksC: receivedBlocksC,
		queueDeliveryCs: make(map[uint32]<-chan rmq.Delivery),
		logger:          logger,
	}

	return listener
}

func (listener *QueuesListener) Add(rollupId uint32, rollupDataC <-chan rmq.Delivery, ctx context.Context) error {
	if _, exists := listener.queueDeliveryCs[rollupId]; exists {
		return QueueExistsError
	}

	listener.queueDeliveryCs[rollupId] = rollupDataC
	go listener.listen(rollupId, rollupDataC, ctx)

	return nil
}

func (listener *QueuesListener) listen(rollupId uint32, rollupDataC <-chan rmq.Delivery, ctx context.Context) {
	for {
		select {
		case d, ok := <-rollupDataC:
			if !ok {
				listener.logger.Info("Deliveries channel close", "rollupId", rollupId)
				break
			}

			listener.logger.Info("New delivery", "rollupId", rollupId)

			var block types.Block
			if err := rlp.DecodeBytes(d.Body, &block); err != nil {
				listener.logger.Warn("Invalid block", "rollupId", rollupId, "err", err)
				continue
			}

			listener.receivedBlocksC <- BlockData{RollupId: rollupId, Block: block}
			d.Ack(false)

		case <-ctx.Done():
			listener.logger.Info("Consumer context canceled")
			// TODO: some closing and canceling here
			return
		}
	}
}
