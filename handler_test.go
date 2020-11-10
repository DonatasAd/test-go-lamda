package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

func TestHandlerBasicOK(t *testing.T) {
	SessionClient = MockAWSSessionClient

	response, _ := Handler(events.APIGatewayProxyRequest{
		Body:           `{"loanID":12345,"loanAmount":99999.99}`,
		RequestContext: events.APIGatewayProxyRequestContext{RequestID: uuid.New().String()},
	})
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestHandlerBadRequest(t *testing.T) {
	SessionClient = MockAWSSessionClient

	response, _ := Handler(events.APIGatewayProxyRequest{
		Body:           `{"loanID":"12345","loanAmount":99999.99}`,
		RequestContext: events.APIGatewayProxyRequestContext{RequestID: uuid.New().String()},
	})
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}
