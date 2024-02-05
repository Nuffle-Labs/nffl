package consumer

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	rmq "github.com/rabbitmq/amqp091-go"
)

// TODO: rename to deliveries smth?
type QueuesListener struct {
	blockstream chan BlockData

	queueDeliveries map[string]<-chan rmq.Delivery
}

func NewQueuesListener(deliveries map[string]<-chan rmq.Delivery, ctx context.Context) QueuesListener {
	listener := QueuesListener{
		blockstream:     make(chan BlockData),
		queueDeliveries: deliveries,
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
			// TODO
			if ok {
				break
			}

			// Decode block
			var block types.Block
			if err := rlp.DecodeBytes(d.Body, &block); err != nil {
				fmt.Println("invalid block")
				// TODO: pass error smwr
				continue
			}

			listener.blockstream <- BlockData{NetworkId: id, Block: block}

		case <-ctx.Done():
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
