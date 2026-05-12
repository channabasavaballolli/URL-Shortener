package models

import "time"

type Request struct { //// struct for url -type string and it is in json format
	Name  string `json:"name"`
	URL   string `json:"url"`
	Alias string `json:"alias"`
	Email string `json:"email"`
}

type Response struct { //a server sends back an url in json format
	ShortURL string `json:"short_url"`
}

type URLDocument struct {
	ShortCode   string    `bson:"short_code"`
	OriginalURL string    `bson:"original_url"`                 //bson tags are MongoDB version of JSON tags.
	CreatedAt   time.Time `bson:"created_at" json:"created_at"` // the time stamp created
	ExpiresAt   time.Time `bson:"expires_at" json:"expires_at"`
}

type Counter struct {
	ID  string `bson:"_id"`
	Seq int    `bson:"seq"` // This matches MongoDB document.
}
