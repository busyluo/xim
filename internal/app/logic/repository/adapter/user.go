package adapter

import "xim/internal/app/logic/entity"

type UserRepository interface {
	AddUser(user entity.User) error
}
