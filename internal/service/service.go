package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/go-redis/redis/v9"
)

const  (
	ErrNotFound = "url not found"
	TokenLength = 8
)

type urlShortener struct {
	redisClient redis.Client
}

func newUrlShortener(redisClient redis.Client) *urlShortener {
	return &urlShortener{redisClient: redisClient}
}

func (s *urlShortener) getUrl(token string) (string, error) {
	url, err := s.redisClient.Get(context.Background(), token).Result()
	if err == redis.Nil {
		return "", errors.New(ErrNotFound)
	} else if err != nil {
		return "", err
	}
	return url, nil
}

func (s *urlShortener) saveUrl(url string) (string, error) {
	var token string
	for {
		token = generateToken()
		_, err := s.redisClient.Get(context.Background(), token).Result()
		if err == redis.Nil {
			break
		}
		if err != nil {
			return "", err
		}
	}
	err := s.redisClient.Set(context.Background(), token, url, 0).Err()
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
