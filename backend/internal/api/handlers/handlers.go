package handlers

import (
	"nestle/internal/api"
	"nestle/internal/api/handlers/common"

	"github.com/labstack/echo/v4"
)

func AttachAllRoutes(s *api.Server) {
	s.Router.Routes = []*echo.Route{
		common.GetAgents(s),
		common.GetPoints(s),
	}
}
