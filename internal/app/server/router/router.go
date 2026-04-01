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

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "web/dist",
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))

	e.Validator = NewValidator()
	return e
}
