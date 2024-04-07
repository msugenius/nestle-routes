package common

import (
	"nestle/internal/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPoints(s *api.Server) *echo.Route {
	return s.Router.Root.GET("/points", getPointsHandler(s))
}

type pointsPayload struct {
	Route string `query:"route"`
}

func getPointsHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload pointsPayload
		err := c.Bind(&payload)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if s.Repository == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		points := s.Repository.SelectPoints(c.Request().Context(), payload.Route)
		return c.JSON(http.StatusOK, points)
	}
}
