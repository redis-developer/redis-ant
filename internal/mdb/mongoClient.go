package mdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ramantehlan/redis-ant/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout = 300
)

// Client is the MongoDB client.
var Client *mongo.Client

// InitMongoDB is to initialize the database
func InitMongoDB() {
	// Open connection for mongo
	Client = OpenConnection(
		config.Config.MongoURI,
		config.Config.MongoUser,
		config.Config.MongoPass,
	)
}

// OpenConnection is to open a connection with MongoDB and get the client
func OpenConnection(URI, username, password string) *mongo.Client {
	// Step: Get the context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	// Step: Create a connection string using credentials
	connectionURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/test?retryWrites=true&w=majority", username, password, URI)

	// Step: Create a client using the connectionURI
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI), options.Client().SetMinPoolSize(100))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Step: Connect client to the database using the context
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	// Step: Ping client to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	return client
}
