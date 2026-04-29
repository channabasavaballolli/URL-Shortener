package services

import (
	"context"

	"url-shortener/internal/db"
	"url-shortener/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetNextID() int {
	counterCollection := db.Client.Database("urlshortener").Collection("counters") //use collection counters

	filter := bson.M{"_id": "url_counter"} // finds the document

	update := bson.M{ //updates th seq
		"$inc": bson.M{
			"seq": 1,
		},
	}

	options := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	var updated models.Counter //This is where MongoDB result gets stored.

	err := counterCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options,
	).Decode(&updated) //finds the document increases th seq and stores result in updated

	if err != nil {
		panic(err)
	}

	return updated.Seq //returns the new id
}
