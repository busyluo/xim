package controller

import (
	"github.com/labstack/echo/v4"
	"xim/internal/app/logic/entity"
	"xim/internal/app/logic/service"
)

func NewUser(ctx echo.Context) error {
	var u entity.User
	u.Nickname = ctx.Param("name")
	return service.User.NewUser(u)
}
