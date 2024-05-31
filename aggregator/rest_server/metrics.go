package rest_server

import (
	"fmt"

	"github.com/NethermindEth/near-sffl/aggregator"
	"github.com/prometheus/client_golang/prometheus"
)

type EventListener interface {
	IncStateRootUpdateRequests()
	IncOperatorSetUpdateRequests()
	IncCheckpointMessagesRequests()
	APIErrors()
}

type SelectiveListener struct {
	IncStateRootUpdateRequestsCb    func()
	IncOperatorSetUpdateRequestsCb  func()
	IncCheckpointMessagesRequestsCb func()
	APIErrorsCb                     func()
}

func (l *SelectiveListener) IncStateRootUpdateRequests() {
	if l.IncStateRootUpdateRequestsCb != nil {
		l.IncStateRootUpdateRequestsCb()
	}
}

func (l *SelectiveListener) IncOperatorSetUpdateRequests() {
	if l.IncOperatorSetUpdateRequestsCb != nil {
		l.IncOperatorSetUpdateRequestsCb()
	}
}

func (l *SelectiveListener) IncCheckpointMessagesRequests() {
	if l.IncCheckpointMessagesRequestsCb != nil {
		l.IncCheckpointMessagesRequestsCb()
	}
}

func (l *SelectiveListener) APIErrors() {
	if l.APIErrorsCb != nil {
		l.APIErrorsCb()
	}
}

func MakeServerMetrics(registry *prometheus.Registry) (EventListener, error) {
	stateRootUpdateRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: aggregator.AggregatorNamespace,
		Name:      "state_root_update_requests_total",
		Help:      "Total number of state root update requests received",
	})
	if err := registry.Register(stateRootUpdateRequests); err != nil {
		return nil, fmt.Errorf("error registering stateRootUpdateRequests counter: %w", err)
	}

	operatorSetUpdateRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: aggregator.AggregatorNamespace,
		Name:      "operator_set_update_requests_total",
		Help:      "Total number of operator set update requests received",
	})
	if err := registry.Register(operatorSetUpdateRequests); err != nil {
		return nil, fmt.Errorf("error registering operatorSetUpdateRequests counter: %w", err)
	}

	checkpointMessagesRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: aggregator.AggregatorNamespace,
		Name:      "checkpoint_messages_requests_total",
		Help:      "Total number of checkpoint messages requests received",
	})
	if err := registry.Register(checkpointMessagesRequests); err != nil {
		return nil, fmt.Errorf("error registering checkpointMessagesRequests counter: %w", err)
	}

	apiErrors := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: aggregator.AggregatorNamespace,
		Name:      "api_errors_total",
		Help:      "Total number of API errors",
	})
	if err := registry.Register(apiErrors); err != nil {
		return nil, fmt.Errorf("error registering apiErrors counter: %w", err)
	}

	return &SelectiveListener{
		IncStateRootUpdateRequestsCb: func() {
			stateRootUpdateRequests.Inc()
		},
		IncOperatorSetUpdateRequestsCb: func() {
			operatorSetUpdateRequests.Inc()
		},
		IncCheckpointMessagesRequestsCb: func() {
			checkpointMessagesRequests.Inc()
		},
		APIErrorsCb: func() {
			apiErrors.Inc()
		},
	}, nil
}
