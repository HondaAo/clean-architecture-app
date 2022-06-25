package usecase

import (
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/domain/repository"
)

type VideoUsecase struct {
	VideoInteractor repository.VideoRepository
	LikeInteractor  repository.LikeRepository
}

func (repo *VideoUsecase) GetVideoById(id int) (video entity.Video, err error) {
	video, err = repo.VideoInteractor.GetVideo(id)
	return
}

func (repo *VideoUsecase) GetVideosByLikeIds(id int) (videos []entity.Video, err error) {
	ids, err := repo.LikeInteractor.GetLikes(id)
	videos, err = repo.VideoInteractor.GetVideosByIds(ids)
	return
}

func (repo *VideoUsecase) GetAllVideos() (videos []entity.Video, err error) {
	videos, err = repo.VideoInteractor.GetVideos()
	return
}

func (repo *VideoUsecase) RegisterVideo(video entity.Video) (err error) {
	err = repo.VideoInteractor.CreateVideo(video)
	return
}

func (repo *VideoUsecase) UpdateVideo(video entity.Video) (err error) {
	err = repo.VideoInteractor.UpdateVideo(video)
	return
}

func (repo *VideoUsecase) SearchVideo(category string) (videos []entity.Video, err error) {
	videos, err = repo.VideoInteractor.SearchVideo(category)
	return
}
