package main

import (
	"net/http"

	zone "ofa/bin/modules/utilities/zone/hendlers"
	"ofa/bin/repositories"

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
	e.GET("/sum/:n1/:n2", func(c echo.Context) error {
		return c.String(http.StatusOK, repositories.Summation(c))
	})
	e.GET("/subtract/:n1/:n2", func(c echo.Context) error {
		return c.String(http.StatusOK, repositories.Subtraction(c))
	})
	e.GET("/divide/:n1/:n2", func(c echo.Context) error {
		return c.String(http.StatusOK, repositories.Division(c))
	})

	grupUtilities := e.Group("/utilities")
	zone.Init(grupUtilities)

}
