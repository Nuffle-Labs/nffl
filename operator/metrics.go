package operator

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	OperatorNamespace = "sffl_operator"
	OperatorSubsytem  = "operator"
)

type OperatorEventListener interface {
	OnTasksReceived()
	IncInitializationCount()
	ObserveLastInitializedTime()
}

type SelectiveOperatorListener struct {
	OnTasksReceivedCb            func()
	IncInitializationCountCb     func()
	ObserveLastInitializedTimeCb func()
}

func (l *SelectiveOperatorListener) OnTasksReceived() {
	if l.OnTasksReceivedCb != nil {
		l.OnTasksReceivedCb()
	}
}

func (l *SelectiveOperatorListener) IncInitializationCount() {
	if l.IncInitializationCountCb != nil {
		l.IncInitializationCountCb()
	}
}

func (l *SelectiveOperatorListener) ObserveLastInitializedTime() {
	if l.ObserveLastInitializedTimeCb != nil {
		l.ObserveLastInitializedTimeCb()
	}
}

type RpcClientEventListener interface {
	OnMessagesReceived()
	ObserveResendQueueSize(size int)
	ObserveLastCheckpointIdResponded(checkpointId uint32)
	ObserveLastOperatorSetUpdateIdResponded(operatorSetUpdateId uint64)
	IncStateRootUpdateSubmissions(rollupId uint32, resend bool)
	IncOperatorSetUpdateUpdateSubmissions(resend bool)
	IncCheckpointTaskResponseSubmissions(resend bool)
	IncErroredStateRootUpdateSubmissions(rollupId uint32, resend bool)
	IncErroredOperatorSetUpdateSubmissions(resend bool)
	IncErroredCheckpointSubmissions(resend bool)
}

type SelectiveRpcClientListener struct {
	OnMessagesReceivedCb                      func()
	ObserveResendQueueSizeCb                  func(size int)
	ObserveLastCheckpointIdRespondedCb        func(checkpointId uint32)
	ObserveLastOperatorSetUpdateIdRespondedCb func(operatorSetUpdateId uint64)
	IncStateRootUpdateSubmissionsCb           func(rollupId uint32, resend bool)
	IncOperatorSetUpdateUpdateSubmissionsCb   func(resend bool)
	IncCheckpointTaskResponseSubmissionsCb    func(resend bool)
	IncErroredStateRootUpdateSubmissionsCb    func(rollupId uint32, resend bool)
	IncErroredOperatorSetUpdateSubmissionsCb  func(resend bool)
	IncErroredCheckpointSubmissionsCb         func(resend bool)
}

func (l *SelectiveRpcClientListener) OnMessagesReceived() {
	if l.OnMessagesReceivedCb != nil {
		l.OnMessagesReceivedCb()
	}
}

func (l *SelectiveRpcClientListener) ObserveResendQueueSize(size int) {
	if l.ObserveResendQueueSizeCb != nil {
		l.ObserveResendQueueSizeCb(size)
	}
}

func (l *SelectiveRpcClientListener) ObserveLastCheckpointIdResponded(checkpointId uint32) {
	if l.ObserveLastCheckpointIdRespondedCb != nil {
		l.ObserveLastCheckpointIdRespondedCb(checkpointId)
	}
}

func (l *SelectiveRpcClientListener) ObserveLastOperatorSetUpdateIdResponded(operatorSetUpdateId uint64) {
	if l.ObserveLastOperatorSetUpdateIdRespondedCb != nil {
		l.ObserveLastOperatorSetUpdateIdRespondedCb(operatorSetUpdateId)
	}
}

func (l *SelectiveRpcClientListener) IncStateRootUpdateSubmissions(rollupId uint32, resend bool) {
	if l.IncStateRootUpdateSubmissionsCb != nil {
		l.IncStateRootUpdateSubmissionsCb(rollupId, resend)
	}
}

