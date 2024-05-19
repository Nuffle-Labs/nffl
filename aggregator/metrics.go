package aggregator

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	AggregatorNamespace           = "sffl_aggregator"
	StateRootUpdateMessageLabel   = "state_root_update_message"
	OperatorSetUpdateMessageLabel = "operator_set_update_message"
	CheckpointTaskResponseLabel   = "checkpoint_task_response"
)

type RpcEventListener interface {
	IncSignedCheckpointTaskResponse(operatorId [32]byte, errored, notFound bool)
	IncSignedStateRootUpdateMessage(operatorId [32]byte, errored, hasNearDa bool)
	IncSignedOperatorSetUpdateMessage(operatorId [32]byte, errored bool)
	IncTotalSignedCheckpointTaskResponse()
	IncTotalSignedStateRootUpdateMessage()
	IncTotalSignedOperatorSetUpdateMessage()
	ObserveLastMessageReceivedTime(operatorId [32]byte, messageType string)
}

type SelectiveRpcListener struct {
	IncSignedCheckpointTaskResponseCb        func(operatorId [32]byte, errored, notFound bool)
	IncSignedStateRootUpdateMessageCb        func(operatorId [32]byte, errored, hasNearDa bool)
	IncSignedOperatorSetUpdateMessageCb      func(operatorId [32]byte, errored bool)
	IncTotalSignedCheckpointTaskResponseCb   func()
	IncTotalSignedStateRootUpdateMessageCb   func()
	IncTotalSignedOperatorSetUpdateMessageCb func()
	ObserveLastMessageReceivedTimeCb         func(operatorId [32]byte, messageType string)
}

func (l *SelectiveRpcListener) IncSignedCheckpointTaskResponse(operatorId [32]byte, errored, notFound bool) {
	if l.IncSignedCheckpointTaskResponseCb != nil {
		l.IncSignedCheckpointTaskResponseCb(operatorId, errored, notFound)
	}
}

func (l *SelectiveRpcListener) IncSignedStateRootUpdateMessage(operatorId [32]byte, errored, hasNearDa bool) {
	if l.IncSignedStateRootUpdateMessageCb != nil {
		l.IncSignedStateRootUpdateMessageCb(operatorId, errored, hasNearDa)
	}
}

