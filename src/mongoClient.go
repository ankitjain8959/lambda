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

// DeleteUser deletes a user document from the MongoDB collection based on the user ID.
func DeleteUser(client *mongo.Client, userId string) (interface{}, error) {
	database := client.Database("tmf-productorder").Collection("productorder")
	filter := map[string]interface{}{"_id": userId}
	deletedResponse, err := database.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting user from MongoDB")
		return 0, err
	}
	return deletedResponse.DeletedCount, nil
}

// GetUser retrieves a user document from the MongoDB collection based on the user ID.
func GetUser(client *mongo.Client, userId string) (User, error) {
	database := client.Database("tmf-productorder").Collection("productorder")
	filter := map[string]interface{}{"_id": userId}
	var user User
	err := database.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println("Error retrieving user from MongoDB")
		return User{}, err
	}
	return user, nil
}

// UpdateUser updates an existing user document in the MongoDB collection based on the user ID.
func UpdateUser(client *mongo.Client, user User) (interface{}, error) {
	database := client.Database("tmf-productorder").Collection("productorder")
	filter := map[string]interface{}{"_id": user.Id}
	update := map[string]interface{}{
		"$set": user,
	}
	result, err := database.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating user in MongoDB")
		return 0, err
	}
	return result.ModifiedCount, nil
}
