package infrastructure

import (
	"github.com/HondaAo/go_blog_app/adapters/controller"
	"github.com/HondaAo/go_blog_app/infrastructure/db"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	userController := controller.NewUserController(db.NewSqlHandler())

	e.GET("/users/:id", userController.Index)
	e.GET("/users", userController.Show)
	e.POST("/user_create", userController.Create)
	e.POST("/signup", userController.SignUp)
}
