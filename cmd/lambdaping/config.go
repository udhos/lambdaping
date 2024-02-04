package main

import (
	"time"

	"github.com/udhos/boilerplate/envconfig"
)

type config struct {
	interval                           time.Duration
	body                               string
	healthAddr                         string
	healthPath                         string
	lambdaArn                          string
	lambdaRoleArn                      string
	debug                              bool
	metricsAddr                        string
	metricsPath                        string
	metricsNamespace                   string
	metricsNameLatencySpringNameClient string
	metricsLabelMethod                 string
	metricsLabelStatus                 string
	metricsLabelURI                    string
	metricsLatencyBucketsClient        []float64
}

func getConfig(roleSessionName string) config {

	env := envconfig.NewSimple(roleSessionName)

	return config{
		interval:                           env.Duration("INTERVAL", 10*time.Second),
		body:                               env.String("BODY", `{"hello":"world"}`),
		healthAddr:                         env.String("HEALTH_ADDR", ":8888"),
		healthPath:                         env.String("HEALTH_PATH", "/health"),
		lambdaArn:                          env.String("LAMBDA_ARN", ""),
		lambdaRoleArn:                      env.String("LAMBDA_ROLE_ARN", ""),
		debug:                              env.Bool("DEBUG", false),
		metricsAddr:                        env.String("METRICS_ADDR", ":3000"),
		metricsPath:                        env.String("METRICS_PATH", "/metrics"),
		metricsNamespace:                   env.String("METRICS_NAMESPACE", "lambda"),
		metricsNameLatencySpringNameClient: env.String("METRICS_LATENCY_SPRING_NAME_CLIENT", "client_requests_seconds"),
		metricsLabelMethod:                 env.String("METRICS_LABEL_METHOD", "method"),
		metricsLabelStatus:                 env.String("METRICS_LABEL_STATUS", "status"),
		metricsLabelURI:                    env.String("METRICS_LABEL_URI", "uri"),
		metricsLatencyBucketsClient:        env.Float64Slice("METRICS_BUCKETS_LATENCY_CLIENT", []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, .5, 1, 2.5, 5, 10, 25, 50}),
	}
}
