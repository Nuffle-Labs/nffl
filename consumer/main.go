package main

import (
	// "context"
	"flag"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	rmq "github.com/rabbitmq/amqp091-go"
)

const (
	rmqAddressF     = "rmq-address"
	rmqConsumerTagF = "consumer-tag"

	defaultRmqAddress  = ""
	defaultConsumerTag = "da-consumer"
	defaultQueueName   = "da-mq"

	reconnectDelay = 3 * time.Second
	rechannelDelay = 2 * time.Second
)

type consumerData struct {
	addr        string
	consumerTag string
}

type Consumer struct {
	// TODO: via getter?
	QueueName string
	// TODO: via getter?
	ConsumerTag  string
	conn         *rmq.Connection
	onConnClosed <-chan *rmq.Error
	channel      *rmq.Channel
	onChanClosed <-chan *rmq.Error
}

func (consumer *Consumer) Reconnect(addr string) <-chan types.Block {
	for {
		conn, err := consumer.connect(addr)
		if err != nil {
			fmt.Println(err)
			<-time.After(reconnectDelay)
			continue
		}

		blockStream, rmqError := consumer.ResetChannel(conn)
		if rmqError != nil {
			fmt.Println(rmqError)
			continue
		}

		return blockStream
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
	consumer.conn = conn

	closeNotifier := make(chan *rmq.Error)
	consumer.onConnClosed = conn.NotifyClose(closeNotifier)
}

func (consumer *Consumer) ResetChannel(conn *rmq.Connection) (<-chan types.Block, *rmq.Error) {
	for {
		blockStream, err := consumer.setupChannel(conn)
		if err != nil {
			fmt.Println(err)

			select {
			case rmqError := <-consumer.onConnClosed:
				return nil, rmqError
			case <-time.After(rechannelDelay):
			}

			continue
		}

		return blockStream, nil
	}
}

func (consumer *Consumer) setupChannel(conn *rmq.Connection) (<-chan types.Block, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(consumer.QueueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	deliveries, err := channel.Consume(
		queue.Name,
		consumer.ConsumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	blockStream := make(chan types.Block)

	// TODO: improve logic
	// client on start right away returns a channel and
	// calls reconnect in backgroudn
	// Also fix situation with onConnClosed(not used anymore)
	go porcessDeliveries(blockStream, deliveries)

	consumer.changeChannel(channel)
	return blockStream, nil
}

func (consumer *Consumer) changeChannel(channel *rmq.Channel) {
	consumer.channel = channel

	closeNotifer := make(chan *rmq.Error)
	consumer.onChanClosed = channel.NotifyClose(closeNotifer)
}

func porcessDeliveries(blocksCh chan<- types.Block, deliveries <-chan rmq.Delivery) {
	defer close(blocksCh)

	for {
		d, ok := <-deliveries
		if !ok {
			break
		}
		var block types.Block

		// Decode block
		if err := rlp.DecodeBytes(d.Body, &block); err != nil {
			break
		}

		blocksCh <- block
	}
}

func parse() consumerData {
	addr := flag.String(rmqAddressF, defaultRmqAddress, "RMQ address(required)")
	consumerTag := flag.String(rmqConsumerTagF, defaultConsumerTag, "Consumer tag")

	flag.Parse()

	if *addr == "" {
		flag.Usage()
		panic("rmq-address is required")
	}

	return consumerData{
		addr:        *addr,
		consumerTag: *consumerTag,
	}
}

func main() {
	consumerData := parse()
	consumer := Consumer{
		QueueName:    defaultQueueName,
		ConsumerTag:  consumerData.consumerTag,
		conn:         nil,
		onConnClosed: nil,
		channel:      nil,
		onChanClosed: nil,
	}

	blockStream := consumer.Reconnect(consumerData.addr)
	for {
		block := <-blockStream
		fmt.Println(block)
	}
}
