package storage

import "gorm.io/gorm"

type Record struct {
	gorm.Model

	Token string `gorm:"unique"`
	Url   string
}

type Storage interface {
	GetUrl(token string) (string, error)
	AddUrl(token, url string) error
}
