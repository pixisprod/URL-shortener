package cache

import "context"

type Cacher[T any] interface {
	Set(ctx context.Context, key string, val T) error
	Get(ctx context.Context, key string) (T, error)
}
