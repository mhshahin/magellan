package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mhshahin/magellan/models"
	"github.com/mhshahin/magellan/repository"
)

type Resp map[string]interface{}

type RouterHandlers struct {
	repo *repository.Repository
}

func NewRouterHandlers(repo *repository.Repository) *RouterHandlers {
	return &RouterHandlers{
		repo: repo,
	}
}

func (rh *RouterHandlers) CalculateRouteHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		searchParams := models.GeoSearchParams{}
		err := c.Bind(&searchParams)
		if err != nil {
			fmt.Println(err)
			return echo.ErrInternalServerError
		}

		members, err := rh.repo.Router.GetRoute(c.Request().Context(), searchParams)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}

		routes, err := rh.repo.Router.GetMemberCoordinates(c.Request().Context(), searchParams.Name, members)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, Resp{"message": "ok", "data": routes})
	}
}

func (rh *RouterHandlers) AddPointsHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		addParams := models.GeoAddParams{}
		err := c.Bind(&addParams)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}

		err = rh.repo.Router.AddPoints(c.Request().Context(), addParams)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, Resp{"message": "ok", "data": addParams})
	}
}
