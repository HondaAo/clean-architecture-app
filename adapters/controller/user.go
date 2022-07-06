package controller

import (
	"log"
	"strconv"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/usecase"
	"github.com/HondaAo/go_blog_app/utils"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userController usecase.UserUsecase
}

type SignupInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username  string `json:"nickname"`
	Email     string `json:"email"`
	Icon      string `json:"picture"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserController(sqlHandler gateways.SQLHandler) *UserController {
	return &UserController{
		userController: usecase.UserUsecase{
			UserInteractor: &gateways.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}
func (controller *UserController) Index(c echo.Context) (err error) {
	log.Print(c)
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.userController.GetUserById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	return
}

func (controller *UserController) Show(c echo.Context) (err error) {
	users, err := controller.userController.GetAllUsers()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	user := UserResponse{}
	if err := c.Bind(&user); err != nil {
		log.Print(err)
		return err
	}
	getUser, err := controller.userController.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(500, err)
		return
	}
	if getUser != nil {
		cookie, _ := utils.GetSession(c, getUser.Username)
		log.Print(cookie)
		id, _ := strconv.Atoi(cookie)
		if id != getUser.Id || cookie == "" {
			utils.NewSession(c, getUser.Username, strconv.Itoa(getUser.Id))
		}
		return c.JSON(200, getUser)
	}
	newUser := entity.User{
		Username:  user.Username,
		Email:     user.Email,
		Icon:      user.Icon,
		UpdatedAt: user.UpdatedAt,
	}
	createdUser, err := controller.userController.CreateOneUser(newUser)
	if err != nil {
		return c.JSON(500, err)
	}

	utils.NewSession(c, createdUser.Username, "secret")
	return c.JSON(200, createdUser)
}

// func (controller *UserController) SignUp(c echo.Context) error {
// 	input := new(SignupInput)

// 	if err := c.Bind(input); err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	if err := controller.userController.SignUp(c.Request().Context(), input.Username, input.Email, input.Password); err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	return c.JSON(200, nil)
// }

// func (controller *UserController) SignIn(c echo.Context) error {
// 	input := new(SigninInput)

// 	if err := c.Bind(input); err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	token, err := controller.userController.SignIn(c.Request().Context(), input.Email, input.Password)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(200, token)
// }
