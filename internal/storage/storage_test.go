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
	db   *gorm.DB
	impl *Impl
}

func (s *StorageSuite) SetupSuite() {
	var err error
	s.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		s.T().Fatalf("Failed to connect to database: %v", err)
	}

	err = s.db.AutoMigrate(&storage.Record{})
	if err != nil {
		s.T().Fatalf("Failed to migrate database schema: %v", err)
	}

	s.impl = &Impl{db: s.db}
}

func (s *StorageSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.T().Errorf("Failed to get database connection: %v", err)
	}
	sqlDB.Close()
}

func (s *StorageSuite) TearDownTest() {
	s.db.Exec("DELETE FROM records")
}

func (s *StorageSuite) TestAddUrl() {
	err := s.impl.AddUrl("abc123", "http://example.com")
	s.Nil(err)

	var rec storage.Record
	err = s.db.First(&rec, "token = ?", "abc123").Error
	s.Nil(err)
	s.Equal("http://example.com", rec.Url)
	s.Equal("abc123", rec.Token)
}

func (s *StorageSuite) TestAddUrl_DuplicateToken() {
	err := s.impl.AddUrl("abc123", "http://example.com")
	s.Nil(err)

	err = s.impl.AddUrl("abc123", "http://example2.com")
	s.Error(err)
}
gi
func (s *StorageSuite) TestGetUrl() {
	err := s.db.Create(&storage.Record{
		Token: "abc123",
		Url:   "http://example.com",
	}).Error
	s.Nil(err)

	url, err := s.impl.GetUrl("abc123")
	s.Nil(err)
	s.Equal("http://example.com", url)
}

func (s *StorageSuite) TestGetUrl_NotFound() {
	_, err := s.impl.GetUrl("nonexistent")
	s.Equal(storage.ErrNotFound, err)
}

func TestStorageSuite(t *testing.T) {
	suite.Run(t, new(StorageSuite))
}
