package transport

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RestServer contains global server
type RestServer struct {
	echo   *echo.Echo
	server *http.Server
}

// Setup creates router instance
func (s *RestServer) Setup() {
	s.echo = echo.New()

	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
		AllowHeaders: []string{
			echo.HeaderContentType,
			echo.HeaderAuthorization,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
	}))

	s.server = &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: s.echo,
	}
	addRoutes(s.echo)
}

// Start begins the server to listen
func (s *RestServer) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
