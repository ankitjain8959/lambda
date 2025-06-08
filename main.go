package main

import (
	"context"
	"lambda-using-go/src"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.Println("Initializing Lambda function...")
}

func main() {
	lambda.Start(handleRequest)
}

// This is a simple AWS Lambda function written in Go that returns a user object.
// The function is triggered by an event and returns a hardcoded user response.
func handleRequest(ctx context.Context, event map[string]interface{}) (src.User, error) {
	log.Println("Hello from AWS Lambda!")

	response := src.User{
		Id:          event["id"].(string),
		Name:        event["name"].(string),
		Email:       event["email"].(string),
		Age:         event["age"].(float64),
		Address:     event["address"].(string),
		Phone:       event["phone"].(string),
		DateOfBirth: event["dateOfBirth"].(string),
	}

	return response, nil
}
