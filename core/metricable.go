package core

import "github.com/prometheus/client_golang/prometheus"

type Metricable interface {
	WithMetrics(registry *prometheus.Registry) error
}
