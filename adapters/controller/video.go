package controller

import (
	"log"
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

type VideoInput struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Script   string `json:"script"`
	View     int    `json:"view"`
	Category string `json:"category"`
	Series   string `json:"series"`
	End      int    `json:"end"`
	Level    string `json:"level"`
}

func (controller *VideoController) Index(c echo.Context) (err error) {
	log.Print(c)
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.videoController.GetVideoById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	return
}

func (controller *VideoController) Show(c echo.Context) (err error) {
	videos, err := controller.videoController.GetAllVideos()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, videos)
	return
}

func (controller *VideoController) Create(c echo.Context) (err error) {
	video := entity.Video{}
	if err := c.Bind(&video); err != nil {
		log.Print(err)
		return err
	}
	log.Print(video)
	err = controller.videoController.RegisterVideo(video)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, nil)
}
