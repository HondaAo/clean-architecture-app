package usecase

import (
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/domain/repository"
)

type UserUsecase struct {
	UserInteractor repository.UserRepository
}

func (userInteractor *UserUsecase) GetUser(id int) (user entity.User, err error) {
	user, err = userInteractor.UserInteractor.GetUser(id)
	return
}

func (userInteractor *UserUsecase) GetUsers() (users []entity.User, err error) {
	users, err = userInteractor.UserInteractor.GetUsers()
	return
}
