package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	// Get MongoDB connection string from .env file
	MongoDb := os.Getenv("MONGODB_URL")
	if MongoDb == "" {
		log.Fatal("MONGODB_URL is not set in the .env file")
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create client options and connect
	clientOptions := options.Client().ApplyURI(MongoDb)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	// Ping the database to check if the connection is successful
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB Ping Error:", err)
	}

	fmt.Println("Connected to MongoDB.")
	return client
}

var Client *mongo.Client = DBinstance()
func OpenCollection(client *mongo.Client, collectionName string)*mongo.Collection{
	var collection *mongo.Collection = client.Database("Cluster 0").Collection(collectionName)
	return collection
}