package attestor

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type EventListener interface {
	OnMissedMQBlock(rollupId uint32)
	OnBlockMismatch(rollupId uint32)
	OnBlockReceived(rollupId uint32)
	ObserveLastBlockReceived(rollupId uint32, blockNumber uint64)
	ObserveLastBlockReceivedTimestamp(rollupId uint32, timestamp uint64)
	ObserveInitializationInitialBlockNumber(rollupId uint32, blockNumber uint64)
}

const OperatorNamespace = "sffl_operator"
const AttestorSubsystem = "attestor"

type SelectiveEventListener struct {
	OnMissedMQBlockCb                         func(rollupId uint32)
	OnBlockMismatchCb                         func(rollupId uint32)
	OnBlockReceivedCb                         func(rollupId uint32)
	ObserveLastBlockReceivedCb                func(rollupId uint32, blockNumber uint64)
	ObserveLastBlockReceivedTimestampCb       func(rollupId uint32, timestamp uint64)
	ObserveInitializationInitialBlockNumberCb func(rollupId uint32, blockNumber uint64)
}

func (l *SelectiveEventListener) OnMissedMQBlock(rollupId uint32) {
	if l.OnMissedMQBlockCb != nil {
		l.OnMissedMQBlockCb(rollupId)
	}
}

func (l *SelectiveEventListener) OnBlockMismatch(rollupId uint32) {
	if l.OnBlockMismatchCb != nil {
		l.OnBlockMismatchCb(rollupId)
	}
}

func (l *SelectiveEventListener) OnBlockReceived(rollupId uint32) {
	if l.OnBlockReceivedCb != nil {
		l.OnBlockReceivedCb(rollupId)
	}
}

func (l *SelectiveEventListener) ObserveLastBlockReceived(rollupId uint32, blockNumber uint64) {
	if l.ObserveLastBlockReceivedCb != nil {
		l.ObserveLastBlockReceivedCb(rollupId, blockNumber)
	}
}

func (l *SelectiveEventListener) ObserveLastBlockReceivedTimestamp(rollupId uint32, timestamp uint64) {
	if l.ObserveLastBlockReceivedTimestampCb != nil {
		l.ObserveLastBlockReceivedTimestampCb(rollupId, timestamp)
	}
}

func (l *SelectiveEventListener) ObserveInitializationInitialBlockNumber(rollupId uint32, blockNumber uint64) {
	if l.ObserveInitializationInitialBlockNumberCb != nil {
		l.ObserveInitializationInitialBlockNumberCb(rollupId, blockNumber)
	}
}

func MakeAttestorMetrics(registry *prometheus.Registry) (EventListener, error) {
	numMissedMqBlocks := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: AttestorSubsystem,
			Name:      "num_of_missed_mq_blocks",
			Help:      "The number of late blocks from MQ",
		},
		[]string{"rollup_id"},
	)

	if err := registry.Register(numMissedMqBlocks); err != nil {
		return nil, fmt.Errorf("error registering numMissedMqBlocks counter: %w", err)
	}

	numBlocksMismatched := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: AttestorSubsystem,
			Name:      "num_of_mismatched_blocks",
			Help:      "The number of blocks from MQ mismatched with RPC ones.",
		},
		[]string{"rollup_id"},
	)

	if err := registry.Register(numBlocksMismatched); err != nil {
		return nil, fmt.Errorf("error registering numBlocksMismatched counter: %w", err)
	}

	numBlocksReceived := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: OperatorNamespace,
			Subsystem: AttestorSubsystem,
			Name:      "num_of_received_blocks",
			Help:      "The number of blocks received from RPC.",
		},
		[]string{"rollup_id"},
	)

	if err := registry.Register(numBlocksReceived); err != nil {
		return nil, fmt.Errorf("error registering numBlocksReceived counter: %w", err)
	}

	lastBlockReceived := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: OperatorNamespace,
			Name:      "last_block_received",
			Help:      "Last block received per rollup ID",
		},
		[]string{"rollup_id"},
	)

	if err := registry.Register(lastBlockReceived); err != nil {
		return nil, fmt.Errorf("error registering lastBlockReceived gauge: %w", err)
	}

	lastBlockReceivedTimestamp := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: OperatorNamespace,
			Name:      "last_block_received_timestamp",
			Help:      "Timestamp of last block received per rollup ID",
		},
		[]string{"rollup_id"},
	)

	if err := registry.Register(lastBlockReceivedTimestamp); err != nil {
		return nil, fmt.Errorf("error registering lastBlockReceivedTimestamp gauge: %w", err)
	}

	initializationInitialBlockNumber := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: OperatorNamespace,
			Name:      "initialization_initial_block_number",
			Help:      "Initialization initial block number per rollup ID",
		},
		[]string{"rollup_id"},
	)

	if err := registry.Register(initializationInitialBlockNumber); err != nil {
		return nil, fmt.Errorf("error registering initializationInitialBlockNumber gauge: %w", err)
	}

	return &SelectiveEventListener{
		OnMissedMQBlockCb: func(rollupId uint32) {
			numMissedMqBlocks.WithLabelValues(fmt.Sprint(rollupId)).Inc()
		},
		OnBlockMismatchCb: func(rollupId uint32) {
			numBlocksMismatched.WithLabelValues(fmt.Sprint(rollupId)).Inc()
		},
		ObserveLastBlockReceivedCb: func(rollupId uint32, blockNumber uint64) {
			lastBlockReceived.WithLabelValues(fmt.Sprint(rollupId)).Set(float64(blockNumber))
		},
		ObserveLastBlockReceivedTimestampCb: func(rollupId uint32, timestamp uint64) {
			lastBlockReceivedTimestamp.WithLabelValues(fmt.Sprint(rollupId)).Set(float64(timestamp))
		},
		ObserveInitializationInitialBlockNumberCb: func(rollupId uint32, blockNumber uint64) {
			initializationInitialBlockNumber.WithLabelValues(fmt.Sprint(rollupId)).Set(float64(blockNumber))
		},
	}, nil
}