func (l *SelectiveRpcClientListener) IncOperatorSetUpdateUpdateSubmissions(resend bool) {
	if l.IncOperatorSetUpdateUpdateSubmissionsCb != nil {
		l.IncOperatorSetUpdateUpdateSubmissionsCb(resend)
	}
}

func (l *SelectiveRpcClientListener) IncCheckpointTaskResponseSubmissions(resend bool) {
	if l.IncCheckpointTaskResponseSubmissionsCb != nil {
		l.IncCheckpointTaskResponseSubmissionsCb(resend)
	}
}

func (l *SelectiveRpcClientListener) IncErroredStateRootUpdateSubmissions(rollupId uint32, resend bool) {
	if l.IncErroredStateRootUpdateSubmissionsCb != nil {
		l.IncErroredStateRootUpdateSubmissionsCb(rollupId, resend)
	}
}

func (l *SelectiveRpcClientListener) IncErroredCheckpointSubmissions(resend bool) {
	if l.IncErroredCheckpointSubmissionsCb != nil {
		l.IncErroredCheckpointSubmissionsCb(resend)
	}
}

func (l *SelectiveRpcClientListener) IncErroredOperatorSetUpdateSubmissions(resend bool) {
	if l.IncErroredOperatorSetUpdateSubmissionsCb != nil {
		l.IncErroredOperatorSetUpdateSubmissionsCb(resend)
	}
}

