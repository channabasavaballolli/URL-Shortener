Here’s your **updated `README.md`** with all the new features you implemented (API keys, auth middleware, rate limiting, expiry) — clean, professional, and ready for GitHub.

---

````md
# URL Shortener with MongoDB, Redis, API Keys & Rate Limiting

A production-style URL Shortener backend built in Go, featuring random TinyURL-style short codes, MongoDB persistent storage, Redis caching, API key authentication, rate limiting, and scalable middleware architecture.

---

## Architecture Overview

### Request Flow

```mermaid
graph TD
    Client --> API[Go HTTP Server]
    API --> Auth[API Key Middleware]
    Auth --> RateLimit[Rate Limit Middleware]
    RateLimit --> Redis[Redis Cache]
    Redis -->|Cache Miss| MongoDB[MongoDB Database]
    Redis -->|Cache Hit| Redirect[Instant Redirect]
    MongoDB --> Redirect
````

---

## URL Creation Flow

```mermaid
graph TD
    Client --> Shorten[POST /shorten]
    Shorten --> Auth[Validate API Key]
    Auth --> RateLimit[Check Rate Limit]
    RateLimit --> Validate[Validate URL]
    Validate --> Generate[Generate Random Code / Alias]
    Generate --> MongoDB[Store in MongoDB]
    MongoDB --> Response[Return Short URL]
```

---

## Key Features

* **Professional Project Structure** using `cmd`, `handlers`, `models`, `db`, `middleware`, `services`
* **Random 8 Character Short Codes** similar to TinyURL
* **Custom Alias Support**
* **MongoDB Persistent Storage**
* **Redis Cache Layer** for faster redirects
* **URL Validation** using Go `net/url`
* **API Key Generation & Storage**
* **API Key Authentication Middleware**
* **30-Day API Key Expiration**
* **Redis Fixed Window Rate Limiting (per API key)**
* **Secure Header-Based Access (`X-API-Key`)**
* **Scalable Middleware Architecture**

---

## Tech Stack

* **Language:** Go (Golang)
* **Database:** MongoDB
* **Cache Layer:** Redis
* **Authentication:** API Key (custom implementation)
* **Rate Limiting:** Redis (Fixed Window)
* **API Testing:** Postman
* **Version Control:** Git + GitHub

---

## API Endpoints

---

### 1. Generate API Key

```http
POST /api/key
```

#### Request Body

```json
{
  "client": "frontend-app"
}
```

#### Response

```json
{
  "api_key": "abc123...",
  "expires_at": "2026-05-29T10:00:00Z"
}
```

---

### 2. Create Short URL (Protected)

```http
POST /shorten
```

#### Headers

```http
X-API-Key: your_api_key
```

#### Request Body

```json
{
  "url": "https://google.com"
}
```

#### With Custom Alias

```json
{
  "url": "https://google.com",
  "alias": "google"
}
```

#### Response

```json
{
  "short_url": "http://localhost:8080/54z7zk3d"
}
```

---

### 3. Redirect URL

```http
GET /{short_code}
```

Example:

```http
GET /54z7zk3d
```

Redirects user to original long URL.

---

## API Security & Rate Limiting

### API Key Authentication

All `/shorten` requests require:

```http
X-API-Key: your_key
```

Validation checks:

* Key exists in MongoDB
* Key is active
* Key is not expired

---

### API Key Expiration

* Default validity: **30 days**
* Expired keys return:

```http
401 Unauthorized
```

---

### Rate Limiting (Fixed Window)

* Limit: **10 requests per minute per API key**
* Implemented using Redis:

```text
rate_limit:<apikey>
```

---

### Behavior

| Requests        | Response              |
| --------------- | --------------------- |
| 1–10 per minute | 200 OK                |
| >10 per minute  | 429 Too Many Requests |

---

## Redis Caching Logic

### First Request

```text
/54z7zk3d
→ Redis Miss
→ MongoDB Lookup
→ Save in Redis
→ Redirect
```

### Next Requests

```text
/54z7zk3d
→ Redis Hit
→ Instant Redirect
```

---

## Environment Variables

```env
MONGO_URI=mongodb://localhost:27017
PORT=8080
BASE_URL=http://localhost:8080
```

---

## Execution Guide

### 1. Start MongoDB

```text
mongodb://localhost:27017
```

---

### 2. Start Redis

```text
localhost:6379
```

---

### 3. Run Application

```bash
go run ./cmd/main.go
```

---

### Expected Output

```text
MongoDB connected
Redis connected
Server Listening on port 8080
```

---

## Testing with Postman

### Step 1: Generate API Key

```http
POST /api/key
```

### Step 2: Use API Key

```http
POST /shorten
X-API-Key: your_key
```

---

## Project Structure

```text
url-shortener/
│── cmd/
│   └── main.go
│
│── internal/
│   ├── db/
│   │   ├── mongo.go
│   │   └── redis.go
│   │
│   ├── handlers/
│   │   ├── urlhandler.go
│   │   └── apikey_handler.go
│   │
│   ├── middleware/
│   │   ├── apikey.go
│   │   └── ratelimit.go
│   │
│   ├── models/
│   │   ├── url.go
│   │   └── apikey.go
│   │
│   ├── utils/
│   │   └── base62.go
│   │
│   └── services/
│       └── apikey_service.go
```

---

## Core Engineering Concepts

### Cache Aside Pattern

Application checks Redis first, then MongoDB.

---

### Fixed Window Rate Limiting

Simple and fast Redis-based rate limiting using `INCR + EXPIRE`.

---

### API Key Lifecycle

* Creation → Active → Expiration (30 days)
* Secure access control for API endpoints

---

## Future Improvements

* API Key Revocation & Rotation
* Usage Analytics per API Key
* Tiered Rate Limits (Free / Pro)
* Sliding Window or Token Bucket Rate Limiting
* URL Expiration
* User Authentication System
* Docker Deployment
* CI/CD Pipeline

---

## Author

Built by Channabasava Ballolli as a backend engineering project using Go, MongoDB, Redis, and scalable API design.

```



