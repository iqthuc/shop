package cache

import (
	"context"
	"shop/internal/infrastructure/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cf *config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cf.DataSourceName(),
		Password: cf.Password,
		DB:       cf.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("failed to connect to Redis: " + err.Error())
	}

	return rdb
}
