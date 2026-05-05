
# URL Shortener (Go + MongoDB + Redis)

A production-style URL shortener backend built in Go with API key authentication, Redis caching, and rate limiting.

This project demonstrates how to design a scalable backend system using middleware, caching, and request control mechanisms.

---

## Key Features

- Short URL generation with optional custom aliases  
- MongoDB for persistent storage  
- Redis caching for fast redirects (cache-aside pattern)  
- API key authentication using header (`X-API-Key`)  
- 30-day API key expiration  
- Redis-based rate limiting (fixed window: 10 req/min per key)  
- Clean modular architecture (`handlers`, `db`, `models`, `middleware`, `services`)  

---

## How It Works

### Short URL Flow

1. Client sends request to `/shorten`
2. API key is validated (MongoDB)
3. Rate limit is checked (Redis)
4. URL is validated and short code generated
5. Stored in MongoDB
6. Short URL returned

### Redirect Flow

1. Request hits `/short_code`
2. Redis is checked first  
3. If cache miss → MongoDB lookup  
4. Result cached in Redis  
5. User redirected  

---

## API Endpoints

### Generate API Key

```http
POST /api/key
````

```json
{
  "client": "frontend-app"
}
```

---

### Create Short URL

```http
POST /shorten
```

Headers:

```http
X-API-Key: your_api_key
```

Body:

```json
{
  "url": "https://google.com"
}
```

---

### Redirect

```http
GET /{short_code}
```

---

## Rate Limiting

* Fixed window rate limiting using Redis
* Key format: `rate_limit:<apikey>`
* Limit: **10 requests per minute per API key**

---

## API Key Lifecycle

* Generated via `/api/key`
* Stored in MongoDB
* Valid for **30 days**
* Can be revoked (`active = false`)
* Middleware validates:

  * existence
  * active status
  * expiration

---

## Project Structure

```
internal/
  db/
  handlers/
  middleware/
  models/
  services/
```

---

## Running Locally

### Start MongoDB

```
mongodb://localhost:27017
```

### Start Redis

```
localhost:6379
```

### Run Server

```bash
go run ./cmd/main.go
```

---

## Why This Project Matters

This project goes beyond a basic URL shortener by implementing:

* authentication via API keys
* middleware-based request validation
* Redis caching and rate limiting
* clean backend architecture

It reflects real-world backend design patterns used in production systems.

---

## Future Improvements

* Usage analytics per API key
* Tier-based rate limiting
* Token bucket implementation
* URL expiration
* Docker setup

---

## Author

Channabasava Ballolli

```
```
