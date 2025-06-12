package src

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectToMongoDB establishes a connection to the MongoDB server and returns a client instance.
// It uses the MongoDB Go Driver to connect to a MongoDB instance running in a Docker container.
func ConnectToMongoDB() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := "mongodb://my-mongo:27017" // Run docker container on a network `docker run --name my-mongo --network lambda-mongo-net -d -p 27017:27017 mongo`
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

// InsertUser inserts a new user document into the MongoDB collection.
func InsertUser(client *mongo.Client, user User) (interface{}, error) {

	database := client.Database("tmf-productorder").Collection("productorder")
	insertedResponse, err := database.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error inserting user into MongoDB")
		return 0, err
	}
	return insertedResponse.InsertedID, nil
}
