package models

type Request struct { //// struct for url -type string and it is in json format
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

type Response struct {
	ShortURL string `json:"short_url"`
}

type URLDocument struct {
	ShortCode   string `bson:"short_code"`
	OriginalURL string `bson:"original_url"` //bson tags are MongoDB version of JSON tags.
}

type Counter struct {
	ID  string `bson:"_id"`
	Seq int    `bson:"seq"` // This matches MongoDB document.
}
