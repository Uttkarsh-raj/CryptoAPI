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

func DBInstance() *mongo.Client {

	err := godotenv.Load() //To connect locally uncomment this and add the .env file
	if err != nil {
		log.Fatal("Error loading the environment file.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	MongoURL := os.Getenv("MONGODB_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoURL))
	if err != nil {
		log.Fatalf("Error connecting to the url: %s", err)
	}

	fmt.Println("Connected to MongoDb.")

	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	clusterName := os.Getenv("CLUSTER")
	var collection *mongo.Collection = client.Database(clusterName).Collection(collectionName)
	return collection
}
