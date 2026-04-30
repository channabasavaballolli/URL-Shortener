package middleware

import (
	"context"
	"net/http"
	"time"

	"url-shortener/internal/db"
	"url-shortener/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func APIKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("X-API-Key") //reads the header the key

		if apiKey == "" {
			http.Error(w, "API key required", http.StatusUnauthorized)
			return
		}

		collection := db.Client.
			Database("urlshortener").
			Collection("api_keys") //accessing where api keys are stored

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //creats a context of 5 secs
		defer cancel()

		var keyDoc models.APIKey //key for result from mongo

		err := collection.FindOne(ctx, bson.M{ //searching the mongo
			"key":    apiKey,
			"active": true,
		}).Decode(&keyDoc) //if document found copy it to keyDoc

		if err != nil {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		if time.Now().After(keyDoc.ExpiresAt) {
			http.Error(w, "API key expired", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
