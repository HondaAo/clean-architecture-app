package usecase

import (
	"context"
	"crypto/sha1"
	"time"

	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/domain/repository"
)

type UserUsecase struct {
	UserInteractor repository.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

// type AuthClaims struct {
// 	jwt.StandardClaims
// 	User *entity.User `json:"user"`
// }

func (repo *UserUsecase) GetUserById(id int) (user entity.User, err error) {
	user, err = repo.UserInteractor.GetUser(id)
	return
}

func (repo *UserUsecase) GetAllUsers() (users []entity.User, err error) {
	users, err = repo.UserInteractor.GetUsers()
	return
}

func (repo *UserUsecase) CreateOneUser(user entity.User) (err error) {
	err = repo.UserInteractor.CreateUser(user)
	return
}

func (repo *UserUsecase) SignUp(c context.Context, username string, email string, password string) (err error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(repo.hashSalt))

	user := &entity.User{
		Username: username,
		Email:    email,
		Password: password,
		Bio:      "",
	}

	err = repo.UserInteractor.CreateUser(*user)
	return err
}
