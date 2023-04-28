package storage

import (
	"errors"

	"github.com/00mohamad00/url-shortener/pkg/storage"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Impl struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) storage.Storage {
	err := db.AutoMigrate(&storage.Record{})
	if err != nil {
		logrus.Fatalf("Failed to migrate db: %v\n", err)
	}
	return &Impl{db: db}
}

func (i Impl) GetUrl(token string) (string, error) {
	var rec storage.Record
	err := i.db.Where("token = ?", token).First(&rec).Error
	if err != nil {
		return "", checkError(err)
	}
	return rec.Url, nil
}

func (i Impl) AddUrl(token, url string) error {
	rec := storage.Record{Token: token, Url: url}
	err := i.db.Create(&rec).Error
	return err
}

func checkError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return storage.ErrNotFound
	}
	return err
}
