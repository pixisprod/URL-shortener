package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacher struct {
	r *redis.Client
}

func (c *RedisCacher) Get(ctx context.Context, k string) (string, error) {
	v, err := c.r.Get(ctx, k).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return v, nil
}

func (c *RedisCacher) Set(ctx context.Context, k string, v string) error {
	_, err := c.r.Set(ctx, k, v, time.Duration(15)*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func NewRedisCacher(r *redis.Client) *RedisCacher {
	return &RedisCacher{r}
}

func InitRedisCacher(addr string, timeout int) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()
	for {
		_, err := c.Ping(ctx).Result()
		if err == nil {
			log.Println("[Redis] ✅ Connection set")
			break
		}

		log.Printf("[Redis] ❌ Connection attempt failed\n⏳ Retrying in %d seconds..\n", timeout)
		time.Sleep(time.Duration(timeout) * time.Second)
	}

	return c
}
