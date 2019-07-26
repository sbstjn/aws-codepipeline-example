package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var version string

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if version == "" {
		version = "v0.0.0-dev"
	}
	
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: version}, nil
}
