package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// 定义全区变量 为了保证执行顺序 初始化均在main中执行
var (
	// gorm mysql db connection
	db *gorm.DB
	// redis client
	// rdb *redis.Client
	// global cache
	// cc *cache.Codec
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
