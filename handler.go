package main

import (
	"encoding/json"
	"fmt"

	"github.com/maxexllc/logging"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// initialize dependencies
	sessClient := SessionClient
	sess := sessClient.Must(sessClient.NewSession())

	// unmarshal request from apigatewayrequest body
	var exampleRequest ExampleRequest
	err := json.Unmarshal([]byte(request.Body), &exampleRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, ErrCouldNotUnmarshalRequest
	}

	// initialize logger
	log, err := logging.NewLambdaAPIRequestLoggerEntry(
		logrus.InfoLevel.String(),
		"data-receiver/mls",
		&request.RequestContext,
		request,
		sess,
	)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, ErrFailedToInitializeLogger
	}

	log.Printf("Processing Lambda request.")

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
