package controller

import (
	"strconv"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/usecase"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userController usecase.UserUsecase
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
func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	guest, err := controller.userController.GetUser(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, guest)
	return
}

func (controller *UserController) Index(c echo.Context) (err error) {
	guests, err := controller.userController.GetUsers()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, guests)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	input := entity.User{
		Username: user.Username,
		Email:    user.Email,
	}
	err = controller.userController.CreateUser(input)

	return
}
