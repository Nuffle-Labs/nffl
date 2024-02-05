package main

import (
	"errors"
	"log"
	"os"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/urfave/cli"

	"github.com/NethermindEth/near-sffl/consumer"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "rmq-address",
			Required: true,
			Usage:    "Connect to RMQ publisher at `ADDRESS`",
		},
		cli.StringFlag{
			Name:  "consumer-tag",
			Value: "da-consumer",
			Usage: "Connect to RMQ publisher using `TAG`",
		},
		cli.Int64SliceFlag{
			Name:     "rollup-ids",
			Required: true,
			Usage:    "Consume data from rollup `ID`",
		},
	}
	app.Name = "sffl-indexer-consumer"
	app.Usage = "SFFL Indexer Consumer"
	app.Description = "Super Fast Finality Layer test service to consume NEAR DA published block data from the indexer"

	app.Action = consumerMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func consumerMain(ctx *cli.Context) error {
	log.Println("Initializing Consumer")

	logLevel := logging.Development
	logger, err := logging.NewZapLogger(logLevel)
	if err != nil {
		panic(err)
	}

	rollupIdsArg := ctx.GlobalInt64Slice("rollup-ids")
	rollupIds := make([]uint32, len(rollupIdsArg))
	for i, el := range rollupIdsArg {
		if el < 0 {
			return errors.New("Rollup IDs should not be < 0")
		}

		rollupIds[i] = uint32(el)
	}

	consumer := consumer.NewConsumer(consumer.ConsumerConfig{
		Addr:        ctx.GlobalString("rmq-address"),
		ConsumerTag: ctx.GlobalString("consumer-tag"),
		RollupIds:   rollupIds,
	}, logger)

	blockStream := consumer.GetBlockStream()

	for {
		block := <-blockStream
		logger.Info("Block received", "block", block)
	}
}
