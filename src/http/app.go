package http

import "github.com/labstack/echo"

func CreateServer() (e *echo.Echo) {
	e = echo.New()

	return e
}