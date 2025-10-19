package config

import (
	"os"
	"strconv"
)

type hash struct {
	Charset string
	Length  int
}

func InitHashConfig() *hash {
	hash_length, err := strconv.Atoi(os.Getenv("HASH_LENGTH"))
	if err != nil {
		hash_length = 6
	}
	return &hash{
		Charset: os.Getenv("HASH_CHARSET"),
		Length:  hash_length,
	}

}
