package gateways

import (
	"github.com/HondaAo/go_blog_app/domain/entity"
)

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) GetUser(id int) (user *entity.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUserByEmail(email string) (user *entity.User, err error) {
	if err = repo.Find(&user, "email = ?", email).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUsers() (users []entity.User, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) CreateUser(user entity.User) (*entity.User, error) {
	if err := repo.Create(&user).Error; err != nil {
		return nil, err
	}
	createdUser := &user
	return createdUser, nil
}
