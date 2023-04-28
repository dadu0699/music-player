package main

import (
	"context"
	"log"
	"os"
	"time"

	"music-service/pkg/app"
	"music-service/pkg/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func setupDatabase(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}

func run() error {
	// // Setup database connection
	mongoURI := os.Getenv("MONGO_URI")
	db, err := setupDatabase(mongoURI)
	if err != nil {
		return err
	}

	// Create storage dependency
	storage := repository.NewStorage(db)

	// Create Music service dependency
	log.Println(storage)

	// Create router dependecy
	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Create server dependency and run it
	server := app.NewServer(router)
	err = server.Run()
	if err := router.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	if err := run(); err != nil {
		log.Fatalf("This is the startup error: %v", err)
		os.Exit(1)
	}
}
