package rpc_server

import (
	"fmt"
	"github.com/Nuffle-Labs/nffl/aggregator"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	OperatorSetUpdateMessageLabel = "operator_set_update_message"
	StateRootUpdateMessageLabel   = "state_root_update_message"
	CheckpointTaskResponseLabel   = "checkpoint_task_response"
)

type EventListener interface {
	IncOperatorInitializations(operatorId [32]byte)
	IncSignedCheckpointTaskResponse(operatorId [32]byte, errored, notFound bool)
	IncSignedStateRootUpdateMessage(operatorId [32]byte, rollupId uint32, errored, hasNearDa bool)
	IncSignedOperatorSetUpdateMessage(operatorId [32]byte, errored bool)
	IncTotalSignedCheckpointTaskResponse()
	IncTotalSignedStateRootUpdateMessage()
	IncTotalSignedOperatorSetUpdateMessage()
	ObserveLastMessageReceivedTime(operatorId [32]byte, messageType string)
}

type SelectiveRpcListener struct {
	IncOperatorInitializationsCb             func(operatorId [32]byte)
	IncSignedCheckpointTaskResponseCb        func(operatorId [32]byte, errored, notFound bool)
	IncSignedStateRootUpdateMessageCb        func(operatorId [32]byte, rollupId uint32, errored, hasNearDa bool)
	IncSignedOperatorSetUpdateMessageCb      func(operatorId [32]byte, errored bool)
	IncTotalSignedCheckpointTaskResponseCb   func()
	IncTotalSignedStateRootUpdateMessageCb   func()
	IncTotalSignedOperatorSetUpdateMessageCb func()
	ObserveLastMessageReceivedTimeCb         func(operatorId [32]byte, messageType string)
}

func (l *SelectiveRpcListener) IncOperatorInitializations(operatorId [32]byte) {
	if l.IncOperatorInitializationsCb != nil {
		l.IncOperatorInitializationsCb(operatorId)
	}
}

func (l *SelectiveRpcListener) IncSignedCheckpointTaskResponse(operatorId [32]byte, errored, notFound bool) {
	if l.IncSignedCheckpointTaskResponseCb != nil {
		l.IncSignedCheckpointTaskResponseCb(operatorId, errored, notFound)
	}
}

func (l *SelectiveRpcListener) IncSignedStateRootUpdateMessage(operatorId [32]byte, rollupId uint32, errored, hasNearDa bool) {
	if l.IncSignedStateRootUpdateMessageCb != nil {
		l.IncSignedStateRootUpdateMessageCb(operatorId, rollupId, errored, hasNearDa)
	}
}

func (l *SelectiveRpcListener) IncSignedOperatorSetUpdateMessage(operatorId [32]byte, errored bool) {
	if l.IncSignedOperatorSetUpdateMessageCb != nil {
		l.IncSignedOperatorSetUpdateMessageCb(operatorId, errored)
	}
}

func (l *SelectiveRpcListener) ObserveLastMessageReceivedTime(operatorId [32]byte, messageType string) {
	if l.ObserveLastMessageReceivedTimeCb != nil {
		l.ObserveLastMessageReceivedTimeCb(operatorId, messageType)
	}
}

func (l *SelectiveRpcListener) IncTotalSignedCheckpointTaskResponse() {
	if l.IncTotalSignedCheckpointTaskResponseCb != nil {
		l.IncTotalSignedCheckpointTaskResponseCb()
	}
}

func (l *SelectiveRpcListener) IncTotalSignedStateRootUpdateMessage() {
	if l.IncTotalSignedStateRootUpdateMessageCb != nil {
		l.IncTotalSignedStateRootUpdateMessageCb()
	}
}

func (l *SelectiveRpcListener) IncTotalSignedOperatorSetUpdateMessage() {
	if l.IncTotalSignedOperatorSetUpdateMessageCb != nil {
		l.IncTotalSignedOperatorSetUpdateMessageCb()
	}
}

