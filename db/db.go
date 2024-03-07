package db

import (
	"context"
	"time"

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

	// Check if the database already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = DB.Client().Ping(ctx, nil)
	if err != nil {
		log.Fatal().Msgf("Database 'restweb' does not exist: %v", err)
	}

	log.Info().Msg("Database 'restweb' already exists")

	// Create a collection in the database
	scancollection := DB.Collection("restscan")

	exists, err := CollectionExists(context.Background(), DB, "restscan")
	if err != nil {
		log.Fatal().Msgf("Error occured while checking collection: %v", err)
	}

	if exists {
		log.Info().Msg("Collection 'restscan' exists")
	} else {
		log.Info().Msg("Collection 'restscan' does not exist")
	}
}

func CollectionExists(ctx context.Context, db *mongo.Database, collectionName string) (bool, error) {
	collectionNames, err := db.ListCollectionNames(ctx, nil)
	if err != nil {
		return false, err
	}
	for _, name := range collectionNames {
		if name == collectionName {
			return true, nil
		}
	}
	return false, nil
}
