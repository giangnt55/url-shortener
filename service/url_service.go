package service

import (
	"errors"
	"strings"

	"url-shortener/storage"

	"github.com/teris-io/shortid"
)

type URLService struct {
	store *storage.MemoryStore
	base  string
}

func NewURLService(store *storage.MemoryStore, baseURL string) *URLService {
	return &URLService{store: store, base: strings.TrimRight(baseURL, "/")}
}

func (s *URLService) Encode(originalURL string) (string, error) {
	id, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	shortURL := s.base + "/" + id
	s.store.Save(id, originalURL)
	return shortURL, nil
}

func (s *URLService) Decode(shortURL string) (string, error) {
	parts := strings.Split(shortURL, "/")
	id := parts[len(parts)-1]

	if original, ok := s.store.Get(id); ok {
		return original, nil
	}
	return "", errors.New("short URL not found")
}
