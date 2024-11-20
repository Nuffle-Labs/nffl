package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli"

	"github.com/Nuffle-Labs/nffl/relayer"
	"github.com/Nuffle-Labs/nffl/relayer/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "sffl-test-relayer"
	app.Usage = "SFFL Test Relayer"
	app.Description = "Super Fast Finality testing service that reads block data from an EVM network and feeds it to a NEAR DA contract."
	app.Commands = []cli.Command{
		{
			Name:  "run-args",
			Usage: "Start the relayer with direct CLI options",
			Flags: []cli.Flag{
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
					Value: "http://127.0.0.1:3030",
					Usage: "Network for NEAR client to use (options: Mainnet, Testnet, Custom url, default: http://127.0.0.1:3030)",
				},
				cli.StringFlag{
					Name:  "metrics-ip-port-address",
					Value: "",
					Usage: "Metrics scrape address",
				},
			},
			Action: relayerMainFromArgs,
		},
		{
			Name:  "run-config",
			Usage: "Start the relayer using a configuration file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "path, p",
					Usage:    "Load configuration from `FILE`",
					Required: true,
				},
			},
			Action: relayerMainFromConfig,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func relayerMainFromArgs(ctx *cli.Context) error {
	config := config.RelayerConfig{
		Production:        ctx.Bool("production"),
		RpcUrl:            ctx.String("rpc-url"),
		DaAccountId:       ctx.String("da-account-id"),
		KeyPath:           ctx.String("key-path"),
		Network:           ctx.String("network"),
		MetricsIpPortAddr: ctx.String("metrics-ip-port-address"),
	}

	return relayerMain(config)
}

func relayerMainFromConfig(ctx *cli.Context) error {
	filePath := ctx.String("path")
	config := config.RelayerConfig{}
	if err := utils.ReadYamlConfig(filePath, &config); err != nil {
		return err
	}

	return relayerMain(config)
}

func relayerMain(config config.RelayerConfig) error {
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

	{
		logger.Info("Initializing Relayer")
		configJson, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			logger.Fatal(err.Error())
		}

		logger.Info("Read config", "config", string(configJson))
	}

	logger.Info("initializing relayer")
	rel, err := relayer.NewRelayerFromConfig(&config, logger)
	if err != nil {
		logger.Error("Error creating relayer", "err", err)
		return err
	}

	ctx := context.Background()
	if config.MetricsIpPortAddr != "" {
		registry := prometheus.NewRegistry()
		if err = rel.EnableMetrics(registry); err != nil {
			return err
		}

		relayer.StartMetricsServer(ctx, config.MetricsIpPortAddr, registry, logger)
	}

	logger.Info("initialized relayer")

	logger.Info("starting relayer")
	err = rel.Start(ctx)
	if err != nil {
		logger.Error("Error starting relayer", "err", err)
		return err
	}

	return nil
}
