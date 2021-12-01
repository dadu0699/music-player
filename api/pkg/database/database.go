package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBInstance return database.
var DBInstance *mongo.Client

// Init initializes the database.
func Init() {
	DBInstance = connect()
}

// connect establishes connection to the Mongo DB instance.
func connect() *mongo.Client {
	mongoURI := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("%s: %s", "Could not connect to the mongodb instance", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("%s: %s", "Could not connect to the mongodb instance", err)
	}

	return client
}

// OpenCollection makes a connection with a collection in the database.
func OpenCollection(databaseName string, collectionName string) *mongo.Collection {
	return DBInstance.Database(databaseName).Collection(collectionName)
}
