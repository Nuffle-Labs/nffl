package attestor

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type EventListener interface {
	OnMissedMQBlock()
	OnBlockMismatch()
}

const OperatorNamespace = "sffl_operator"
const AttestorSubsystem = "attestor"

type SelectiveEventListener struct {
	OnMissedMQBlockCb func()
	OnBlockMismatchCb func()
}

func (l *SelectiveEventListener) OnMissedMQBlock() {
	if l.OnMissedMQBlockCb != nil {
		l.OnBlockMismatchCb()
	}
}

func (l *SelectiveEventListener) OnBlockMismatch() {
	if l.OnBlockMismatchCb != nil {
		l.OnBlockMismatchCb()
	}
}

func MakeAttestorMetrics(registry *prometheus.Registry) (EventListener, error) {
	numMissedMqBlocks := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: AttestorSubsystem,
			Name:      "num_of_missed_mq_blocks",
			// TODO: late? better desc
			Help: "The number of late blocks from MQ",
		})

	if err := registry.Register(numMissedMqBlocks); err != nil {
		return nil, fmt.Errorf("error registering numMissedMqBlocks counter: %w", err)
	}

	numBlocksMismatched := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: AttestorSubsystem,
			Name:      "num_of_mismatched_blocks",
			Help:      "The number of blocks from MQ mismatched with RPC ones.",
		})

	if err := registry.Register(numBlocksMismatched); err != nil {
		return nil, fmt.Errorf("error registering numBlocksMismatched counter: %w", err)
	}

	return &SelectiveEventListener{
		OnMissedMQBlockCb: func() {
			numMissedMqBlocks.Inc()
		},
		OnBlockMismatchCb: func() {
			numBlocksMismatched.Inc()
		},
	}, nil
}
