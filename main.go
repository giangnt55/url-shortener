package main

import (
	"os"
	"url-shortener/handler"
	"url-shortener/service"
	"url-shortener/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(CORSMiddleware())

	store := storage.NewMemoryStore()

	// Determine base URL
	baseURL := getBaseURL()

	svc := service.NewURLService(store, baseURL)
	h := handler.NewURLHandler(svc)

	// Serve static UI file (index.html) for testing API in browser
	r.StaticFile("/", "./index.html")

	r.POST("/encode", h.Encode)
	r.POST("/decode", h.Decode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

func getBaseURL() string {
	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		return baseURL
	}

	// Check if running on Render
	if renderURL := os.Getenv("RENDER_EXTERNAL_URL"); renderURL != "" {
		return renderURL
	}

	// Default to localhost for local development
	return "http://localhost:8080"
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Requested-With, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
