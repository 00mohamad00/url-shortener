package storage

import (
	"testing"

	"github.com/00mohamad00/url-shortener/pkg/storage"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StorageSuite struct {
	suite.Suite
	db      *gorm.DB
	Storage storage.Storage
}

func (s *StorageSuite) SetupSuite() {
	var err error
	s.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.Nil(err)
	s.Storage = NewStorage(s.db)
}

func (s *StorageSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	s.Nil(err)
	sqlDB.Close()
}

func (s *StorageSuite) TearDownTest() {
	s.db.Exec("DELETE FROM records")
}

func (s *StorageSuite) TestAddUrl() {
	err := s.Storage.AddUrl("abc123", "http://example.com")
	s.Nil(err)

	var rec storage.Record
	err = s.db.First(&rec, "token = ?", "abc123").Error
	s.Nil(err)
	s.Equal("http://example.com", rec.Url)
	s.Equal("abc123", rec.Token)
}

func (s *StorageSuite) TestAddUrl_DuplicateToken() {
	err := s.Storage.AddUrl("abc123", "http://example.com")
	s.Nil(err)

	err = s.Storage.AddUrl("abc123", "http://example2.com")
	s.Error(err)
}

func (s *StorageSuite) TestGetUrl() {
	err := s.db.Create(&storage.Record{
		Token: "abc123",
		Url:   "http://example.com",
	}).Error
	s.Nil(err)

	url, err := s.Storage.GetUrl("abc123")
	s.Nil(err)
	s.Equal("http://example.com", url)
}

func (s *StorageSuite) TestGetUrl_NotFound() {
	_, err := s.Storage.GetUrl("nonexistent")
	s.Equal(storage.ErrNotFound, err)
}

func TestStorageSuite(t *testing.T) {
	suite.Run(t, new(StorageSuite))
}
