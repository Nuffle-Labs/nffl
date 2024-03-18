package consumer

import (
	"context"
	"errors"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	rmq "github.com/rabbitmq/amqp091-go"
)

var (
	QueueExistsError = errors.New("Queue already exists")
)

type QueuesListener struct {
	receivedBlocksC    chan<- BlockData
	queueDeliveryCs    map[uint32]<-chan rmq.Delivery
	queueDeliveryMutex sync.Mutex

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

func (l *QueuesListener) Add(ctx context.Context, rollupId uint32, rollupDataC <-chan rmq.Delivery) error {
	l.queueDeliveryMutex.Lock()
	defer l.queueDeliveryMutex.Unlock()

	if _, exists := l.queueDeliveryCs[rollupId]; exists {
		return QueueExistsError
	}
	l.queueDeliveryCs[rollupId] = rollupDataC

	go l.listen(ctx, rollupId, rollupDataC)

	return nil
}

func (l *QueuesListener) Remove(rollupId uint32) {
	l.queueDeliveryMutex.Lock()
	delete(l.queueDeliveryCs, rollupId)
	l.queueDeliveryMutex.Unlock()
}

func (l *QueuesListener) listen(ctx context.Context, rollupId uint32, rollupDataC <-chan rmq.Delivery) {
	for {
		select {
		case d, ok := <-rollupDataC:
			if !ok {
				l.logger.Info("Deliveries channel close", "rollupId", rollupId)
				l.Remove(rollupId)
				return
			}

			l.logger.Info("New delivery", "rollupId", rollupId)

			var block types.Block
			if err := rlp.DecodeBytes(d.Body, &block); err != nil {
				l.logger.Warn("Invalid block", "rollupId", rollupId, "err", err)
				continue
			}

			l.receivedBlocksC <- BlockData{RollupId: rollupId, Block: block}
			d.Ack(false)

		case <-ctx.Done():
			l.logger.Info("Consumer context canceled")
			// TODO: some closing and canceling here
			return
		}
	}
}
