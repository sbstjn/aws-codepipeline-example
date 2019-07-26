package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var version string

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if name, ok := request.QueryStringParameters["name"]; ok {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: fmt.Sprintf("Bonjour, %s!", name) }, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Query parameter name required."}, errors.New("Query parameter name required.")
}
