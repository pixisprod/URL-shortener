package config

import (
	"os"
	"strconv"
)

type app struct {
	RetryInterval int
}

type Config struct {
	Database *database
	App      *app
	Redis    *redis
	Hash     *hash
}

func LoadConfig() *Config {
	retry_interval, err := strconv.Atoi(os.Getenv("APP_RETRY_INTERVAL"))
	if err != nil {
		retry_interval = 5
	}
	app_settings := app{
		RetryInterval: retry_interval,
	}

	return &Config{
		Database: InitDatabaseConfig(),
		App:      &app_settings,
		Redis:    InitRedisConfig(),
		Hash:     InitHashConfig(),
	}
}
