package repository

import "github.com/HondaAo/go_blog_app/domain/entity"

type VideoRepository interface {
	GetVideo(id int) (video entity.Video, err error)
	GetVideos() (videos []entity.Video, err error)
	CreateVideo(video entity.Video) (err error)
	UpdateVideo(video entity.Video) (err error)
	SearchVideo(category string) (videos []entity.Video, err error)
}
