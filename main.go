package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, event map[string]interface{}) {
	log.Println("Hello from AWS Lambda!")
}
