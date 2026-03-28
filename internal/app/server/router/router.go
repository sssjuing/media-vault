package router

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Logger = newLogger()
	e.Use(middleware.RequestLogger())

	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.Validator = NewValidator()
	return e
}
