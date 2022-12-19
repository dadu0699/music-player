package repository

import "go.mongodb.org/mongo-driver/mongo"

// Storage is the interface for the repository
type Storage interface{}

type storage struct {
	db *mongo.Client
}

// NewStorage creates a new storage instance
func NewStorage(db *mongo.Client) Storage {
	return &storage{db: db}
}
