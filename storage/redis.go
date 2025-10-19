package storage

import (
	"github.com/redis/go-redis/v9"

	"urlShortner/config"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(config config.RedisConfig) *Redis {
	return &Redis{Client: redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})}
}
