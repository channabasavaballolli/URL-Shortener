package middleware

import (
	"net/http"
	"strconv"
	"time"

	"url-shortener/internal/db"
)

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc { //like API middleware runs before actual handler
	return func(w http.ResponseWriter, r *http.Request) { //return func

		apiKey := r.Header.Get("X-API-Key") //reads api key

		redisKey := "rate_limit:" + apiKey //creates redis key

		count, err := db.RedisClient.Incr(db.Ctx, redisKey).Result() //increases the count by 1 if the key is not present a key is created and set to 1
		if err != nil {
			http.Error(w, "Redis error", http.StatusInternalServerError) //if increment fails
			return
		}

		if count == 1 {
			db.RedisClient.Expire(db.Ctx, redisKey, time.Minute) //if this is the first request of this key it will be automatically deleted after 1 minute.
		}

		if count > 10 { //blocks if over limit
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Retry-After", "60") //tells to retry after 60 seconds

			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		w.Header().Set("X-RateLimit-Limit", "10")
		w.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(10-count, 10)) //helful for client to  see remaining requests.

		next.ServeHTTP(w, r)
	}
}
