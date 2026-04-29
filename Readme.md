````md
# URL Shortener with MongoDB & Redis Caching

A production-style URL Shortener backend built in Go, featuring random TinyURL-style short codes, MongoDB persistent storage, Redis caching, custom aliases, URL validation, and scalable project architecture.

---

## Architecture Overview

### Request Flow

```mermaid
graph TD
    Client --> API[Go HTTP Server]
    API --> Redis[Redis Cache]
    Redis -->|Cache Miss| MongoDB[MongoDB Database]
    Redis -->|Cache Hit| Redirect[Instant Redirect]
    MongoDB --> Redirect
````

---

## URL Creation Flow

```mermaid
graph TD
    Client --> Shorten[POST /shorten]
    Shorten --> Validate[Validate URL]
    Validate --> Generate[Generate Random Code / Alias]
    Generate --> MongoDB[Store in MongoDB]
    MongoDB --> Response[Return Short URL]
```

---

## Key Features

* **Professional Project Structure** using `cmd`, `handlers`, `models`, `db`, `utils`, `services`
* **Random 8 Character Short Codes** similar to TinyURL
* **Custom Alias Support**
* **MongoDB Persistent Storage**
* **Redis Cache Layer** for faster redirects
* **URL Validation** using Go `net/url`
* **Environment Variables** for deployment-ready configuration
* **Fast Redirect System**
* **Scalable Backend Design**

---

## Tech Stack

* **Language:** Go (Golang)
* **Database:** MongoDB
* **Cache Layer:** Redis / Memurai
* **API Testing:** Postman / PowerShell
* **Version Control:** Git + GitHub

---

## API Endpoints

### 1. Create Short URL

```http
POST /shorten
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

### 2. Redirect URL

```http
GET /{short_code}
```

Example:

```http
GET /54z7zk3d
```

Redirects user to original long URL.

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

This reduces MongoDB reads and improves response speed.

---

## Environment Variables

Create `.env` or set manually:

```env
MONGO_URI=mongodb://localhost:27017
PORT=8080
BASE_URL=http://localhost:8080
```

---

## Execution Guide

### 1. Start MongoDB

Ensure MongoDB is running locally:

```text
mongodb://localhost:27017
```

---

### 2. Start Redis / Memurai

Default:

```text
localhost:6379
```

---

### 3. Run Application

```powershell
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

## Example Testing with PowerShell

### Create URL

```powershell
Invoke-RestMethod -Method Post `
-Uri "http://localhost:8080/shorten" `
-ContentType "application/json" `
-Body '{"url":"https://google.com"}'
```

---

### Open Short URL

```text
http://localhost:8080/54z7zk3d
```

---

## Core Engineering Concepts

### Base62 Encoding

Uses:

```text
0-9 a-z A-Z
```

to generate compact short codes.

---

### Cache Aside Pattern

Application checks Redis first, then MongoDB on cache miss.

---

### Separation of Concerns

* `handlers` → HTTP logic
* `db` → Database connections
* `models` → Structs
* `utils` → Helpers
* `services` → Business logic

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
│   │   └── url_handler.go
│   │
│   ├── models/
│   │   └── url.go
│   │
│   ├── utils/
│   │   └── base62.go
│   │
│   └── services/
│       └── counter.go
```

---

## Future Improvements

* Click Analytics
* Rate Limiting using Redis
* Expiring URLs
* User Accounts / API Keys
* QR Code Generation
* Admin Dashboard
* Docker Deployment

---

## Author

Built by Channabasava Ballolli as a backend engineering project using Go, MongoDB, and Redis.

```
```
