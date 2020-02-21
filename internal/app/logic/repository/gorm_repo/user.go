package gorm_repo

import (
	"xim/internal/app/logic/entity"
)

type userRepository struct{}

func (userRepository) AddUser(user entity.User) error {
	panic("implement me")
}
