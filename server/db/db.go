package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectToDb(collectionName string) *mongo.Collection {

	// In a production environment (like Render), variables are set directly.
	// The .env file is only for local development.
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Could not load .env file. Using environment variables from the system.")
		}
	}

	connectionString := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	colName := collectionName

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	Collection = client.Database(dbName).Collection(colName)

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return Collection
}