func MakeRpcServerMetrics(registry *prometheus.Registry) (EventListener, error) {
	operatorInitializationsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: aggregator.AggregatorNamespace,
			Name:      "operator_initializations_total",
			Help:      "Total number of operator initializations",
		},
		[]string{"operator_id"},
	)
	if err := registry.Register(operatorInitializationsTotal); err != nil {
		return nil, fmt.Errorf("error registering operatorInitializationsTotal counter: %w", err)
	}

	signedCheckpointTaskResponsesTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: aggregator.AggregatorNamespace,
			Name:      "signed_checkpoint_task_responses_total",
			Help:      "Total number of signed checkpoint task responses received per operator",
		},
		[]string{"operator_id", "errored", "not_found"},
	)
	if err := registry.Register(signedCheckpointTaskResponsesTotal); err != nil {
		return nil, fmt.Errorf("error registering signedCheckpointTaskResponsesTotal counter: %w", err)
	}

	signedStateRootUpdateMessagesTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: aggregator.AggregatorNamespace,
			Name:      "signed_state_root_update_messages_total",
			Help:      "Total number of signed state root update messages received per operator",
		},
		[]string{"operator_id", "rollup_id", "errored", "has_near_da"},
	)
	if err := registry.Register(signedStateRootUpdateMessagesTotal); err != nil {
		return nil, fmt.Errorf("error registering signedStateRootUpdateMessagesTotal counter: %w", err)
	}

	signedOperatorSetUpdateMessagesTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: aggregator.AggregatorNamespace,
			Name:      "signed_operator_set_update_messages_total",
			Help:      "Total number of signed operator set update messages received per operator",
		},
		[]string{"operator_id", "errored"},
	)
	if err := registry.Register(signedOperatorSetUpdateMessagesTotal); err != nil {
		return nil, fmt.Errorf("error registering signedOperatorSetUpdateMessagesTotal counter: %w", err)
	}

	lastMessageReceivedTime := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: aggregator.AggregatorNamespace,
			Name:      "last_message_received_time",
			Help:      "Timestamp of the last message received per operator and message type",
		},
		[]string{"operator_id", "message_type"},
	)
	if err := registry.Register(lastMessageReceivedTime); err != nil {
		return nil, fmt.Errorf("error registering lastMessageReceivedTime gauge: %w", err)
	}

	return &SelectiveRpcListener{
		IncOperatorInitializationsCb: func(operatorId [32]byte) {
			operatorInitializationsTotal.WithLabelValues(fmt.Sprintf("%x", operatorId)).Inc()
		},
		IncSignedCheckpointTaskResponseCb: func(operatorId [32]byte, errored, expired bool) {
			signedCheckpointTaskResponsesTotal.WithLabelValues(fmt.Sprintf("%x", operatorId), fmt.Sprintf("%t", errored), fmt.Sprintf("%t", expired)).Inc()
		},
		IncSignedStateRootUpdateMessageCb: func(operatorId [32]byte, rollupId uint32, errored, hasNearDa bool) {
			signedStateRootUpdateMessagesTotal.WithLabelValues(fmt.Sprintf("%x", operatorId), fmt.Sprintf("%d", rollupId), fmt.Sprintf("%t", errored), fmt.Sprintf("%t", hasNearDa)).Inc()
		},
		IncSignedOperatorSetUpdateMessageCb: func(operatorId [32]byte, errored bool) {
			signedOperatorSetUpdateMessagesTotal.WithLabelValues(fmt.Sprintf("%x", operatorId), fmt.Sprintf("%t", errored)).Inc()
		},
		ObserveLastMessageReceivedTimeCb: func(operatorId [32]byte, messageType string) {
			lastMessageReceivedTime.WithLabelValues(fmt.Sprintf("%x", operatorId), messageType).SetToCurrentTime()
		},
	}, nil
}
