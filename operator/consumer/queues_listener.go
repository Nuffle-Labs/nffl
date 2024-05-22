package consumer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/near/borsh-go"
	rmq "github.com/rabbitmq/amqp091-go"
)

var (
	QueueExistsError = errors.New("Queue already exists")
)

// Type reflections of NEAR DA client submission format
type Commitment = [32]byte
type TransactionId = [32]byte

type CommittedBlob struct {
	Commitment Commitment
	Data       []byte
}

// Type reflection of MQ format
type PublishPayload struct {
	TransactionId TransactionId
	Data          []byte
}

type QueuesListener struct {
	receivedBlocksC    chan<- BlockData
	queueDeliveryCs    map[uint32]<-chan rmq.Delivery
	queueDeliveryMutex sync.Mutex

	logger        logging.Logger
	eventListener EventListener
}

func NewQueuesListener(receivedBlocksC chan<- BlockData, eventListener EventListener, logger logging.Logger) *QueuesListener {
	listener := QueuesListener{
		receivedBlocksC: receivedBlocksC,
		queueDeliveryCs: make(map[uint32]<-chan rmq.Delivery),
		logger:          logger,
		eventListener:   eventListener,
	}

	return &listener
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
			l.eventListener.OnArrival()

			publishPayload := new(PublishPayload)
			err := borsh.Deserialize(publishPayload, d.Body)
			if err != nil {
				l.logger.Error("Error deserializing payload")
				l.eventListener.OnFormatError()
				d.Reject(false)

				continue
			}

			committedBlob := new(CommittedBlob)
			err = borsh.Deserialize(committedBlob, publishPayload.Data)
			if err != nil {
				l.logger.Error("Invalid blob", "d.Body", d.Body, "err", err)
				l.eventListener.OnFormatError()
				d.Reject(false)

				continue
			}

			var blocks []*types.Block
			if err := rlp.DecodeBytes(committedBlob.Data, &blocks); err != nil {
				l.logger.Warn("Invalid block", "rollupId", rollupId, "err", err)
				l.eventListener.OnFormatError()

				continue
			}

			for _, block := range blocks {
				blockData := BlockData{
					RollupId:      rollupId,
					TransactionId: publishPayload.TransactionId,
					Commitment:    committedBlob.Commitment,
					Block:         *block,
				}

				l.logger.Info(
					"MQ Block",
					"rollupId", rollupId,
					"blockHeight", blockData.Block.Header().Number.Uint64(),
					"transactionId", blockData.TransactionId,
					"commitment", blockData.Commitment,
					"listener", fmt.Sprintf("%p", l),
				)
				l.receivedBlocksC <- blockData
			}

			l.logger.Info("Acking delivery", "rollupId", rollupId)
			d.Ack(false)

		case <-ctx.Done():
			l.logger.Info("Consumer context canceled")
			// TODO: some closing and canceling here
			return
		}
	}
}
