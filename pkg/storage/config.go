package storage

import "fmt"

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (config *DatabaseConfig) GetURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port,
	)
}
