package service

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/00mohamad00/url-shortener/pkg/storage"
	"github.com/00mohamad00/url-shortener/pkg/urlshortener"
)

const (
	TokenLength = 2
)

type urlShortenerService struct {
	storage storage.Storage
}

func NewUrlShortenerService(storage storage.Storage) urlshortener.Service {
	return &urlShortenerService{storage: storage}
}

func (s *urlShortenerService) GetUrl(token string) (string, error) {
	return s.storage.GetUrl(token)
}

func (s *urlShortenerService) AddUrl(url string) (string, error) {
	var token string
	for {
		token = generateToken()
		_, err := s.storage.GetUrl(token)
		if err == storage.ErrNotFound {
			break
		}
		if err != nil {
			return "", err
		}
	}
	err := s.storage.AddUrl(token, url)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken() string {
	b := make([]byte, TokenLength)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
