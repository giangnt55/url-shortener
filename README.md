# 🔗 URL Shortener

A simple URL shortening service built with **Golang** and **Gin**.

## 🌐 Live Demo

**Try it now:** [https://go-url-shortener-euss.onrender.com/](https://go-url-shortener-euss.onrender.com/)


---

## ✨ Features

- 🔗 **URL Encoding**: Converts long URLs into short, shareable links
- 🔍 **URL Decoding**: Restores short URLs back to their original form
- 🔒 **Thread-Safe**: In-memory store with `sync.RWMutex` for concurrent access

---

## 🏗️ Tech Stack

- **Backend**: Go (Golang) with Gin Framework
- **Frontend**: HTML5, CSS3, JavaScript
- **Storage**: In-memory data store
- **Deployment**: Render
---

## ⚙️ Run Locally

### Prerequisites
- Go 1.24 or higher
- Git

### Installation Steps

#### 1️⃣ Clone the repository
```bash
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

#### 2️⃣ Install dependencies
```bash
go mod tidy
```

#### 3️⃣ Run the service
```bash
go run main.go
```

Server starts by default on: **http://localhost:8080**

#### 4️⃣ Open your browser
Navigate to `http://localhost:8080` to use the web interface!

---

## 📡 API Endpoints

### ➤ POST `/encode`
Create a shortened URL from a long URL.

**Request:**
```json
{
  "url": "https://example.com/very/long/link/to/shorten"
}
```

**Response:**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/encode \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/very/long/link"}'
```

---

### ➤ POST `/decode`
Retrieve the original URL from a shortened URL.

**Request:**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

**Response:**
```json
{
  "original_url": "https://example.com/very/long/link/to/shorten"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/decode \
  -H "Content-Type: application/json" \
  -d '{"short_url":"http://localhost:8080/abc123"}'
```

---

## 🧪 Testing

### Run all tests
```bash
go test ./...
```

### Run with detailed output
```bash
go test -v
```

---

## 🙏 Acknowledgments

- Built with [Gin Web Framework](https://gin-gonic.com/)
- UI design enhanced with AI assistance
- Deployed on [Render](https://render.com/)

---