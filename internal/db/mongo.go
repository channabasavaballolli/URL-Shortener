package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client //two global variables start with capital letter so they are exported outside the package
var Collection *mongo.Collection

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //creates a context with timeout
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
