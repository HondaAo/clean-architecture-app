package main

import (
	"net/http"

	"github.com/HondaAo/go_blog_app/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	infrastructure.UserRoute(e)

	// e.Use(middleware.JWT([]byte("secret")))
	e.GET("/restricted", restricted)
	e.Logger.Fatal(e.Start(":4000"))
}

func restricted(c echo.Context) error {
	return c.String(http.StatusOK, "helloworld")
}
