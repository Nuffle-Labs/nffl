package consumer

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
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
		l.OnArrivalCb()
	}
}

func (l *SelectiveListener) OnFormatError() {
	if l.OnFormatErrorCb != nil {
		l.OnFormatErrorCb()
	}
}

func MakeConsumerMetrics(registry *prometheus.Registry) (EventListener, error) {
	numBlocksArrived := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: ConsumerSubsystem,
			Name:      "num_of_mq_arrivals",
			Help:      "The number of consumed blocks from MQ",
		})

	if err := registry.Register(numBlocksArrived); err != nil {
		return nil, fmt.Errorf("error registering numBlocksArrived counter: %w", err)
	}

	numFormatErrors := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: ConsumerSubsystem,
			Name:      "num_of_mismatched_blocks",
			Help:      "The number of blocks from MQ with invalid format.",
		})

	if err := registry.Register(numFormatErrors); err != nil {
		return nil, fmt.Errorf("error registering numFormatErrors counter: %w", err)
	}

	return &SelectiveListener{
		OnArrivalCb: func() {
			numBlocksArrived.Inc()
		},
		OnFormatErrorCb: func() {
			numFormatErrors.Inc()
		},
	}, nil
}
