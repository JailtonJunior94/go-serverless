package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {
	val, ok := req.Headers["api-key"]
	if !ok {
		return Response{StatusCode: http.StatusUnprocessableEntity}, nil
	}

	if val != "45aabf67-8337-4eb7-8d2c-2cf6b554fbf4" {
		return Response{StatusCode: http.StatusUnauthorized}, nil
	}

	return Response{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
