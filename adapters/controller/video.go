package controller

import (
	"strconv"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/usecase"
	"github.com/labstack/echo/v4"
)

type VideoController struct {
	videoController usecase.VideoUsecase
}

func NewVideoController(sqlHandler gateways.SQLHandler) *VideoController {
	return &VideoController{
		videoController: usecase.VideoUsecase{
			VideoInteractor: &gateways.VideoRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *VideoController) Index(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	video, err := controller.videoController.GetVideoById(id)
	if err != nil {
		echo.NewHTTPError(500, "failed to get video")
		return
	}
	c.JSON(200, video)
	return
}

func (controller *VideoController) Show(c echo.Context) (err error) {
	videos, err := controller.videoController.GetAllVideos()
	if err != nil {
		echo.NewHTTPError(500, "failed to get video")
		return
	}
	c.JSON(200, videos)
	return
}

func (controller *VideoController) Create(c echo.Context) (err error) {
	video := entity.Video{}
	if err := c.Bind(&video); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	//仮置き
	// video.Id = 1
	err = controller.videoController.RegisterVideo(video)
	if err != nil {
		return echo.NewHTTPError(500, "failed to create &video")
	}

	return c.JSON(200, nil)
}

func (controller *VideoController) Watch(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	video, err := controller.videoController.GetVideoById(id)
	if err != nil {
		echo.NewHTTPError(500, "failed to get video")
		return
	}
	video.View += 1
	if err = controller.videoController.UpdateVideo(video); err != nil {
		return
	}
	return c.JSON(200, nil)
}

func (controller *VideoController) Update(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	video, err := controller.videoController.GetVideoById(id)
	if err != nil {
		echo.NewHTTPError(500, "failed to get video")
		return
	}
	updatedVideo := entity.Video{}
	if err := c.Bind(&updatedVideo); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	video = updatedVideo
	if err = controller.videoController.UpdateVideo(video); err != nil {
		return
	}
	return c.JSON(200, nil)
}

func (controller *VideoController) Search(c echo.Context) (err error) {
	category := c.Param("category")
	videos, err := controller.videoController.SearchVideo(category)
	if err != nil {
		echo.NewHTTPError(500, "failed to get video")
		return
	}
	return c.JSON(200, videos)
}

func (controller *VideoController) Like(c echo.Context) (err error) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	videos, err := controller.videoController.GetVideosByLikeIds(user_id)
	if err != nil {
		echo.NewHTTPError(500, "failed to get video")
		return
	}
	return c.JSON(200, videos)
}
