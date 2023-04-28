package main

import (
	"github.com/00mohamad00/url-shortener/internal/httpserver"
	"github.com/00mohamad00/url-shortener/internal/service"
	"github.com/00mohamad00/url-shortener/internal/storage"
	pkgStorage "github.com/00mohamad00/url-shortener/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(cmd, args); err != nil {
			logrus.WithError(err).Fatal("Failed to serve.")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, _ []string) error {
	conf := loadConfigOrPanic(cmd)

	db := getDatabaseOrPanic(conf.Database)
	defer closeDatabaseConnection(db)

	urlShortenerService := service.NewUrlShortenerService(storage.NewStorage(db))
	router := httpserver.NewRouter(urlShortenerService)

	err := router.Run(conf.server.Addr)
	if err != nil {
		return err
	}

	return nil
}

func loadConfigOrPanic(cmd *cobra.Command) *Config {
	conf, err := LoadConfig(cmd)
	if err != nil {
		logrus.WithError(err).Panic("Failed to load configurations")
	}
	return conf
}

func getDatabaseOrPanic(conf pkgStorage.DatabaseConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.GetURL()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func closeDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		logrus.WithError(err).Error("Error closing connection")
	}
	err = sqlDB.Close()
	if err != nil {
		return
	}
}
