package usecase

import (
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/domain/repository"
	"github.com/dgrijalva/jwt-go/v4"
)

type UserUsecase struct {
	UserInteractor repository.UserRepository
}

type AuthClaims struct {
	jwt.StandardClaims
	User *entity.User `json:"user"`
}

const (
	hashSalt       = "secret"
	secret         = "secret"
	expireDuration = 60 * 60 * 24 * 7
)

func (repo *UserUsecase) GetUserById(id int) (user *entity.User, err error) {
	user, err = repo.UserInteractor.GetUser(id)
	return
}

func (repo *UserUsecase) GetUserByEmail(email string) (user *entity.User, err error) {
	user, err = repo.UserInteractor.GetUserByEmail(email)
	return
}

func (repo *UserUsecase) GetAllUsers() (users []entity.User, err error) {
	users, err = repo.UserInteractor.GetUsers()
	return
}

func (repo *UserUsecase) CreateOneUser(user entity.User) (createdUser *entity.User, err error) {
	createdUser, err = repo.UserInteractor.CreateUser(user)
	return
}

// func (repo *UserUsecase) SignUp(c context.Context, username string, email string, password string) (err error) {
// 	pwd := sha1.New()
// 	pwd.Write([]byte(password))
// 	pwd.Write([]byte(hashSalt))

// 	user := entity.User{
// 		Username: username,
// 		Email:    email,
// 	}

// 	err = repo.UserInteractor.CreateUser(user)
// 	return
// }

// func (repo *UserUsecase) SignIn(c context.Context, email, password string) (string, error) {
// 	pwd := sha1.New()
// 	pwd.Write([]byte(password))
// 	pwd.Write([]byte(hashSalt))
// 	password = fmt.Sprintf("%x", pwd.Sum(nil))

// 	user, err := repo.UserInteractor.GetUserByEmail(email)

// 	if err != nil {
// 		return "", err
// 	}

// 	claims := AuthClaims{
// 		User: &user,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: jwt.At(time.Now().Add(expireDuration)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	return token.SignedString([]byte(secret))
// }
