package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleWithoutQueryParameter(t *testing.T) {
	res, err := handler(context.Background(), events.APIGatewayProxyRequest{})

	if err == nil {
		t.Error("Expected error")
	}

	if res.StatusCode != 400 {
		t.Errorf("Unexpected HTTP Status Code: %d", res.StatusCode)
	}
}

func TestHandleWithQueryParameter(t *testing.T) {
	res, err := handler(context.Background(), events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"name": "Franc",
		},
	})

	if err != nil {
		t.Error("Unexpected error")
	}

	if res.Body != "Bonjour, Franc!" {
		t.Errorf("Unexpected HTTP Response Body: %s", res.Body)
	}
}