package consumer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	rmq "github.com/rabbitmq/amqp091-go"
)

const (
	reconnectDelay = 3 * time.Second
	rechannelDelay = 2 * time.Second
)

var (
	// TODO:
	QueueNamesToNetworkId = map[string]uint{
		"da-mq": 0,
		// Add mappings
	}
	defaultQueues = compilerDefaultQueues()

	errAlreadyClosed = errors.New("already closed: not connected to the server")
)

func compilerDefaultQueues() []string {
	keys := make([]string, 0, len(QueueNamesToNetworkId))
	for k := range QueueNamesToNetworkId {
		keys = append(keys, k)
	}

	return keys
}

type ConsumerConfig struct {
	Addr        string
	ConsumerTag string
}

type BlockData struct {
	NetworkId uint
	Block     types.Block
}

// TODO: add logger
type Consumer struct {
	consumerTag string
	blockstream chan BlockData

	isReady           bool
	contextCancelFunc context.CancelFunc
	connection        *rmq.Connection
	onConnClosed      <-chan *rmq.Error
	channel           *rmq.Channel
	onChanClosed      <-chan *rmq.Error

	queues         []string
	queuesListener QueuesListener
}

// TODO: Pass default queues in config?
func NewConsumer(config ConsumerConfig) Consumer {
	// TODO: context.TODO() or background?
	ctx, cancel := context.WithCancel(context.TODO())
	consumer := Consumer{
		consumerTag:       config.ConsumerTag,
		queues:            defaultQueues,
		blockstream:       make(chan BlockData),
		contextCancelFunc: cancel,
	}

	go consumer.Reconnect(config.Addr, ctx)
	return consumer
}

func (consumer *Consumer) Reconnect(addr string, ctx context.Context) {
	for {
		fmt.Println("Reconnecting...")

		consumer.isReady = false
		conn, err := consumer.connect(addr)
		if err != nil {
			fmt.Println(err)

			select {
			case <-ctx.Done():
				return
			case <-time.After(reconnectDelay):
			}

			continue
		}

		if done := consumer.ResetChannel(conn, ctx); done {
			return
		}

		fmt.Println("Connected")
		select {
		case err := <-ctx.Done():
			ctx.Err()
			fmt.Println(err)
			// deref cancel smth?
			break

		case err := <-consumer.onConnClosed:
			if !err.Recover {
				fmt.Println(err)
				break
			}

			fmt.Println("Connection closed with err:", err, "Reconnecting...")

		case err := <-consumer.onChanClosed:
			if !err.Recover {
				fmt.Println(err)
				break
			}

			// TODO: Reconnect not the whole connection but just a channel?
			fmt.Println("Channel closed with err:", err, "Reconnecting...")
		}
	}
}

func (consumer *Consumer) connect(addr string) (*rmq.Connection, error) {
	conn, err := rmq.Dial(addr)
	if err != nil {
		return nil, err
	}

	consumer.changeConnection(conn)
	return conn, nil
}

func (consumer *Consumer) changeConnection(conn *rmq.Connection) {
	consumer.connection = conn

	closeNotifier := make(chan *rmq.Error)
	consumer.onConnClosed = conn.NotifyClose(closeNotifier)
}

func (consumer *Consumer) ResetChannel(conn *rmq.Connection, ctx context.Context) bool {
	for {
		consumer.isReady = false

		err := consumer.setupChannel(conn, ctx)
		if err != nil {
			fmt.Println(err)

			select {
			case <-ctx.Done():
				fmt.Println("Consumer ctx canceled")
				return true

			case rmqError := <-consumer.onConnClosed:
				if rmqError.Recover {
					fmt.Println("channel can't recover error")
					return true
				}

				fmt.Println("ResetChannel err:", rmqError)
				return false
			case <-time.After(rechannelDelay):
			}

			continue
		}

		return false
	}
}

func (consumer *Consumer) setupChannel(conn *rmq.Connection, ctx context.Context) error {
	// TODO: create multiple chanels?
	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	queueDeliveries := make(map[string]<-chan rmq.Delivery)
	for i := range consumer.queues {
		queue, err := channel.QueueDeclare(consumer.queues[i], true, false, false, false, nil)
		if err != nil {
			return err
		}

		deliveries, err := channel.Consume(
			queue.Name,
			consumer.consumerTag,
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			return err
		}

		queueDeliveries[queue.Name] = deliveries
	}

	listener := NewQueuesListener(queueDeliveries, consumer.blockstream, ctx)
	consumer.queuesListener = listener

	consumer.changeChannel(channel)
	consumer.isReady = true
	return nil
}

func (consumer *Consumer) changeChannel(channel *rmq.Channel) {
	consumer.channel = channel

	closeNotifer := make(chan *rmq.Error)
	consumer.onChanClosed = channel.NotifyClose(closeNotifer)
}

func (consumer *Consumer) Close(ctx context.Context) error {
	if !consumer.isReady {
		return errAlreadyClosed
	}

	// shut down goroutines
	consumer.contextCancelFunc()

	err := consumer.channel.Close()
	if err != nil {
		return err
	}

	err = consumer.connection.Close()
	if err != nil {
		return err
	}
	consumer.isReady = false
	return nil
}

func (consumer Consumer) GetBlockStream() <-chan BlockData {
	return consumer.blockstream
}
