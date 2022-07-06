package repository

import "github.com/HondaAo/go_blog_app/domain/entity"

type VideoRepository interface {
	GetVideo(id int) (video entity.Video, err error)
	GetVideos() (videos []entity.Video, err error)
	CreateVideo(video entity.Video) (err error)
	UpdateVideo(video entity.Video, id int) (err error)
	SearchVideoByCategory(category string) (videos []entity.Video, err error)
	SearchVideoBySeries(series string) (videos []entity.Video, err error)
	GetVideosByIds(ids []int) (videos []entity.Video, err error)
}
