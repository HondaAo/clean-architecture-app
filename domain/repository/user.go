package repository

import (
	"github.com/HondaAo/go_blog_app/domain/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUser(id int) (entity.User, error)
	GetUsers() ([]entity.User, error)
}
