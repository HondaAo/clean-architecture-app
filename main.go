package main

import (
	"errors"
	"fmt"

	"github.com/HondaAo/go_blog_app/infrastructure"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	infrastructure.UserRoute(e)
	config := middleware.JWTConfig{
		SigningKey: []byte("secret"),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("SECRET_KEY"), nil
			}
			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}
	e.Use(middleware.JWTWithConfig(config))
	e.Logger.Fatal(e.Start(":4000"))
}
