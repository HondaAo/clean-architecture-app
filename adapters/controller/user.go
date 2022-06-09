package controller

import (
	"log"
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
func (controller *UserController) Index(c echo.Context) (err error) {
	log.Print(c)
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.userController.GetUser(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	return
}

func (controller *UserController) Show(c echo.Context) (err error) {
	users, err := controller.userController.GetUsers()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		log.Print(err)
		return err
	}
	log.Print(user)
	err = controller.userController.CreateUser(user)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, nil)
}
