package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandle(t *testing.T) {
	res, err := handler(context.Background(), events.APIGatewayProxyRequest{})

	if err != nil {
		t.Error("Unexpected error")
	}

	if res.Body != "v0.0.0-dev" {
		t.Errorf("Unexpected HTTP Response Body: %s", res.Body)
	}

	if res.StatusCode != 200 {
		t.Errorf("Unexpected HTTP Status Code: %d", res.StatusCode)
	}
}
