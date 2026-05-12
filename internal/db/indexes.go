package db

import (
	"context"
	"log" //for printing erros if index fails

	"go.mongodb.org/mongo-driver/bson" //usd to define index fields
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes() {

	//this is URL short_code index
	urlIndex := mongo.IndexModel{ //it is index blueprint
		Keys: bson.M{
			"short_code": 1, //creates a index on short field 1 means stores index in ascending order
		},
		Options: options.Index().SetUnique(true), //means it will not allow duplicate values
	}

	_, err := Collection.Indexes(). //mongo DB index manager
					CreateOne(context.Background(), urlIndex) //creates one index in database with context

	if err != nil {
		log.Fatal("Failed to create short_code index:", err)
	}

	// this is for API key index
	apiKeyCollection := Client. //accessing the api_keys collection
					Database("urlshortener").
					Collection("api_keys")

	apiKeyIndex := mongo.IndexModel{ //Index API key field
		Keys: bson.M{
			"key": 1,
		},
		Options: options.Index().SetUnique(true), //ensures no duplicate api keys
	}

	_, err = apiKeyCollection.Indexes().
		CreateOne(context.Background(), apiKeyIndex)

	if err != nil {
		log.Fatal("Failed to create api key index:", err)
	}
}
