package relayer

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
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

func startMetrics(metricsAddr string, reg prometheus.Gatherer) (<-chan error, func(ctx context.Context) error) {
	errC := make(chan error, 1)
	server := &http.Server{Addr: metricsAddr, Handler: nil}

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	shutdown := func(ctx context.Context) error {
		if err := server.Shutdown(ctx); err != nil {
			return err
		}

		return nil
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, shutdown
}

func StartMetricsServer(ctx context.Context, metricsAddr string, reg prometheus.Gatherer, logger sdklogging.Logger) {
	const (
		RetryCount    = 3
		RetryInterval = time.Second
	)

	errC, shutdownMetrics := startMetrics(metricsAddr, reg)
	go func() {
		retryCount := 0
		for {
			select {
			case err := <-errC:
				if err == nil {
					continue
				}

				if retryCount >= RetryCount {
					logger.Error("Failed to restart metrics server after multiple attempts", "err", err)
					return
				} else {
					logger.Error("Metrics server error", "err", err)
				}

				err = shutdownMetrics(ctx)
				if err != nil {
					logger.Error("Error while shutting down", "err", err)
					return
				}

				// Sleep before restart
				time.Sleep(RetryInterval)

				logger.Info("Attempting to restart metrics server")
				errC, shutdownMetrics = startMetrics(metricsAddr, reg)
				retryCount++
			case <-ctx.Done():
				return
			}
		}
	}()
}
