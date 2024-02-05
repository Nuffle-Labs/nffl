package main

import (
	"flag"
	"fmt"

	"github.com/NethermindEth/near-sffl/comsumer"
)

const (
	rmqAddressF     = "rmq-address"
	rmqConsumerTagF = "consumer-tag"

	defaultRmqAddress  = ""
	defaultConsumerTag = "da-consumer"
)

func parse() consumer.ConsumerConfig {
	addr := flag.String(rmqAddressF, defaultRmqAddress, "RMQ address(required)")
	consumerTag := flag.String(rmqConsumerTagF, defaultConsumerTag, "Consumer tag")

	flag.Parse()

	if *addr == "" {
		flag.Usage()
		panic("rmq-address is required")
	}

	return consumer.ConsumerConfig{
		Addr:        *addr,
		ConsumerTag: *consumerTag,
	}
}

func main() {
	config := parse()
	consumer := consumer.NewConsumer(config)

	blockStream := consumer.GetBlockStream()

	for {
		block := <-blockStream
		fmt.Println(block)
	}
}
