# Lambda using GO
This project demonstrates how to build an AWS Lambda function in Go that performs CRUD operations on a MongoDB database. The Lambda function is designed to run locally (using AWS SAM CLI or similar tools) and connect to a MongoDB instance running in a Docker container.

## Features
- **Insert User:** Adds a new user document to MongoDB.
- **Get User:** Retrieves a user document by ID.
- **Update User:** Updates an existing user document.
- **Delete User:** Deletes a user document by ID.

## Project Structure
- `main.go`: Entry point for the Lambda function. Handles incoming events and routes them to the appropriate MongoDB operation.
- `src/mongoClient.go`: Contains MongoDB connection logic and CRUD helper functions.
- `src/userType.go`: Defines the `User` struct.
- `go.mod`: Go module file that specifies dependencies.
- `go.sum`: Go module checksum file.

## How It Works
1. **MongoDB Connection:**  
   The Lambda function connects to a MongoDB instance using the Go MongoDB driver. The connection URI is set to `mongodb://my-mongo:27017`, which assumes MongoDB is running in a Docker container named `my-mongo` on a user-defined Docker network.

2. **Event Handling:**  
   The Lambda handler (`handleRequest`) receives an event (as a map), parses it into a `User` struct, and performs the requested operation (insert, get, update, or delete).

3. **CRUD Operations:**  
   - **Insert:** Adds a new user document.
   - **Get:** Retrieves a user by ID.
   - **Update:** Updates user details.
   - **Delete:** Removes a user by ID.

## Running Locally
### 1. Start MongoDB in Docker

```sh
docker network create lambda-mongo-net
docker run --name my-mongo --network lambda-mongo-net -d -p 27017:27017 mongo
```

### 2. Run Lambda Locally

If using AWS SAM CLI:

```sh
sam local invoke --docker-network lambda-mongo-net ...
```

### 3. MongoDB URI

The connection string in `src/mongoClient.go` should be:

```go
uri := "mongodb://my-mongo:27017"
```

## Environment Assumptions

- Both Lambda and MongoDB containers are on the same Docker network (`lambda-mongo-net`).
- The MongoDB database and collection (`tmf-productorder.productorder`) are used for storing user documents.

## Customization

- Modify the `User` struct and CRUD logic in `src/mongoClient.go` as needed for your application.
- Adjust the Lambda handler in `main.go` to support additional event types or business logic.

## Troubleshooting

- Ensure both containers are on the same Docker network.
- Use the container name (`my-mongo`) as the MongoDB host in your URI.
- Check Docker logs for connectivity issues.

---

**This project is a template for integrating AWS Lambda (Go) with MongoDB in a local Dockerized development environment.**