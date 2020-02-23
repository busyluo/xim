package service

import (
	"xim/internal/app/logic/entity"
)

type user struct{}

var User user

func (user) NewUser(u entity.User) error {
	return nil
}
