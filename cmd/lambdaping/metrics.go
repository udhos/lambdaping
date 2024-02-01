package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	latencySpringClient *prometheus.HistogramVec
}

var (
	dimensionsSpring = []string{"method", "status", "uri"}
)

const (
	latencySpringNameClient = "lambda_client_requests_seconds"
)

func newMetrics(namespace string, latencyBucketsClient []float64) *metrics {
	const me = "newMetrics"

	//
	// latency client
	//

	latencySpringClient := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      latencySpringNameClient,
			Help:      "Spring-like client request duration in seconds.",
			Buckets:   latencyBucketsClient,
		},
		dimensionsSpring,
	)

	if err := prometheus.Register(latencySpringClient); err != nil {
		log.Fatalf("%s: client latency was not registered: %s", me, err)
	}

	//
	// all metrics
	//

	m := &metrics{
		latencySpringClient: latencySpringClient,
	}

	return m
}

func (m *metrics) recordLatencyClient(method, status, path string, latency time.Duration) {
	m.latencySpringClient.WithLabelValues(method, status, path).Observe(float64(latency) / float64(time.Second))
}