func (l *SelectiveRpcListener) IncSignedOperatorSetUpdateMessage(operatorId [32]byte, errored bool) {
	if l.IncSignedOperatorSetUpdateMessageCb != nil {
		l.IncSignedOperatorSetUpdateMessageCb(operatorId, errored)
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

func (l *SelectiveRpcListener) ObserveLastMessageReceivedTime(operatorId [32]byte, messageType string) {
	if l.ObserveLastMessageReceivedTimeCb != nil {
		l.ObserveLastMessageReceivedTimeCb(operatorId, messageType)
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

func (l *SelectiveRestListener) IncStateRootUpdateRequests() {
	if l.IncStateRootUpdateRequestsCb != nil {
		l.IncStateRootUpdateRequestsCb()
	}
}

func (l *SelectiveRestListener) IncOperatorSetUpdateRequests() {
	if l.IncOperatorSetUpdateRequestsCb != nil {
		l.IncOperatorSetUpdateRequestsCb()
	}
}

func (l *SelectiveRestListener) IncCheckpointMessagesRequests() {
	if l.IncCheckpointMessagesRequestsCb != nil {
		l.IncCheckpointMessagesRequestsCb()
	}
}

func (l *SelectiveRestListener) APIErrors() {
	if l.APIErrorsCb != nil {
		l.APIErrorsCb()
	}
}

type AggregatorEventListener interface {
	ObserveLastStateRootUpdateAggregated(rollupId uint32, blockNumber uint64)
	ObserveLastStateRootUpdateReceived(rollupId uint32, blockNumber uint64)
	ObserveLastOperatorSetUpdateAggregated(operatorSetUpdateId uint64)
	ObserveLastOperatorSetUpdateReceived(operatorSetUpdateId uint64)
	IncExpiredMessages()
	IncExpiredTasks()
	IncErroredSubmissions()
	IncAggregatorInitializations()
	ObserveLastCheckpointReferenceSent(referenceId uint32)
	ObserveLastCheckpointTaskReferenceReceived(referenceId uint32)
	ObserveLastCheckpointTaskReferenceAggregated(referenceId uint32)
}

type SelectiveAggregatorListener struct {
	ObserveLastStateRootUpdateAggregatedCb         func(rollupId uint32, blockNumber uint64)
	ObserveLastStateRootUpdateReceivedCb           func(rollupId uint32, blockNumber uint64)
	ObserveLastOperatorSetUpdateAggregatedCb       func(operatorSetUpdateId uint64)
	ObserveLastOperatorSetUpdateReceivedCb         func(operatorSetUpdateId uint64)
	IncExpiredMessagesCb                           func()
	IncExpiredTasksCb                              func()
	IncErroredSubmissionsCb                        func()
	IncAggregatorInitializationsCb                 func()
	ObserveLastCheckpointReferenceSentCb           func(referenceId uint32)
	ObserveLastCheckpointTaskReferenceReceivedCb   func(referenceId uint32)
	ObserveLastCheckpointTaskReferenceAggregatedCb func(referenceId uint32)
}

func (l *SelectiveAggregatorListener) ObserveLastOperatorSetUpdateAggregated(operatorSetUpdateId uint64) {
	if l.ObserveLastOperatorSetUpdateAggregatedCb != nil {
		l.ObserveLastOperatorSetUpdateAggregatedCb(operatorSetUpdateId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastOperatorSetUpdateReceived(operatorSetUpdateId uint64) {
	if l.ObserveLastOperatorSetUpdateReceivedCb != nil {
		l.ObserveLastOperatorSetUpdateReceivedCb(operatorSetUpdateId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastStateRootUpdateAggregated(rollupId uint32, blockNumber uint64) {
	if l.ObserveLastStateRootUpdateAggregatedCb != nil {
		l.ObserveLastStateRootUpdateAggregatedCb(rollupId, blockNumber)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastStateRootUpdateReceived(rollupId uint32, blockNumber uint64) {
	if l.ObserveLastStateRootUpdateReceivedCb != nil {
		l.ObserveLastStateRootUpdateReceivedCb(rollupId, blockNumber)
	}
}

func (l *SelectiveAggregatorListener) IncExpiredMessages() {
	if l.IncExpiredMessagesCb != nil {
		l.IncExpiredMessagesCb()
	}
}

func (l *SelectiveAggregatorListener) IncExpiredTasks() {
	if l.IncExpiredTasksCb != nil {
		l.IncExpiredTasksCb()
	}
}

func (l *SelectiveAggregatorListener) IncErroredSubmissions() {
	if l.IncErroredSubmissionsCb != nil {
		l.IncErroredSubmissionsCb()
	}
}

func (l *SelectiveAggregatorListener) IncAggregatorInitializations() {
	if l.IncAggregatorInitializationsCb != nil {
		l.IncAggregatorInitializationsCb()
	}
}

func (l *SelectiveAggregatorListener) ObserveLastCheckpointReferenceSent(referenceId uint32) {
	if l.ObserveLastCheckpointReferenceSentCb != nil {
		l.ObserveLastCheckpointReferenceSentCb(referenceId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastCheckpointTaskReferenceReceived(referenceId uint32) {
	if l.ObserveLastCheckpointTaskReferenceReceivedCb != nil {
		l.ObserveLastCheckpointTaskReferenceReceivedCb(referenceId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastCheckpointTaskReferenceAggregated(referenceId uint32) {
	if l.ObserveLastCheckpointTaskReferenceAggregatedCb != nil {
		l.ObserveLastCheckpointTaskReferenceAggregatedCb(referenceId)
	}
}

func MakeRestServerMetrics(registry *prometheus.Registry) (RestEventListener, error) {
	stateRootUpdateRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "state_root_update_requests_total",
		Help:      "Total number of state root update requests received",
	})
	if err := registry.Register(stateRootUpdateRequests); err != nil {
		return nil, fmt.Errorf("error registering stateRootUpdateRequests counter: %w", err)
	}

	operatorSetUpdateRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "operator_set_update_requests_total",
		Help:      "Total number of operator set update requests received",
	})
	if err := registry.Register(operatorSetUpdateRequests); err != nil {
		return nil, fmt.Errorf("error registering operatorSetUpdateRequests counter: %w", err)
	}

	checkpointMessagesRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "checkpoint_messages_requests_total",
		Help:      "Total number of checkpoint messages requests received",
	})
	if err := registry.Register(checkpointMessagesRequests); err != nil {
		return nil, fmt.Errorf("error registering checkpointMessagesRequests counter: %w", err)
	}

	apiErrors := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: AggregatorNamespace,
		Name:      "api_errors_total",
		Help:      "Total number of API errors",
	})
	if err := registry.Register(apiErrors); err != nil {
		return nil, fmt.Errorf("error registering apiErrors counter: %w", err)
	}

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
	}, nil
}

func MakeRpcServerMetrics(registry *prometheus.Registry) (RpcEventListener, error) {
	signedCheckpointTaskResponsesTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
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
			Namespace: AggregatorNamespace,
			Name:      "signed_state_root_update_messages_total",
			Help:      "Total number of signed state root update messages received per operator",
		},
		[]string{"operator_id", "errored", "has_near_da"},
	)
	if err := registry.Register(signedStateRootUpdateMessagesTotal); err != nil {
		return nil, fmt.Errorf("error registering signedStateRootUpdateMessagesTotal counter: %w", err)
	}

	signedOperatorSetUpdateMessagesTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
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
			Namespace: AggregatorNamespace,
			Name:      "last_message_received_time",
			Help:      "Timestamp of the last message received per operator and message type",
		},
		[]string{"operator_id", "message_type"},
	)
	if err := registry.Register(lastMessageReceivedTime); err != nil {
		return nil, fmt.Errorf("error registering lastMessageReceivedTime gauge: %w", err)
	}

	return &SelectiveRpcListener{
		IncSignedCheckpointTaskResponseCb: func(operatorId [32]byte, errored, expired bool) {
			signedCheckpointTaskResponsesTotal.WithLabelValues(fmt.Sprintf("%x", operatorId), fmt.Sprintf("%t", errored), fmt.Sprintf("%t", expired)).Inc()
		},
		IncSignedStateRootUpdateMessageCb: func(operatorId [32]byte, errored, hasNearDa bool) {
			signedStateRootUpdateMessagesTotal.WithLabelValues(fmt.Sprintf("%x", operatorId), fmt.Sprintf("%t", errored), fmt.Sprintf("%t", hasNearDa)).Inc()
		},
		IncSignedOperatorSetUpdateMessageCb: func(operatorId [32]byte, errored bool) {
			signedOperatorSetUpdateMessagesTotal.WithLabelValues(fmt.Sprintf("%x", operatorId), fmt.Sprintf("%t", errored)).Inc()
		},
		ObserveLastMessageReceivedTimeCb: func(operatorId [32]byte, messageType string) {
			lastMessageReceivedTime.WithLabelValues(fmt.Sprintf("%x", operatorId), messageType).SetToCurrentTime()
		},
	}, nil
}

func MakeAggregatorMetrics(registry *prometheus.Registry) (AggregatorEventListener, error) {
	lastStateRootUpdateAggregated := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_state_root_update_aggregated",
			Help:      "Last state root update aggregated per rollup ID",
		},
		[]string{"rollup_id"},
	)
	if err := registry.Register(lastStateRootUpdateAggregated); err != nil {
		return nil, fmt.Errorf("error registering lastStateRootUpdateAggregated gauge: %w", err)
	}

	lastStateRootUpdateReceived := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_state_root_update_received",
			Help:      "Last state root update received per rollup ID",
		},
		[]string{"rollup_id"},
	)
	if err := registry.Register(lastStateRootUpdateReceived); err != nil {
		return nil, fmt.Errorf("error registering lastStateRootUpdateReceived gauge: %w", err)
	}

	lastOperatorSetUpdateAggregated := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_operator_set_update_aggregated",
			Help:      "Last operator set update aggregated",
		},
	)
	if err := registry.Register(lastOperatorSetUpdateAggregated); err != nil {
		return nil, fmt.Errorf("error registering lastOperatorSetUpdateAggregated gauge: %w", err)
	}

	lastOperatorSetUpdateReceived := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_operator_set_update_received",
			Help:      "Last operator set update received",
		},
	)
	if err := registry.Register(lastOperatorSetUpdateReceived); err != nil {
		return nil, fmt.Errorf("error registering lastBlockReceived gauge: %w", err)
	}

	expiredMessages := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "expired_messages_total",
			Help:      "Total number of expired messages",
		},
	)
	if err := registry.Register(expiredMessages); err != nil {
		return nil, fmt.Errorf("error registering expiredMessages counter: %w", err)
	}

	expiredTasks := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "expired_tasks_total",
			Help:      "Total number of expired tasks",
		},
	)
	if err := registry.Register(expiredTasks); err != nil {
		return nil, fmt.Errorf("error registering expiredTasks counter: %w", err)
	}

	erroredSubmissions := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "errored_submissions_total",
			Help:      "Total number of errored submissions",
		},
	)
	if err := registry.Register(erroredSubmissions); err != nil {
		return nil, fmt.Errorf("error registering erroredSubmissions counter: %w", err)
	}

	aggregatorInitializations := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "reinitializations_total",
			Help:      "Total number of aggregator reinitializations",
		},
	)
	if err := registry.Register(aggregatorInitializations); err != nil {
		return nil, fmt.Errorf("error registering aggregatorInitializations counter: %w", err)
	}

	lastCheckpointReferenceSent := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_checkpoint_reference_sent",
			Help:      "Last checkpoint reference sent",
		},
	)
	if err := registry.Register(lastCheckpointReferenceSent); err != nil {
		return nil, fmt.Errorf("error registering lastCheckpointReferenceSent gauge: %w", err)
	}

	lastCheckpointTaskReferenceReceived := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_checkpoint_task_reference_received",
			Help:      "Last checkpoint task reference received",
		},
	)
	if err := registry.Register(lastCheckpointTaskReferenceReceived); err != nil {
		return nil, fmt.Errorf("error registering lastCheckpointTaskReferenceReceived gauge: %w", err)
	}

	lastCheckpointTaskReferenceAggregated := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_checkpoint_task_reference_aggregated",
			Help:      "Last checkpoint task reference aggregated",
		},
	)
	if err := registry.Register(lastCheckpointTaskReferenceAggregated); err != nil {
		return nil, fmt.Errorf("error registering lastCheckpointTaskReferenceAggregated gauge: %w", err)
	}

	return &SelectiveAggregatorListener{
		ObserveLastStateRootUpdateAggregatedCb: func(rollupId uint32, blockNumber uint64) {
			lastStateRootUpdateAggregated.WithLabelValues(fmt.Sprintf("%x", rollupId)).Set(float64(blockNumber))
		},
		ObserveLastStateRootUpdateReceivedCb: func(rollupId uint32, blockNumber uint64) {
			lastStateRootUpdateReceived.WithLabelValues(fmt.Sprintf("%x", rollupId)).Set(float64(blockNumber))
		},
		ObserveLastOperatorSetUpdateAggregatedCb: func(operatorSetUpdateId uint64) {
			lastOperatorSetUpdateAggregated.Set(float64(operatorSetUpdateId))
		},
		ObserveLastOperatorSetUpdateReceivedCb: func(operatorSetUpdateId uint64) {
			lastOperatorSetUpdateReceived.Set(float64(operatorSetUpdateId))
		},
		IncExpiredMessagesCb: func() {
			expiredMessages.Inc()
		},
		IncExpiredTasksCb: func() {
			expiredTasks.Inc()
		},
		IncErroredSubmissionsCb: func() {
			erroredSubmissions.Inc()
		},
		IncAggregatorInitializationsCb: func() {
			aggregatorInitializations.Inc()
		},
		ObserveLastCheckpointReferenceSentCb: func(referenceId uint32) {
			lastCheckpointReferenceSent.Set(float64(referenceId))
		},
		ObserveLastCheckpointTaskReferenceReceivedCb: func(referenceId uint32) {
			lastCheckpointTaskReferenceReceived.Set(float64(referenceId))
		},
		ObserveLastCheckpointTaskReferenceAggregatedCb: func(referenceId uint32) {
			lastCheckpointTaskReferenceAggregated.Set(float64(referenceId))
		},
	}, nil
}
