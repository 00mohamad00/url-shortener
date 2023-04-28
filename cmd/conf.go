package main

import (
	"strings"

	"github.com/00mohamad00/url-shortener/internal/httpserver"
	"github.com/00mohamad00/url-shortener/pkg/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	server   httpserver.Config
	Database storage.DatabaseConfig
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	viper.SetDefault("Server.Addr", "127.0.0.1:8888")

	// Read Config from ENV
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
