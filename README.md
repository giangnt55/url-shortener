# URL Shortener

A simple URL shortening service built with **Golang** and **Gin**.

---

## 🚀 Features
- `/encode`: Converts a long URL into a short URL  
- `/decode`: Restores a short URL back to its original  
- Thread-safe in-memory store using `sync.RWMutex`  
- JSON input/output  
- Unit tests for both endpoints  

---

## ⚙️ Run Locally

#### 1️⃣ Install dependencies
```bash
go mod tidy
```

#### 2️⃣ Run the service
```bash
go run main.go
```
Server starts by default on: http://localhost:8080

---

## 📡 API Endpoints

#### ➤ POST /encode

Request body:
```json
{
  "url": "https://example.com/very/long/link"
}
```

Response:
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

#### ➤ POST /decode


Request body:
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

Response:
```json
{
  "original_url": "https://example.com/very/long/link"
}
```

---

## 🧪 Run Tests

#### 1️⃣ Run all tests
```bash
go test ./...
```

#### 2️⃣ Run with detailed output
```bash
go test -v
```