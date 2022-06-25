package controller

import (
	"strconv"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/usecase"
	"github.com/labstack/echo/v4"
)

type LikeController struct {
	likeController usecase.LikeUseCase
}

func NewLikeController(sqlHandler gateways.SQLHandler) *LikeController {
	return &LikeController{
		likeController: usecase.LikeUseCase{
			LikeInteractor: &gateways.LikeRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *LikeController) Index(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("user_id"))
	likes, err := controller.likeController.GetLikes(id)
	if err != nil {
		echo.NewHTTPError(500, "failed to get video")
		return
	}
	c.JSON(200, likes)
	return
}

func (controller *LikeController) Create(c echo.Context) (err error) {
	like := entity.Like{}
	if err := c.Bind(&like); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	err = controller.likeController.LikeVideo(like)
	if err != nil {
		return echo.NewHTTPError(500, "failed to create &video")
	}

	return c.JSON(200, nil)
}
