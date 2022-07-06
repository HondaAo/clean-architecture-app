package controller

import (
	"log"
	"strconv"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/usecase"
	"github.com/HondaAo/go_blog_app/utils"
	"github.com/labstack/echo/v4"
)

type LikeController struct {
	likeController usecase.LikeUseCase
}

type Response struct {
	Username string `json:"user_name"`
	VideoId  int    `json:"video_id"`
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
	return c.JSON(200, likes)

}

func (controller *LikeController) Create(c echo.Context) (err error) {
	like := Response{}
	if err := c.Bind(&like); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	log.Print(like)
	id, _ := utils.GetSession(c, like.Username)
	intId, _ := strconv.Atoi(id)
	newLike := entity.Like{
		UserId:  intId,
		VideoId: like.VideoId,
	}
	err = controller.likeController.LikeVideo(newLike)
	if err != nil {
		return echo.NewHTTPError(500, "failed to create &video")
	}

	return c.JSON(200, nil)
}
