package model

type EncodeRequest struct {
	URL string `json:"url" binding:"required"`
}

type EncodeResponse struct {
	ShortURL string `json:"short_url"`
}

type DecodeRequest struct {
	ShortURL string `json:"short_url" binding:"required"`
}

type DecodeResponse struct {
	OriginalURL string `json:"original_url"`
}
