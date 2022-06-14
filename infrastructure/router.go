package infrastructure

import (
	"github.com/HondaAo/go_blog_app/adapters/controller"
	"github.com/HondaAo/go_blog_app/infrastructure/db"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	userController := controller.NewUserController(db.NewSqlHandler())

	e.GET("/api/users/:id", userController.Index)
	e.GET("/api/users", userController.Show)
	e.POST("/api/user_create", userController.Create)
	e.POST("/api/signup", userController.SignUp)
	e.POST("/api/signin", userController.SignIn)
}
