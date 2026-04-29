package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var idCounter int = 0                  //This simulates database auto-increment.
// var urlStore = make(map[string]string) // A global var for mapping
var Client *mongo.Client
var Collection *mongo.Collection

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI)) //Go app connects to MongoDB server running on your laptop.
	if err != nil {
		panic(err)
	}

	Collection = Client.Database("urlshortener").Collection("urls") // using db: urlshortener and collection urls

	fmt.Println("MongoDB connected")
}
