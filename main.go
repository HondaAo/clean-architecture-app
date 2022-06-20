package main

import (
	"github.com/HondaAo/go_blog_app/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	infrastructure.UserRoute(e)

	// e.Use(middleware.JWT([]byte("secret")))
	e.Logger.Fatal(e.Start(":4000"))
}
