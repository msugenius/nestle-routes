package api

import (
	"fmt"
	"nestle/internal/config"
	"nestle/internal/database/repository/query"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Routes     []*echo.Route
	Root       *echo.Group
	Management *echo.Group
	APIV1Auth  *echo.Group
	APIV1Push  *echo.Group
}

type Server struct {
	Repository *query.Core
	Config     config.Server
	Echo       *echo.Echo
	Router     *Router
}

func NewServer(cfg config.Server) *Server {
	return &Server{
		Config: cfg,
	}
}

func (s *Server) RegisterDb(repository *query.Core) *Server {
	s.Repository = repository
	return s
}

func (s Server) StartServer() {
	s.Echo.Start(fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
}
