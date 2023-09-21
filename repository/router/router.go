package router

import (
	"context"

	"github.com/mhshahin/magellan/models"
	"github.com/redis/go-redis/v9"
)

type RoutingRepository struct {
	redisClient *redis.Client
}

func NewRoutingRepository(redisClient *redis.Client) Router {
	return &RoutingRepository{
		redisClient: redisClient,
	}
}

func (r *RoutingRepository) GetRoute(ctx context.Context, searchParams models.GeoSearchParams) ([]string, error) {
	res, err := r.redisClient.GeoSearch(ctx, searchParams.Name, &redis.GeoSearchQuery{
		Longitude:  searchParams.Longitude,
		Latitude:   searchParams.Latitude,
		Radius:     searchParams.Radius,
		RadiusUnit: searchParams.RadiusUnit,
		Sort:       searchParams.Sort,
	}).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	return res, nil

}

func (r *RoutingRepository) AddPoints(ctx context.Context, addParams models.GeoAddParams) error {
	locations := []*redis.GeoLocation{}

	for _, loc := range addParams.Members {
		locations = append(locations, &redis.GeoLocation{
			Name:      loc.Name,
			Longitude: loc.Longitude,
			Latitude:  loc.Latitude,
		})
	}

	err := r.redisClient.GeoAdd(ctx, addParams.Key, locations...).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RoutingRepository) GetMemberCoordinates(ctx context.Context, key string, members []string) ([]models.GeoMember, error) {
	res, err := r.redisClient.GeoPos(ctx, key, members...).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	geoMembers := []models.GeoMember{}

	for i, member := range res {
		geoMembers = append(geoMembers, models.GeoMember{
			Order:     i + 1,
			Name:      members[i],
			Longitude: member.Longitude,
			Latitude:  member.Latitude,
		})
	}

	return geoMembers, nil
}
