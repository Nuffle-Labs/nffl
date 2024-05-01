package relayer

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const RelayerNamespace = "sffl_relayer"

type EventListener interface {
	OnBlockReceived()
	OnDaSubmissionFailed()
	OnDaSubmitted(duration time.Duration)
}

type SelectiveListener struct {
	OnBlockReceivedCb      func()
	OnDaSubmissionFailedCb func()
	OnDaSubmittedCb        func(duration time.Duration)
}

func (l *SelectiveListener) OnBlockReceived() {
	if l.OnBlockReceivedCb != nil {
		l.OnBlockReceivedCb()
	}
}

func (l *SelectiveListener) OnDaSubmissionFailed() {
	if l.OnDaSubmissionFailedCb != nil {
		l.OnDaSubmissionFailedCb()
	}
}

func (l *SelectiveListener) OnDaSubmitted(duration time.Duration) {
	if l.OnDaSubmittedCb != nil {
		l.OnDaSubmittedCb(duration)
	}
}

func MakeRelayerMetrics(registry *prometheus.Registry) (EventListener, error) {
	numBlocksReceived := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: RelayerNamespace,
			Name:      "num_blocks_received",
			Help:      "The number of blocks received from rollup",
		})

	if err := registry.Register(numBlocksReceived); err != nil {
		return nil, fmt.Errorf("error registering numBlocksReceived counter: %w", err)
	}

	numDaSubmissionsFailed := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: RelayerNamespace,
		Name:      "num_da_submissions_failed",
		Help:      "The number of failed da submissions",
	})

	if err := registry.Register(numDaSubmissionsFailed); err != nil {
		return nil, fmt.Errorf("error registering numDaSubmissionsFailed counter: %w", err)
	}

	latencyBuckets := []float64{
		1,
		25,
		50,
		75,
		100,
		250,
		500,
		1000, // 1 sec
		2000,
		3000,
		4000,
		5000,
		10000,
		math.Inf(0),
	}

	submissionDuration := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: RelayerNamespace,
		Name:      "submission_duration_ms",
		Help:      "Duration of successful DA submissions",
		Buckets:   latencyBuckets,
	})

	if err := registry.Register(submissionDuration); err != nil {
		return nil, fmt.Errorf("error registering submissionDuration histogram: %w", err)
	}

	return &SelectiveListener{
		OnBlockReceivedCb: func() {
			numBlocksReceived.Inc()
		},
		OnDaSubmissionFailedCb: func() {
			numDaSubmissionsFailed.Inc()
		},
		OnDaSubmittedCb: func(duration time.Duration) {
			submissionDuration.Observe(float64(duration.Milliseconds()))
		},
	}, nil
}

func StartMetrics(metricsAddr string, reg prometheus.Gatherer) (<-chan error, func()) {
	errC := make(chan error, 1)
	server := &http.Server{Addr: metricsAddr, Handler: nil}

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	shutdown := func() {
		if err := server.Shutdown(context.Background()); err != nil {
			// Handle the error according to your application's needs, e.g., log it
			log.Printf("Error shutting down metrics server: %v", err)
		}
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, shutdown
}
