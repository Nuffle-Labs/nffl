package consumer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type EventListener interface {
	OnArrival()
	OnFormatError()
}

const OperatorNamespace = "sffl_operator"
const ConsumerSubsystem = "consumer"

type SelectiveListener struct {
	OnArrivalCb     func()
	OnFormatErrorCb func()
}

func (l *SelectiveListener) OnArrival() {
	if l.OnArrivalCb != nil {
		l.OnArrival()
	}
}

func (l *SelectiveListener) OnFormatError() {
	if l.OnFormatErrorCb != nil {
		l.OnFormatError()
	}
}

func MakeConsumerMetrics(registry *prometheus.Registry) EventListener {
	numBlocksArrived := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			// TODO: different namespace?
			Namespace: OperatorNamespace,
			Subsystem: ConsumerSubsystem,
			Name:      "num_of_mq_arrivals",
			Help:      "The number of consumed blocks from MQ",
		})

	numFormatErrors := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: ConsumerSubsystem,
			Name:      "num_of_mismatched_blocks",
			Help:      "The number of blocks from MQ with invalid format.",
		})

	return &SelectiveListener{
		OnArrivalCb: func() {
			numBlocksArrived.Inc()
		},
		OnFormatErrorCb: func() {
			numFormatErrors.Inc()
		},
	}
}
