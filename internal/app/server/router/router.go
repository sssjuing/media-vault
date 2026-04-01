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

	e.Group("/login").Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "web/login",
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))

	e.Group("/h5").Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "web/h5",
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))

	e.Group("/console").Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "web/console",
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))

	e.Validator = NewValidator()
	return e
}
