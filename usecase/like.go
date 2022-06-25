package usecase

import (
	"github.com/HondaAo/go_blog_app/domain/entity"
	"github.com/HondaAo/go_blog_app/domain/repository"
)

type LikeUseCase struct {
	LikeInteractor repository.LikeRepository
}

func (repo *LikeUseCase) LikeVideo(like entity.Like) (err error) {
	err = repo.LikeInteractor.CreateLike(like)
	return
}

func (repo *LikeUseCase) GetLikes(id int) (videoIds []int, err error) {
	videoIds, err = repo.LikeInteractor.GetLikes(id)
	return
}
