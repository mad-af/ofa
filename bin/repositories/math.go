package repositories

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func Summation(c echo.Context) string {
	n1, _ := strconv.Atoi(c.Param("n1"))
	n2, _ := strconv.Atoi(c.Param("n2"))
	return strconv.Itoa(n1 + n2)
}

func Subtraction(c echo.Context) string {
	n1, _ := strconv.Atoi(c.Param("n1"))
	n2, _ := strconv.Atoi(c.Param("n2"))
	return strconv.Itoa(n1 - n2)
}

func Division(c echo.Context) string {
	n1, _ := strconv.Atoi(c.Param("n1"))
	n2, _ := strconv.Atoi(c.Param("n2"))
	return strconv.Itoa(n1 / n2)
}
