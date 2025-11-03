package handler

import (
	"net/http"

	"url-shortener/model"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	svc *service.URLService
}

func NewURLHandler(svc *service.URLService) *URLHandler {
	return &URLHandler{svc: svc}
}

// POST /encode
func (h *URLHandler) Encode(c *gin.Context) {
	var req model.EncodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	shortURL, err := h.svc.Encode(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.EncodeResponse{ShortURL: shortURL})
}

// POST /decode
func (h *URLHandler) Decode(c *gin.Context) {
	var req model.DecodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	original, err := h.svc.Decode(req.ShortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.DecodeResponse{OriginalURL: original})
}
