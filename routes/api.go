package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mhshahin/magellan/handlers"
)

func InitializeRoutes(e *echo.Echo, h *handlers.Handlers) {
	api := e.Group("/api/v1")

	api.POST("/calculate-route", h.Router.CalculateRouteHandler())
	api.POST("/add-points", h.Router.AddPointsHandler())
}
