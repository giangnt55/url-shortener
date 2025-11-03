package main

import (
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
	svc := service.NewURLService(store, "http://localhost:8080")
	h := handler.NewURLHandler(svc)

	r.POST("/encode", h.Encode)
	r.POST("/decode", h.Decode)

	r.Run(":8080")
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
