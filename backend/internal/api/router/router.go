package router

import (
	"fmt"
	"nestle/internal/api"
	"nestle/internal/api/handlers"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Init(s *api.Server) {
	s.Echo = echo.New()

	s.Echo.Debug = true
	s.Echo.HideBanner = false

	if s.Config.Echo.EnableCORSMiddleware {
		s.Echo.Use(echoMiddleware.CORS())
	} else {
		fmt.Println("Disabling CORS middleware due to environment config")
	}

	s.Echo.Use(echoMiddleware.StaticWithConfig(echoMiddleware.StaticConfig{
		Skipper: nil,
		// Root directory from where the static content is served.
		Root: "build",
		// Index file for serving a directory.
		// Optional. Default value "index.html".
		Index: "index.html",
		// Enable HTML5 mode by forwarding all not-found requests to root so that
		// SPA (single-page application) can handle the routing.
		HTML5:      true,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: nil,
	}))

	s.Router = &api.Router{
		Routes:     nil,
		Root:       s.Echo.Group("/api/v1"),
		Management: s.Echo.Group("/management"),
		APIV1Auth:  s.Echo.Group("/api/v1/auth"),
		APIV1Push:  s.Echo.Group("/api/v1/push"),
	}
	handlers.AttachAllRoutes(s)
}
