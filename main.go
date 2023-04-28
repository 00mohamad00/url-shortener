package main

import (
	"github.com/00mohamad00/url-shortener/internal/httpserver"
	"github.com/00mohamad00/url-shortener/internal/service"
	"github.com/00mohamad00/url-shortener/internal/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	urlShortenerService := service.NewUrlShortenerService(storage.NewStorage(db))
	router := httpserver.NewRouter(urlShortenerService)

	err = router.Run("127.0.0.1:8887")
	if err != nil {
		panic(err)
	}
}
