package main

import (
	"url-shortener/handler"
	"url-shortener/service"
	"url-shortener/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := storage.NewMemoryStore()
	svc := service.NewURLService(store, "http://localhost:8080")
	h := handler.NewURLHandler(svc)

	r.POST("/encode", h.Encode)
	r.POST("/decode", h.Decode)

	r.Run(":8080")
}
