package core

import (
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/NethermindEth/near-sffl/core/safeclient"
)

type Metricable interface {
	EnableMetrics(registry *prometheus.Registry) error
}

func CreateEthClientWithCollector(id, url string, enableMetrics bool, registry *prometheus.Registry, logger sdklogging.Logger) (safeclient.SafeClient, error) {
	if enableMetrics {
		// Using url as avsName
		rpcCallsCollector := rpccalls.NewCollector(id+url, registry)
		return safeclient.NewSafeEthClient(url, logger, safeclient.WithCollector(rpcCallsCollector))
	}

	return safeclient.NewSafeEthClient(url, logger)
}
