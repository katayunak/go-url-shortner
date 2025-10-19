package cache

import (
	"context"
	"time"
)

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Client.Get(ctx, key).Result()
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return r.redis.Client.Set(ctx, key, value, ttl).Err()
}
