package handlers

import (
	"context"
	"encoding/json" // needed for decoding the json
	"fmt"
	"net/http"
	"os"

	"url-shortener/internal/db"
	"url-shortener/internal/models"
	"url-shortener/internal/utils"

	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "URL Shortener Running")
// }

func URLShortener(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // we will check if the requst method is a Post method or not if not we will return err
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed) //return method not allowed
		return
	}
	var req models.Request

	err := json.NewDecoder(r.Body).Decode(&req) // we decode the json store in req var
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	parsedURL, err := url.ParseRequestURI(req.URL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" { //checks the given input is url or not.
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	// idCounter++
	// code := encodeBase62(idCounter) //convert numeric ID into short code
	var code string

	// if req.Alias != "" {
	// 	code = req.Alias
	// } else {
	// 	id := services.GetNextID()
	// 	code = utils.EncodeBase62(id)
	// }
	if req.Alias != "" {
		code = req.Alias
	} else {
		for {
			code = utils.GenerateRandomCode(8)

			var existing models.URLDocument

			err = db.Collection.FindOne(
				context.Background(),
				bson.M{"short_code": code},
			).Decode(&existing)

			if err != nil {
				break
			}
		}
	}

	doc := models.URLDocument{
		ShortCode:   code,
		OriginalURL: req.URL, //Means create a record in Go memory.
	}
	var existing models.URLDocument

	err = db.Collection.FindOne(
		context.Background(),
		bson.M{"short_code": code},
	).Decode(&existing)

	if err == nil {
		http.Error(w, "Alias already taken", http.StatusConflict)
		return
	}
	_, err = db.Collection.InsertOne(context.Background(), doc) //saves the record to mongo DB
	if err != nil {
		http.Error(w, "Database insert failed", http.StatusInternalServerError)
		return
	}

	baseURL := os.Getenv("BASE_URL")

	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	resp := models.Response{
		ShortURL: baseURL + "/" + code,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp) //we will return the json

	if err != nil {
		http.Error(w, "Response Failed", http.StatusInternalServerError)
		return
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "URL Shortener Running")
		return
	}

	code := r.URL.Path[1:] //[1:] because it is like /1

	longURL, err := db.RedisClient.Get(db.Ctx, code).Result() //Searches Redis cache for key

	if err == nil {
		fmt.Println("Code hit :", code)
		http.Redirect(w, r, longURL, http.StatusFound) //if found then immediate redirect
		return
	}
	//only if redis can't find
	var result models.URLDocument

	err = db.Collection.FindOne( // finds matching document
		context.Background(),
		bson.M{"short_code": code}, // filters the object
	).Decode(&result) //Put DB result into Go struct.

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	db.RedisClient.Set(db.Ctx, code, result.OriginalURL, 0) //save the result found into redis cache

	fmt.Println("Cache Miss -> Saved:", code)

	http.Redirect(w, r, result.OriginalURL, http.StatusFound) //redirect
}
