package repository

import (
	"github.com/mhshahin/magellan/repository/router"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Router router.Router
}

func NewRepository(redisClient *redis.Client) *Repository {
	return &Repository{
		Router: router.NewRoutingRepository(redisClient),
	}
}
