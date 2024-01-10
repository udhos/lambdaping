package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/udhos/boilerplate/awsconfig"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/yaml.v3"
)

func pinger(app *application) {
	const me = "pinger"

	var lambdaARNsList []string

	if errYaml := yaml.Unmarshal([]byte(app.conf.lambdaArn), &lambdaARNsList); errYaml != nil {
		log.Fatalf("%s: parse yaml lambda ARNs: %v", me, errYaml)
	}

	size := len(lambdaARNsList)

	if size < 1 {
		log.Fatalf("%s: bad number of lambda ARNs: %d", me, size)
	}

	countOk := make([]int, size)
	countErrors := make([]int, size)

	for {
		for i, arn := range lambdaARNsList {
			if errInvoke := invoke(arn, app.lambdaClient, app.tracer, app.conf.debug); errInvoke == nil {
				countOk[i]++
			} else {
				log.Printf("%s: invoke error: %v", me, errInvoke)
				countErrors[i]++
			}
			if app.conf.debug {
				log.Printf("%s: %s: success=%d error=%d",
					me, arn, countOk[i], countErrors[i])
			}
		}
		if app.conf.debug {
			log.Printf("%s: sleeping for %v",
				me, app.conf.interval)
		}
		time.Sleep(app.conf.interval)
	}
}

func traceError(span trace.Span, e error) error {
	span.SetStatus(codes.Error, e.Error())
	return e
}

func invoke(lambdaArn string, clientLambda *lambda.Client, tracer trace.Tracer, debug bool) error {
	const me = "invoke"

	_, span := tracer.Start(context.TODO(), me)
	defer span.End()

	requestBytes := []byte(`{"lambdaping":"hello"}`)

	input := &lambda.InvokeInput{
		FunctionName: aws.String(lambdaArn),
		Payload:      requestBytes,
	}

	resp, errInvoke := clientLambda.Invoke(context.TODO(), input)
	if errInvoke != nil {
		return traceError(span, errInvoke)
	}

	if resp.StatusCode != http.StatusOK {
		errStatus := fmt.Errorf("%s: Invoke ARN=%s bad status=%d payload: %s",
			me, lambdaArn, resp.StatusCode, resp.Payload)
		return traceError(span, errStatus)
	}

	var funcError string
	if resp.FunctionError != nil {
		funcError = *resp.FunctionError
	}
	if funcError != "" {
		errFunc := fmt.Errorf("%s: Invoke ARN=%s function_error='%s' payload: %s",
			me, lambdaArn, funcError, resp.Payload)
		return traceError(span, errFunc)
	}

	if debug {
		log.Printf("%s: Invoke ARN=%s status_code=%d function_error='%s' payload: %s",
			me, lambdaArn, resp.StatusCode, funcError, resp.Payload)
	}

	return nil
}

func newLambdaClient(lambdaArn, roleArn, roleSessionName, roleExternalID string) (*lambda.Client, error) {
	region, errRegion := getARNRegion(lambdaArn)
	if errRegion != nil {
		return nil, errRegion
	}

	cfg, _, errConfig := awsConfig(region, roleArn, roleExternalID, roleSessionName)
	if errConfig != nil {
		return nil, errConfig
	}

	client := lambda.NewFromConfig(cfg)

	return client, nil
}

// arn:aws:sns:us-east-1:123456789012:gateboard
// arn:aws:lambda:us-east-1:123456789012:function:forward_to_sqs
func getARNRegion(arn string) (string, error) {
	const me = "getARNRegion"
	fields := strings.SplitN(arn, ":", 5)
	if len(fields) < 5 {
		return "", fmt.Errorf("%s: bad ARN=[%s]", me, arn)
	}
	region := fields[3]
	log.Printf("%s=[%s]", me, region)
	return region, nil
}

func awsConfig(region, roleArn, roleExternalID, roleSessionName string) (aws.Config, string, error) {

	awsConfOptions := awsconfig.Options{
		Region:          region,
		RoleArn:         roleArn,
		RoleSessionName: roleSessionName,
		RoleExternalID:  roleExternalID,
	}
	awsConf, errAwsConf := awsconfig.AwsConfig(awsConfOptions)
	if errAwsConf != nil {
		return awsConf.AwsConfig, awsConf.StsAccountID, errAwsConf
	}

	return awsConf.AwsConfig, awsConf.StsAccountID, errAwsConf
}
