package db

import (
	"context"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DB *mongo.Database

func InitDB() {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal().Msgf("Failed to connect to MongoDB:", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal().Msgf("Failed to ping MongoDB: %v", err)
	}

	// Set the database
	DB = client.Database("restweb")
}
