package common

import (
	"nestle/internal/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAgents(s *api.Server) *echo.Route {
	return s.Router.Root.GET("/agents", getAgentsHandler(s))
}

type agentsPayload struct {
	Days   []int  `query:"day"`
	Region string `query:"region"`
}

func getAgentsHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload agentsPayload
		err := c.Bind(&payload)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if s.Repository == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		agents := s.Repository.SelectAgents(c.Request().Context(), payload.Region, payload.Days)
		return c.JSON(http.StatusOK, agents)
	}
}
