package main

import (
	"context"
	"flag"
	"fmt"
	"time"

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

func (consumer *Consumer) Reconnect(addr string) <-chan rmq.Delivery {
	for {
		conn, err := consumer.connect(addr)
		if err != nil {
			fmt.Println(err)
			<-time.After(reconnectDelay)
			continue
		}

		deliveries, rmqError := consumer.ResetChannel(conn)
		if rmqError != nil {
			fmt.Println(rmqError)
			continue
		}

		return deliveries
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

func (consumer *Consumer) ResetChannel(conn *rmq.Connection) (<-chan rmq.Delivery, *rmq.Error) {
	for {
		deliveries, err := consumer.setupChannel(conn)
		if err != nil {
			fmt.Println(err)

			select {
			case rmqError := <-consumer.onConnClosed:
				return nil, rmqError
			case <-time.After(rechannelDelay):
			}

			continue
		}

		return deliveries, nil
	}
}

func (consumer *Consumer) setupChannel(conn *rmq.Connection) (<-chan rmq.Delivery, error) {
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

	consumer.changeChannel(channel)
	return deliveries, nil
}

func (consumer *Consumer) changeChannel(channel *rmq.Channel) {
	consumer.channel = channel

	closeNotifer := make(chan *rmq.Error)
	consumer.onChanClosed = channel.NotifyClose(closeNotifer)
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

// TODO
func porcessDeliveries(stream <-chan []byte, ctx context.Context) {
	for {
		select {
		case data := <-stream:
			fmt.Println("accepted data", data)
		case <-ctx.Done():
			fmt.Println("done")
			break
		}
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
	deliveries := consumer.Reconnect(consumerData.addr)

	ctx, cancel := context.WithCancel(context.TODO())

	dataStream := make(chan []byte, 10)
	go porcessDeliveries(dataStream, ctx)

	for {
		select {
		case d := <-deliveries:
			dataStream <- d.Body
			d.Ack(true)

		case rmqError := <-consumer.onChanClosed:
			fmt.Println(rmqError)
			if !rmqError.Recover {
				defer cancel()
				break
			}

			deliveries = consumer.Reconnect(consumerData.addr)

		case rmqError := <-consumer.onConnClosed:
			fmt.Println(rmqError)
			if !rmqError.Recover {
				defer cancel()
				break
			}

			deliveries = consumer.Reconnect(consumerData.addr)
		}
	}
}
