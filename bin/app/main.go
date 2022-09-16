package main

import (
	u "ofa/bin/pkg/utils"
	r "ofa/bin/pkg/response"

	"github.com/labstack/echo/v4"
)

func main() {
	var s = &Server{echo.New()}
	e := s.Echo
	e.Validator = u.NewValidator()
	e.HTTPErrorHandler = r.NewHTTPErrorHandler

	s.Routes()

	e.Logger.Fatal(e.Start(":9000"))
}