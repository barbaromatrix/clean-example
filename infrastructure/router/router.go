package router

import (
	"github.com/barbaromatrix/clean-example/interface/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter function
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })

	return e
}
