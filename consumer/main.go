package main

import (
	"context"
	"flag"
	"fmt"

	rmq "github.com/rabbitmq/amqp091-go"
)

const (
	rmqAddressF     = "rmq-address"
	rmqConsumerTagF = "consumer-tag"

	defaultRmqAddress  = ""
	defaultConsumerTag = "da-consumer"
	defaultQueueName   = "da-mq"
)

type consumerData struct {
	addr        string
	consumerTag string
}

func parse() consumerData {
	addr := flag.String(rmqAddressF, defaultRmqAddress, "RMQ address(required)")
	consumerTag := flag.String(rmqConsumerTagF, defaultConsumerTag, "Consumer tag")

	flag.Parse()

	if *addr == "" {
		fmt.Println("rmq-address is required")
		flag.Usage()
		panic("")
	}

	return consumerData{
		addr:        *addr,
		consumerTag: *consumerTag,
	}
}

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

	conn, err := rmq.Dial(consumerData.addr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(defaultQueueName, true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	deliveries, err := channel.Consume(
		queue.Name,
		consumerData.consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(queue.Name)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	dataStream := make(chan []byte, 10)
	go porcessDeliveries(dataStream, ctx)

	for d := range deliveries {
		dataStream <- d.Body
		d.Ack(true)
	}
}
