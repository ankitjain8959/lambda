package main

import (
	"context"
	"lambda-using-go/src"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func init() {
	log.Println("Initializing Lambda function...")
}

func main() {
	lambda.Start(handleRequest)
}

// This function handles the incoming Lambda event, connects to MongoDB, and inserts a user document.
// It returns the inserted user document or an error if the operation fails.
func handleRequest(ctx context.Context, event map[string]interface{}) (interface{}, error) {
	log.Println("Hello from AWS Lambda!")
	client = src.ConnectToMongoDB()

	response := src.User{
		Id:          event["id"].(string),
		Name:        event["name"].(string),
		Email:       event["email"].(string),
		Age:         event["age"].(float64),
		Address:     event["address"].(string),
		Phone:       event["phone"].(string),
		DateOfBirth: event["dateOfBirth"].(string),
	}

	insertedResponse, err := src.InsertUser(client, response)
	if err != nil {
		log.Println("Error inserting user into MongoDB:", err)
		return nil, err
	}

	log.Println("User inserted successfully:", insertedResponse)

	return insertedResponse, nil
}
