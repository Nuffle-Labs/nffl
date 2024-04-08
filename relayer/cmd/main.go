package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/urfave/cli"
	
	"github.com/NethermindEth/near-sffl/relayer"
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
			Usage:    "Connect to the indicated RPC",
		},
		cli.StringFlag{
			Name:     "da-account-id",
			Required: true,
			Usage:    "Publish block data to the indicated NEAR account",
		},
		cli.StringFlag{
			Name:     "key-path",
			Required: true,
			Usage:    "Path to NEAR account's key file",
		},
		cli.StringFlag{
			Name:  "network",
			Value: "127.0.0.1:3030",
			Usage: "Network for NEAR client to use (options: Mainnet, Testnet, [ip]:[port], default: 127.0.0.1:3030)",
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
	config := &relayer.RelayerConfig{
		Production:  ctx.GlobalBool("production"),
		RpcUrl:      ctx.GlobalString("rpc-url"),
		DaAccountId: ctx.GlobalString("da-account-id"),
		KeyPath:     ctx.GlobalString("key-path"),
		Network:     ctx.GlobalString("network"),
	}

	var logLevel sdklogging.LogLevel
	if config.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}

	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return err
	}

	logger.Info("Initializing Relayer")
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Infof("Config: %s", string(configJson))

	logger.Info("initializing relayer")
	rel, err := relayer.NewRelayerFromConfig(config, logger)
	if err != nil {
		logger.Error("Error creating relayer", "err", err)
		return err
	}

	logger.Info("initialized relayer")

	logger.Info("starting relayer")
	err = rel.Start(context.Background())
	if err != nil {
		logger.Error("Error starting relayer", "err", err)
		return err
	}

	return nil
}
