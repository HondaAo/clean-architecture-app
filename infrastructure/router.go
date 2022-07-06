package infrastructure

import (
	"github.com/HondaAo/go_blog_app/adapters/controller"
	"github.com/HondaAo/go_blog_app/infrastructure/db"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	userController := controller.NewUserController(db.NewSqlHandler())
	videoController := controller.NewVideoController(db.NewSqlHandler())
	likeController := controller.NewLikeController(db.NewSqlHandler())

	e.GET("/api/users/:id", userController.Index)
	e.GET("/api/users", userController.Show)
	e.POST("/api/user_create", userController.Create)
	e.GET("/api/videos", videoController.Show)
	e.GET("/api/videos/:id", videoController.Index)
	e.PUT("/api/videos/:id", videoController.Update)
	e.GET("/api/videos/categories/:category", videoController.Search)
	e.GET("/api/videos/series/:series", videoController.Series)
	e.GET("/api/like_videos/:user_id", videoController.Like)
	e.POST("/api/video_create", videoController.Create)
	e.POST("/api/like_video", likeController.Create)
	// e.POST("/api/signup", userController.SignUp)
	// e.POST("/api/signin", userController.SignIn)
}
