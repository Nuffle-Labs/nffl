package main

import (
	"flag"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/NethermindEth/near-sffl/consumer"
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
		RollupIds:   []uint32{0},
	}
}

func main() {
	config := parse()
	logLevel := logging.Development
	logger, err := logging.NewZapLogger(logLevel)
	if err != nil {
		panic(err)
	}

	consumer := consumer.NewConsumer(config, logger)
	blockStream := consumer.GetBlockStream()

	for {
		block := <-blockStream
		fmt.Println(block)
	}
}
