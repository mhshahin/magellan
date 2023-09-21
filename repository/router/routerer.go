package router

import (
	"context"

	"github.com/mhshahin/magellan/models"
)

type Router interface {
	GetRoute(ctx context.Context, origin models.GeoSearchParams) ([]string, error)
	AddPoints(ctx context.Context, addParams models.GeoAddParams) error
	GetMemberCoordinates(ctx context.Context, key string, members []string) ([]models.GeoMember, error)
}
