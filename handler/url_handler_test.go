package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"url-shortener/model"
	"url-shortener/service"
	"url-shortener/storage"

	"github.com/gin-gonic/gin"
)

// setup test router
func setup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	store := storage.NewMemoryStore()
	svc := service.NewURLService(store, "http://localhost:8080")
	h := NewURLHandler(svc)

	r := gin.Default()
	r.POST("/encode", h.Encode)
	r.POST("/decode", h.Decode)
	return r
}

// test encode -> decode flow
func TestEncodeDecode(t *testing.T) {
	r := setup()

	// encode
	reqBody := model.EncodeRequest{URL: "https://example.com"}
	b, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/encode", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("encode failed: got %d", w.Code)
	}

	var enc model.EncodeResponse
	json.Unmarshal(w.Body.Bytes(), &enc)
	if enc.ShortURL == "" {
		t.Fatal("empty short url")
	}

	// decode
	decReq := model.DecodeRequest{ShortURL: enc.ShortURL}
	b, _ = json.Marshal(decReq)
	req, _ = http.NewRequest("POST", "/decode", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("decode failed: got %d", w.Code)
	}

	var dec model.DecodeResponse
	json.Unmarshal(w.Body.Bytes(), &dec)
	if dec.OriginalURL != reqBody.URL {
		t.Fatalf("expected %s, got %s", reqBody.URL, dec.OriginalURL)
	}
}
