package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/NethermindEth/near-sffl/relayer"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "production",
			Usage: "Run in production logging mode",
		},
		cli.StringFlag{
			Name:     "rpc-url",
			Required: true,
			Usage:    "Connect to the indicated RPC`",
		},
		cli.StringFlag{
			Name:     "da-account-id",
			Required: true,
			Usage:    "Publish block data to the indicated NEAR account",
		},
	}
	app.Name = "sffl-test-relayer"
	app.Usage = "SFFL Test Relayer"
	app.Description = "Super Fast Finality testing service that reads block data from an EVM network and feeds it to a NEAR DA contract."

	app.Action = relayerMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func relayerMain(ctx *cli.Context) error {
	log.Println("Initializing Relayer")

	config := &relayer.RelayerConfig{
		Production:  ctx.GlobalBool("production"),
		RpcUrl:      ctx.GlobalString("rpc-url"),
		DaAccountId: ctx.GlobalString("da-account-id"),
	}

	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Config:", string(configJson))

	log.Println("initializing relayer")
	rel, err := relayer.NewRelayerFromConfig(config)
	if err != nil {
		return err
	}
	log.Println("initialized relayer")

	log.Println("starting relayer")
	err = rel.Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}
