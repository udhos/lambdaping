package main

import (
	"time"

	"github.com/udhos/boilerplate/envconfig"
)

type config struct {
	interval      time.Duration
	healthAddr    string
	healthPath    string
	lambdaArn     string
	lambdaRoleArn string
	debug         bool
}

func getConfig(roleSessionName string) config {

	env := envconfig.NewSimple(roleSessionName)

	return config{
		interval:      env.Duration("INTERVAL", 10*time.Second),
		healthAddr:    env.String("HEALTH_ADDR", ":8888"),
		healthPath:    env.String("HEALTH_PATH", "/health"),
		lambdaArn:     env.String("LAMBDA_ARN", ""),
		lambdaRoleArn: env.String("LAMBDA_ROLE_ARN", ""),
		debug:         env.Bool("DEBUG", false),
	}
}
