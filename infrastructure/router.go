package infrastructure

import (
	"github.com/HondaAo/go_blog_app/adapters/controller"
	"github.com/HondaAo/go_blog_app/infrastructure/db"
	"github.com/labstack/echo/v4"
)

func UserRoute(*echo.Echo) {
	e := echo.New()

	userController := controller.NewUserController(db.NewSqlHandler())

	e.GET("/users/:id", userController.Show)
	e.GET("/users", userController.Index)
}
