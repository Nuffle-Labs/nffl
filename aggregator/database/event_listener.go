package database

import (
	"math"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const AggregatorNamespace = "sffl_aggregator"
const DBSubsystem = "db"

type EventListener interface {
	OnStore(duration time.Duration)
	OnFetch(duration time.Duration)
}

type SelectiveListener struct {
	OnStoreCb func(duration time.Duration)
	OnFetchCb func(duration time.Duration)
}

func (l *SelectiveListener) OnStore(duration time.Duration) {
	if l.OnStoreCb != nil {
		l.OnStoreCb(duration)
	}
}

func (l *SelectiveListener) OnFetch(duration time.Duration) {
	if l.OnFetchCb != nil {
		l.OnFetchCb(duration)
	}
}

func MakeDBMetrics(registry *prometheus.Registry) EventListener {
	latencyBuckets := []float64{
		25,
		50,
		75,
		100,
		250,
		500,
		1000, // 1ms
		2000,
		3000,
		4000,
		5000,
		10000,
		50000,
		500000,
		math.Inf(0),
	}
	readLatencyHistogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: AggregatorNamespace,
		Subsystem: DBSubsystem,
		Name:      "read_latency",
		Buckets:   latencyBuckets,
	})
	writeLatencyHistogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: AggregatorNamespace,
		Subsystem: DBSubsystem,
		Name:      "write_latency",
		Buckets:   latencyBuckets,
	})

	registry.MustRegister(readLatencyHistogram, writeLatencyHistogram)
	return &SelectiveListener{
		OnStoreCb: func(duration time.Duration) {
			writeLatencyHistogram.Observe(float64(duration.Microseconds()))
		},
		OnFetchCb: func(duration time.Duration) {
			readLatencyHistogram.Observe(float64(duration.Microseconds()))
		},
	}
}
