package config

import (
	"os"
	"strconv"
)

type database struct {
	User     string
	Password string
	Port     int
	Host     string
	Name     string
	SslMode  string
}

type Config struct {
	Database database
}

func LoadConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 5432
	}
	db_settings := database{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     port,
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSLMODE"),
	}
	return &Config{
		Database: db_settings,
	}
}
