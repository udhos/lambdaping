[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/lambdaping/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/lambdaping)](https://goreportcard.com/report/github.com/udhos/lambdaping)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/lambdaping.svg)](https://pkg.go.dev/github.com/udhos/lambdaping)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/lambdaping)](https://artifacthub.io/packages/search?repo=lambdaping)
[![Docker Pulls lambdaping](https://img.shields.io/docker/pulls/udhos/lambdaping)](https://hub.docker.com/r/udhos/lambdaping)

# lambdaping

# Build

```
git clone https://github.com/udhos/lambdaping

cd lambdaping

go install ./...
```

# Run

```
# optional
export DEBUG=true
export LAMBDA_ROLE_ARN=arn:aws:iam::100010001000:role/invoker
export BODY='{"hello":"world"}'

# mandatory
export LAMBDA_ARN=["arn:aws:lambda:us-east-1:100010001000:function:funcname"]

export OTELCONFIG_EXPORTER=jaeger
export OTEL_TRACES_EXPORTER=jaeger
export OTEL_PROPAGATORS=b3multi
export OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger-collector:14268

lambdaping
```

# Docker images

https://hub.docker.com/r/udhos/lambdaping


# Helm chart

https://udhos.github.io/lambdaping