func MakeOperatorMetrics(registry *prometheus.Registry) (OperatorEventListener, error) {
	numTasksReceived := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "num_tasks_received",
			Help:      "The number of tasks received by reading from the avs service manager contract",
		})

	initializationCount := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "initialization_count",
			Help:      "Initialization count",
		},
	)

	lastInitializedTime := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: OperatorNamespace,
			Name:      "last_initialized_time",
			Help:      "Last initialized time",
		},
	)

	if err := registry.Register(numTasksReceived); err != nil {
		return nil, fmt.Errorf("error registering numTasksReceived counter: %w", err)
	}
	if err := registry.Register(lastInitializedTime); err != nil {
		return nil, fmt.Errorf("error registering lastInitializedTime gauge: %w", err)
	}
	if err := registry.Register(initializationCount); err != nil {
		return nil, fmt.Errorf("error registering initializationCount counter: %w", err)
	}

	return &SelectiveOperatorListener{
		OnTasksReceivedCb: func() {
			numTasksReceived.Inc()
		},
		IncInitializationCountCb: func() {
			initializationCount.Inc()
		},
		ObserveLastInitializedTimeCb: func() {
			lastInitializedTime.SetToCurrentTime()
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

	resendQueueSize := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: OperatorNamespace,
			Name:      "resend_queue_size",
			Help:      "Resend queue size",
		})

	lastCheckpointIdResponded := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: OperatorNamespace,
			Name:      "last_checkpoint_id_responded",
			Help:      "Last checkpoint ID responded",
		},
	)

	lastOperatorSetUpdateIdResponded := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: OperatorNamespace,
			Name:      "last_operator_set_update_id_responded",
			Help:      "Last operator set update ID responded",
		},
	)

	stateRootUpdateSubmissions := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "state_root_update_submissions",
			Help:      "State root updates submitted over time per rollup ID",
		},
		[]string{"rollup_id", "resend"},
	)

	operatorSetUpdateSubmissions := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "operator_set_update_submissions",
			Help:      "Operator set updates submitted over time",
		},
		[]string{"resend"},
	)

	checkpointTaskResponseSubmissions := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "checkpoint_task_response_submissions",
			Help:      "Checkpoint task response submissions over time",
		},
		[]string{"resend"},
	)

	erroredStateRootUpdateSubmissions := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "errored_state_root_update_submissions",
			Help:      "Errored state root update submissions over time per rollup ID",
		},
		[]string{"rollup_id", "resend"},
	)

	erroredCheckpointSubmissions := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "errored_checkpoint_submissions",
			Help:      "Errored checkpoint submissions over time",
		},
		[]string{"resend"},
	)

	erroredOperatorSetUpdateSubmissions := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Name:      "errored_operator_set_update_submissions",
			Help:      "Errored operator set update submissions over time",
		},
		[]string{"resend"},
	)

	if err := registry.Register(numMessagesReceived); err != nil {
		return nil, fmt.Errorf("error registering numMessagesReceived counter: %w", err)
	}
	if err := registry.Register(resendQueueSize); err != nil {
		return nil, fmt.Errorf("error registering resendQueueSize gauge: %w", err)
	}
	if err := registry.Register(lastCheckpointIdResponded); err != nil {
		return nil, fmt.Errorf("error registering lastCheckpointIdResponded gauge: %w", err)
	}
	if err := registry.Register(lastOperatorSetUpdateIdResponded); err != nil {
		return nil, fmt.Errorf("error registering lastOperatorSetUpdateIdResponded gauge: %w", err)
	}
	if err := registry.Register(stateRootUpdateSubmissions); err != nil {
		return nil, fmt.Errorf("error registering stateRootUpdatesSent counter: %w", err)
	}
	if err := registry.Register(operatorSetUpdateSubmissions); err != nil {
		return nil, fmt.Errorf("error registering operatorSetUpdateSubmissions counter: %w", err)
	}
	if err := registry.Register(checkpointTaskResponseSubmissions); err != nil {
		return nil, fmt.Errorf("error registering checkpointTaskResponseSubmissions counter: %w", err)
	}
	if err := registry.Register(erroredCheckpointSubmissions); err != nil {
		return nil, fmt.Errorf("error registering erroredCheckpointSubmissions counter: %w", err)
	}
	if err := registry.Register(erroredOperatorSetUpdateSubmissions); err != nil {
		return nil, fmt.Errorf("error registering erroredOperatorSetUpdateSubmissions counter: %w", err)
	}

	return &SelectiveRpcClientListener{
		OnMessagesReceivedCb: func() {
			numMessagesReceived.Inc()
		},
		ObserveResendQueueSizeCb: func(size int) {
			resendQueueSize.Set(float64(size))
		},
		ObserveLastCheckpointIdRespondedCb: func(checkpointId uint32) {
			lastCheckpointIdResponded.Set(float64(checkpointId))
		},
		ObserveLastOperatorSetUpdateIdRespondedCb: func(operatorSetUpdateId uint64) {
			lastOperatorSetUpdateIdResponded.Set(float64(operatorSetUpdateId))
		},
		IncStateRootUpdateSubmissionsCb: func(rollupId uint32, resend bool) {
			stateRootUpdateSubmissions.WithLabelValues(fmt.Sprintf("%d", rollupId), fmt.Sprintf("%t", resend)).Inc()
		},
		IncOperatorSetUpdateUpdateSubmissionsCb: func(resend bool) {
			operatorSetUpdateSubmissions.WithLabelValues(fmt.Sprintf("%t", resend)).Inc()
		},
		IncCheckpointTaskResponseSubmissionsCb: func(resend bool) {
			checkpointTaskResponseSubmissions.WithLabelValues(fmt.Sprintf("%t", resend)).Inc()
		},
		IncErroredStateRootUpdateSubmissionsCb: func(rollupId uint32, resend bool) {
			erroredStateRootUpdateSubmissions.WithLabelValues(fmt.Sprintf("%d", rollupId), fmt.Sprintf("%t", resend)).Inc()
		},
		IncErroredCheckpointSubmissionsCb: func(resend bool) {
			erroredCheckpointSubmissions.WithLabelValues(fmt.Sprintf("%t", resend)).Inc()
		},
		IncErroredOperatorSetUpdateSubmissionsCb: func(resend bool) {
			erroredOperatorSetUpdateSubmissions.WithLabelValues(fmt.Sprintf("%t", resend)).Inc()
		},
	}, nil
}
