package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"urlShortner/storage"
)

type Operator string

const (
	REDISOPERATOR Operator = "redis"
	MAPOPERATOR   Operator = "map"
)

type Operators struct {
	operators map[Operator]Cache
}

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
}

func (o *Operators) NewCache(operator Operator) (Cache, error) {
	op, ok := o.operators[operator]
	if !ok {
		return nil, errors.New("invalid cache operator")
	}
	return op, nil
}

func InitialCacheOperators(input map[Operator]Cache) *Operators {
	return &Operators{operators: input}
}

type RedisCache struct {
	redis *storage.Redis
}

func NewRedisCache(redis *storage.Redis) *RedisCache {
	return &RedisCache{redis: redis}
}

type MapCache struct {
	mapCache map[string]string
	mu       *sync.Mutex
}

func NewMapCache(mapCache map[string]string, mu *sync.Mutex) (*MapCache, error) {
	if mu == nil {
		return nil, errors.New("mutex in not initialized")
	}
	return &MapCache{mapCache: mapCache, mu: mu}, nil
}
