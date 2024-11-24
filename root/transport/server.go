package transport

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RestServer struct {
	echo   *echo.Echo
	server *http.Server
}

func (s *RestServer) Setup() {
	s.echo = echo.New()
	s.server = &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: s.echo,
	}
	addRoutes(s.echo)
}

func (s *RestServer) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
