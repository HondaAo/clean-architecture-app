package gateways

import (
	"log"

	"github.com/HondaAo/go_blog_app/domain/entity"
)

type LikeRepository struct {
	SQLHandler
}

func (repo LikeRepository) CreateLike(like entity.Like) (err error) {
	if err = repo.Create(&like).Error; err != nil {
		return err
	}
	return nil
}

func (repo LikeRepository) DeleteLike(id int) (err error) {
	if err = repo.Delete(id).Error; err != nil {
		return err
	}
	return nil
}

func (repo LikeRepository) GetLikes(id int) (ids []int, err error) {
	likes := []entity.Like{}
	if err = repo.WhereUserId(&likes, id).Error; err != nil {
		return nil, err
	}
	log.Print(likes)
	ids = []int{}
	for _, v := range likes {
		ids = append(ids, v.VideoId)
	}
	log.Print(ids)
	return ids, err
}
