package handlers

import (
	"context"       //Used for timeout/cancel operations
	"encoding/json" //used for json reading and sending json body
	"net/http"
	"time" //to store created time stamp

	"url-shortener/internal/db"
	"url-shortener/internal/models"
	"url-shortener/internal/services"
)

type CreateAPIKeyRequest struct {
	Client string `json:"client"`
}

func CreateAPIKeyHandler(w http.ResponseWriter, r *http.Request) { //runs when user hits route /api/key
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateAPIKeyRequest //variable

	err := json.NewDecoder(r.Body).Decode(&req) //reads the request body in json
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Client == "" {
		http.Error(w, "client required", http.StatusBadRequest)
		return
	}

	key, err := services.GenerateAPIKey() //calls the func to generate the api key
	if err != nil {
		http.Error(w, "Failed to generate key", http.StatusInternalServerError)
		return
	}

	apiKeyCollection := db.Client.
		Database("urlshortener").
		Collection("api_keys") //get this collection
	now := time.Now()
	apiKey := models.APIKey{
		Key:       key,
		Client:    req.Client,
		Active:    true,
		CreatedAt: time.Now(),
		ExpiresAt: now.AddDate(0, 0, 30), //set th expiry after 30 days
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //should be inserted in 5 secs if not cancel it
	defer cancel()                                                          //cleans memory

	_, err = apiKeyCollection.InsertOne(ctx, apiKey) //stores document permanently in DB
	if err != nil {
		http.Error(w, "Failed to store key", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // alerts client that it is returning json

	json.NewEncoder(w).Encode(map[string]string{ //returns the key in json
		"api_key":    key,
		"expires_at": apiKey.ExpiresAt.Format(time.RFC3339), //returns the expiry time of API key
	})
}
