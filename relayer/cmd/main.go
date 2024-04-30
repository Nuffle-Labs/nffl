package main

import (
	"context"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/urfave/cli"

	"github.com/NethermindEth/near-sffl/relayer"
	"github.com/NethermindEth/near-sffl/relayer/config"
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
			Value: "http://127.0.0.1:3030",
			Usage: "Network for NEAR client to use (options: Mainnet, Testnet, Custom url, default: http://127.0.0.1:3030)",
		},
		cli.StringFlag{
			Name:  "metrics-addr",
			Value: "",
			Usage: "Metrics scrape address",
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

func startMetrics(metricsAddr string, reg prometheus.Gatherer) (<-chan error, func()) {
	errC := make(chan error, 1)
	server := &http.Server{Addr: metricsAddr, Handler: nil}

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	shutdown := func() {
		if err := server.Shutdown(context.Background()); err != nil {
			// Handle the error according to your application's needs, e.g., log it
			log.Printf("Error shutting down metrics server: %v", err)
		}
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, shutdown
}

func relayerMain(ctx *cli.Context) error {
	config := &config.RelayerConfig{
		Production:  ctx.GlobalBool("production"),
		RpcUrl:      ctx.GlobalString("rpc-url"),
		DaAccountId: ctx.GlobalString("da-account-id"),
		KeyPath:     ctx.GlobalString("key-path"),
		Network:     ctx.GlobalString("network"),
		MetricsAddr: ctx.GlobalString("metrics-addr"),
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

	{
		logger.Info("Initializing Relayer")
		configJson, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			logger.Fatal(err.Error())
		}

		logger.Infof("Config: %s", string(configJson))
	}

	logger.Info("initializing relayer")
	rel, err := relayer.NewRelayerFromConfig(config, logger)
	if err != nil {
		logger.Error("Error creating relayer", "err", err)
		return err
	}

	if config.MetricsAddr != "" {
		registry := prometheus.NewRegistry()
		if err = rel.WithMetrics(registry); err != nil {
			return err
		}

		_, shutdown := startMetrics(config.MetricsAddr, registry)
		defer shutdown()
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
