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

	// session
	// GET /session # 获取会话信息
	// POST /session # 创建新的会话（登入）
	// PUT /session # 更新会话信息
	// DELETE /session # 销毁当前会话（登出）
	e.POST("/session", controller.Login)

	// user
	// GET /user/:id # 获取id用户的信息
	// POST /user # 创建新的用户（注册）
	// PUT /user/:id # 更新id用户的信息
	// DELETE /user/:id # 删除id用户（注销）
	e.POST("/user", controller.NewUser)

	// device
	e.POST("/device", controller.Login)
}
