package usecase

import (
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/domain/repository"
)

type UserUsecase struct {
	UserInteractor repository.UserRepository
}

func (userInt *UserUsecase) GetUser(id int) (user entity.User, err error) {
	user, err = userInt.UserInteractor.GetUser(id)
	return
}

func (userInt *UserUsecase) GetUsers() (users []entity.User, err error) {
	users, err = userInt.UserInteractor.GetUsers()
	return
}

func (userInt *UserUsecase) CreateUser(user entity.User) (err error) {
	err = userInt.CreateUser(user)
	return
}
