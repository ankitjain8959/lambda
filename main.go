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

	receivedEvent := src.User{
		Id:          event["id"].(string),
		Name:        event["name"].(string),
		Email:       event["email"].(string),
		Age:         event["age"].(float64),
		Address:     event["address"].(string),
		Phone:       event["phone"].(string),
		DateOfBirth: event["dateOfBirth"].(string),
		AtType:      event["@type"].(string),
	}

	log.Println("Received event:", receivedEvent)
	insertUser(receivedEvent)
	return getUser(receivedEvent.Id)
	// return updateUser(receivedEvent)
	// return deleteUser(receivedEvent.Id)
}

func insertUser(eventJson src.User) (interface{}, error) {
	insertedResponse, err := src.InsertUser(client, eventJson)
	if err != nil {
		log.Println("Error inserting user into MongoDB:", err)
		return nil, err
	}

	log.Println("User inserted successfully:", insertedResponse)
	return insertedResponse, nil
}

func deleteUser(userId string) (interface{}, error) {
	deletedResponse, err := src.DeleteUser(client, userId)
	if err != nil {
		log.Println("Error deleting user from MongoDB:", err)
		return nil, err
	}

	log.Println("User deleted successfully:", deletedResponse)
	return deletedResponse, nil
}

func getUser(userId string) (src.User, error) {
	user, err := src.GetUser(client, userId)
	if err != nil {
		log.Println("Error retrieving user from MongoDB:", err)
		return src.User{}, err
	}
	log.Println("User retrieved successfully:", user)
	return user, nil
}

func updateUser(user src.User) (interface{}, error) {
	count, err := src.UpdateUser(client, user)
	if err != nil {
		log.Println("Error updating user in MongoDB:", err)
		return nil, err
	}
	log.Println("Number of users updated:", count)
	return count, nil
}
