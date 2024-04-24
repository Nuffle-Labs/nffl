package operator

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const OperatorNamespace = "sffl_operator"

type OperatorEventListener interface {
	OnTasksReceived()
}

type SelectiveOperatorListener struct {
	OnTasksReceivedCb func()
}

func (ol *SelectiveOperatorListener) OnTasksReceived() {
	if ol.OnTasksReceivedCb != nil {
		ol.OnTasksReceivedCb()
	}
}

type RpcClientEventListener interface {
	OnMessagesReceived()
}

type SelectiveRpcClientListener struct {
	OnMessagesReceivedCb func()
}

func (l *SelectiveRpcClientListener) OnMessagesReceived() {
	if l.OnMessagesReceivedCb != nil {
		l.OnMessagesReceivedCb()
	}
}

func MakeOperatorMetrics(registry *prometheus.Registry) OperatorEventListener {
	numTasksReceived := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "num_tasks_received",
			Help:      "The number of tasks received by reading from the avs service manager contract",
		})

	return &SelectiveOperatorListener{
		OnTasksReceivedCb: func() {
			numTasksReceived.Inc()
		},
	}
}

func MakeRpcClientMetrics(registry *prometheus.Registry) RpcClientEventListener {
	numMessagesReceived := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "num_messages_received",
			Help:      "The number of messages received by the operator set",
		})

	return &SelectiveRpcClientListener{
		OnMessagesReceivedCb: func() {
			numMessagesReceived.Inc()
		},
	}
}
