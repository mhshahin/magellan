package config

import (
	"context"
	"net"

	"github.com/mhshahin/magellan/models"
	"github.com/redis/go-redis/v9"
)

func InitializeRedis(redisCfg models.RedisCfg) (*redis.Client, error) {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: net.JoinHostPort(redisCfg.Host, redisCfg.Port),
		},
	)

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}
