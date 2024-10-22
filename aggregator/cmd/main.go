package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli"

	"github.com/Nuffle-Labs/nffl/aggregator"
	restserver "github.com/Nuffle-Labs/nffl/aggregator/rest_server"
	rpcserver "github.com/Nuffle-Labs/nffl/aggregator/rpc_server"
	"github.com/Nuffle-Labs/nffl/core/config"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {
	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "sffl"
	app.Usage = "SFFL Aggregator"
	app.Description = "Service that sends checkpoints to be signed by operator nodes."

	app.Action = aggregatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func aggregatorMain(ctx *cli.Context) error {
	log.Println("Initializing Aggregator")

	configRaw, err := config.NewConfigRaw(ctx)
	if err != nil {
		return err
	}

	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	if err != nil {
		return err
	}

	config, err := config.NewConfig(ctx, *configRaw, logger)
	if err != nil {
		return err
	}

	// Print config as JSON
	{
		configJson, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			logger.Fatalf(err.Error())
		}
		fmt.Println("Config:", string(configJson))
	}

	bgCtx := context.Background()
	var optRegistry *prometheus.Registry
	if config.EnableMetrics {
		optRegistry = prometheus.NewRegistry()
	}
	agg, err := aggregator.NewAggregator(bgCtx, config, optRegistry, logger)
	if err != nil {
		return err
	}

	rpcServer := rpcserver.NewRpcServer(config.AggregatorServerIpPortAddr, agg, logger)
	if optRegistry != nil {
		if err = rpcServer.EnableMetrics(optRegistry); err != nil {
			return err
		}
	}
	go rpcServer.Start()

	restServer := restserver.NewRestServer(config.AggregatorRestServerIpPortAddr, agg, logger)
	if optRegistry != nil {
		if err = restServer.EnableMetrics(optRegistry); err != nil {
			return err
		}
	}
	go restServer.Start()

	err = agg.Start(bgCtx)
	if err != nil {
		return err
	}

	return nil
}
