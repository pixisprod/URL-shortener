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

func InitDatabaseConfig() *database {
	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		db_port = 5432
	}
	return &database{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     db_port,
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSLMODE"),
	}
}
