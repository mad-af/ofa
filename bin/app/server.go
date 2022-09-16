package main

import (
	"net/http"

	zone "ofa/bin/modules/utilities/zone/hendlers"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func (s *Server) Routes() {
	e := s.Echo
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	grupUtilities := e.Group("/utilities")
	zone.Init(grupUtilities)

}
