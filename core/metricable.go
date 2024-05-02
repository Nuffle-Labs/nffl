package core

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/prometheus/client_golang/prometheus"
)

type Metricable interface {
	EnableMetrics(registry *prometheus.Registry) error
}

func CreateEthClient(rpcUrl string, collector *rpccalls.Collector, logger sdklogging.Logger) (eth.Client, error) {
	if collector != nil {
		ethClient, err := eth.NewInstrumentedClient(rpcUrl, collector)
		if err != nil {
			logger.Error("Cannot create ethclient", "err", err)
			return nil, err
		}

		return ethClient, nil
	}

	ethClient, err := eth.NewClient(rpcUrl)
	if err != nil {
		logger.Error("Cannot create ethclient", "err", err)
		return nil, err
	}

	return ethClient, nil
}

func CreateEthClientWithCollector(id, url string, enableMetrics bool, registry *prometheus.Registry, logger sdklogging.Logger) (eth.Client, error) {
	if enableMetrics {
		// Using url as avsName
		rpcCallsCollector := rpccalls.NewCollector(id+url, registry)
		return CreateEthClient(url, rpcCallsCollector, logger)
	}

	return CreateEthClient(url, nil, logger)
}
