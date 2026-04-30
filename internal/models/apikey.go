package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APIKey struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`      //Mongo document ID
	Key       string             `bson:"key" json:"key"`               //API Key which is generated
	Client    string             `bson:"client" json:"client"`         //the client who owns the API key
	Active    bool               `bson:"active" json:"active"`         //to disable later
	CreatedAt time.Time          `bson:"created_at" json:"created_at"` // the time stamp created
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"` //the time stamp when api key will expire
}
