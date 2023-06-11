package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name LambdaEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func lambdaHandler() {
	lambda.Start(HandleRequest)
}