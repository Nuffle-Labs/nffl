package operator

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
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

func MakeOperatorMetrics(registry *prometheus.Registry) (OperatorEventListener, error) {
	numTasksReceived := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "num_tasks_received",
			Help:      "The number of tasks received by reading from the avs service manager contract",
		})

	if err := registry.Register(numTasksReceived); err != nil {
		return nil, fmt.Errorf("error registering numTasksReceived counter: %w", err)
	}

	return &SelectiveOperatorListener{
		OnTasksReceivedCb: func() {
			numTasksReceived.Inc()
		},
	}, nil
}

func MakeRpcClientMetrics(registry *prometheus.Registry) (RpcClientEventListener, error) {
	numMessagesReceived := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "num_messages_received",
			Help:      "The number of messages received by the operator set",
		})

	if err := registry.Register(numMessagesReceived); err != nil {
		return nil, fmt.Errorf("error registering numMessagesReceived counter: %w", err)
	}

	return &SelectiveRpcClientListener{
		OnMessagesReceivedCb: func() {
			numMessagesReceived.Inc()
		},
	}, nil
}
