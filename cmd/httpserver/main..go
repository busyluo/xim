package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/login", login)

	// Start server
	err := e.Start(":8080")
	e.Logger.Fatal(err)
}

//// Handler
//func hello(c echo.Context) error {
//	return c.String(http.StatusOK, "Hello, World!")
//}
