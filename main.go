package main

import (
	"github.com/HondaAo/go_blog_app/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	infrastructure.UserRoute(e)
	e.Logger.Fatal(e.Start(":4000"))
}
