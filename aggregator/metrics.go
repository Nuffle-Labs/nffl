package aggregator

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const AggregatorNamespace = "sffl_aggregator"

type RpcEventListener interface {
	IncSignedCheckpointTaskResponse()
	IncSignedStateRootUpdateMessage()
	IncSignedOperatorSetUpdateMessage()
}

type SelectiveRpcListener struct {
	IncSignedCheckpointTaskResponseCb   func()
	IncSignedStateRootUpdateMessageCb   func()
	IncSignedOperatorSetUpdateMessageCb func()
}

func (rpcl *SelectiveRpcListener) IncSignedCheckpointTaskResponse() {
	if rpcl.IncSignedCheckpointTaskResponseCb != nil {
		rpcl.IncSignedCheckpointTaskResponseCb()
	}
}

func (rpcl *SelectiveRpcListener) IncSignedStateRootUpdateMessage() {
	if rpcl.IncSignedStateRootUpdateMessageCb != nil {
		rpcl.IncSignedStateRootUpdateMessageCb()
	}
}

func (rpcl *SelectiveRpcListener) IncSignedOperatorSetUpdateMessage() {
	if rpcl.IncSignedOperatorSetUpdateMessageCb != nil {
		rpcl.IncSignedOperatorSetUpdateMessageCb()
	}
}

type RestEventListener interface {
	IncStateRootUpdateRequests()
	IncOperatorSetUpdateRequests()
	IncCheckpointMessagesRequests()
	APIErrors()
}

type SelectiveRestListener struct {
	IncStateRootUpdateRequestsCb    func()
	IncOperatorSetUpdateRequestsCb  func()
	IncCheckpointMessagesRequestsCb func()
	APIErrorsCb                     func()
}

func (restl *SelectiveRestListener) IncStateRootUpdateRequests() {
	if restl.IncStateRootUpdateRequestsCb != nil {
		restl.IncStateRootUpdateRequestsCb()
	}
}

func (restl *SelectiveRestListener) IncOperatorSetUpdateRequests() {
	if restl.IncOperatorSetUpdateRequestsCb != nil {
		restl.IncOperatorSetUpdateRequestsCb()
	}
}

func (restl *SelectiveRestListener) IncCheckpointMessagesRequests() {
	if restl.IncCheckpointMessagesRequestsCb != nil {
		restl.IncCheckpointMessagesRequestsCb()
	}
}

func (restl *SelectiveRestListener) APIErrors() {
	if restl.APIErrorsCb != nil {
		restl.APIErrorsCb()
	}
}

func MakeRestServerMetrics(registry *prometheus.Registry) RestEventListener {
	stateRootUpdateRequests := promauto.With(registry).NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "state_root_update_requests_total",
		Help:      "Total number of state root update requests received",
	})
	operatorSetUpdateRequests := promauto.With(registry).NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "operator_set_update_requests_total",
		Help:      "Total number of operator set update requests received",
	})
	checkpointMessagesRequests := promauto.With(registry).NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "checkpoint_messages_requests_total",
		Help:      "Total number of checkpoint messages requests received",
	})
	apiErrors := promauto.With(registry).NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "api_errors_total",
		Help:      "Total number of API errors",
	})

	return &SelectiveRestListener{
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
	}
}

func MakeRpcServerMetrics(registry *prometheus.Registry) RpcEventListener {
	numSignedCheckpointTaskResponse := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "num_signed_checkpoints_accepted_by_aggregator",
			Help:      "The number of signed checkpoints responses accepted by the aggregator",
		},
	)

	numSignedStateRootUpdateMessage := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "num_signed_roots_accepted_by_aggregator",
			Help:      "The number of signed state roots accepted by the aggregator",
		},
	)

	numSignedOperatorSetUpdateMessage := promauto.With(registry).NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "num_signed_operators_accepted_by_aggregator",
			Help:      "The number of signed operator updates accepted by the aggregator",
		},
	)

	return &SelectiveRpcListener{
		IncSignedCheckpointTaskResponseCb: func() {
			numSignedCheckpointTaskResponse.Inc()
		},
		IncSignedStateRootUpdateMessageCb: func() {
			numSignedStateRootUpdateMessage.Inc()
		},
		IncSignedOperatorSetUpdateMessageCb: func() {
			numSignedOperatorSetUpdateMessage.Inc()
		},
	}
}
