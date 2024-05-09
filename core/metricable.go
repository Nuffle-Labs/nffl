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

func CreateEthClientWithCollector(id, url string, enableMetrics bool, registry *prometheus.Registry, logger sdklogging.Logger) (eth.Client, error) {
	if enableMetrics {
		// Using url as avsName
		rpcCallsCollector := rpccalls.NewCollector(id+url, registry)
		return NewSafeEthClient(url, logger, WithCollector(rpcCallsCollector))
	}

	return NewSafeEthClient(url, logger)
}
