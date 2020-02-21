package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"xim/internal/app/httpserver/controller"
)

func Route(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/login", controller.Login)
	e.PUT("/user", controller.AddUser)
}
