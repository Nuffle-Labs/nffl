package attestor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

func MakeAttestorMetrics(registry *prometheus.Registry) EventListener {
	numMissedMqBlocks := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			// TODO: different namespace?
			Namespace: OperatorNamespace,
			Subsystem: AttestorSubsystem,
			Name:      "num_of_missed_mq_blocks",
			// TODO: late? better desc
			Help: "The number of late blocks from MQ",
		})

	numBlocksMismatched := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: AttestorSubsystem,
			Name:      "num_of_mismatched_blocks",
			Help:      "The number of blocks from MQ mismatched with RPC ones.",
		})

	return &SelectiveEventListener{
		OnMissedMQBlockCb: func() {
			numMissedMqBlocks.Inc()
		},
		OnBlockMismatchCb: func() {
			numBlocksMismatched.Inc()
		},
	}
}
