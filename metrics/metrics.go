package metrics

import (
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	metrics.Metrics
	IncNumTasksReceived()
	IncNumTasksAcceptedByAggregator()
	IncNumMessagesReceived()
	IncNumMessagesAcceptedByAggregator()
}

// AvsMetrics contains instrumented metrics that should be incremented by the avs node using the methods below
type AvsAndEigenMetrics struct {
	metrics.Metrics
	// if numSignedTaskResponsesAcceptedByAggregator != numTasksReceived, then there is a bug
	numTasksReceived                           prometheus.Counter
	numSignedTaskResponsesAcceptedByAggregator prometheus.Counter

	numMessagesReceived                   prometheus.Counter
	numSignedMessagesAcceptedByAggregator prometheus.Counter
}

const superFastFinalityLayerNamespace = "sffl"

func NewAvsAndEigenMetrics(avsName string, eigenMetrics *metrics.EigenMetrics, reg prometheus.Registerer) *AvsAndEigenMetrics {
	return &AvsAndEigenMetrics{
		Metrics: eigenMetrics,
		numTasksReceived: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: superFastFinalityLayerNamespace,
				Name:      "num_tasks_received",
				Help:      "The number of tasks received by reading from the avs service manager contract",
			}),
		numSignedTaskResponsesAcceptedByAggregator: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: superFastFinalityLayerNamespace,
				Name:      "num_signed_task_responses_accepted_by_aggregator",
				Help:      "The number of signed task responses accepted by the aggregator",
			}),
		numMessagesReceived: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: superFastFinalityLayerNamespace,
				Name:      "num_messages_received",
				Help:      "The number of messages received by the operator set",
			}),
		numSignedMessagesAcceptedByAggregator: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: superFastFinalityLayerNamespace,
				Name:      "num_signed_messages_accepted_by_aggregator",
				Help:      "The number of signed messages accepted by the aggregator",
			}),
	}
}

func (m *AvsAndEigenMetrics) IncNumTasksReceived() {
	m.numTasksReceived.Inc()
}

func (m *AvsAndEigenMetrics) IncNumTasksAcceptedByAggregator() {
	m.numSignedTaskResponsesAcceptedByAggregator.Inc()
}

func (m *AvsAndEigenMetrics) IncNumMessagesReceived() {
	m.numMessagesReceived.Inc()
}

func (m *AvsAndEigenMetrics) IncNumMessagesAcceptedByAggregator() {
	m.numSignedMessagesAcceptedByAggregator.Inc()
}
