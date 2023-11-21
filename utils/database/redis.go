package database

import (
	"github.com/go-redis/redis/v8"
	"waf/config"
)

// NewRedisDB 实例化redis
func NewRedisDB(config config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
}
