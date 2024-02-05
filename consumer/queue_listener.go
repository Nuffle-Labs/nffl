package consumer

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	rmq "github.com/rabbitmq/amqp091-go"
)

type QueuesListener struct {
	blockstream     chan<- BlockData
	queueDeliveries map[string]<-chan rmq.Delivery

	logger logging.Logger
}

func NewQueuesListener(deliveries map[string]<-chan rmq.Delivery, blockstream chan<- BlockData, logger logging.Logger, ctx context.Context) QueuesListener {
	listener := QueuesListener{
		blockstream:     blockstream,
		queueDeliveries: deliveries,
		logger:          logger,
	}

	go listener.initListeners(ctx)
	return listener
}

func (listener *QueuesListener) Add(name string, stream <-chan rmq.Delivery, ctx context.Context) {
	go listener.listen(name, stream, ctx)
}

func (listener *QueuesListener) listen(name string, stream <-chan rmq.Delivery, ctx context.Context) {
	id := QueueNamesToNetworkId[name]
	for {
		select {
		case d, ok := <-stream:
			if !ok {
				listener.logger.Info("Deliveries channel close, network", "id", id)
				break
			}

			listener.logger.Info("New delivery, network", "id", id)

			var block types.Block
			if err := rlp.DecodeBytes(d.Body, &block); err != nil {
				// TODO: pass error smwr
				listener.logger.Warn("Invalid block")
				continue
			}

			// TODO: case with multiple consumers from same queue
			listener.blockstream <- BlockData{NetworkId: id, Block: block}
			d.Ack(true)

		case <-ctx.Done():
			listener.logger.Info("Consumer context canceled")
			// TODO: some closing and canceling here
			break
		}
	}
}

func (listener *QueuesListener) initListeners(ctx context.Context) {
	for name, ch := range listener.queueDeliveries {
		listener.Add(name, ch, ctx)
	}
}
