package main

import (
	"context"
	"encoding/json"
	"fmt"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/NethermindEth/near-sffl/aggregator"
	"github.com/NethermindEth/near-sffl/core/config"
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

	rawConfig := config.NewRawConfig(ctx)
	logger, err := sdklogging.NewZapLogger(rawConfig.Environment)
	if err != nil {
		return err
	}

	config, err := config.NewConfig(rawConfig, ctx)
	if err != nil {
		logger.Error("Error creating config", "err", err)
		return err
	}

	{
		configJson, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			logger.Fatalf(err.Error())
		}
		fmt.Println("Config:", string(configJson))
	}

	bgCtx := context.Background()
	agg, err := aggregator.NewAggregator(config, logger, bgCtx)
	if err != nil {
		return err
	}

	err = agg.Start(bgCtx)
	if err != nil {
		return err
	}

	return nil
}
