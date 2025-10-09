package config

import (
	"os"
	"strconv"
)

type redis struct {
	Port int
	Host string
}

func InitRedisConfig() *redis {
	redis_port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		redis_port = 6379
	}
	return &redis{
		Port: redis_port,
		Host: os.Getenv("REDIS_HOST"),
	}
}
