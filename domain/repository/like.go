package repository

import "github.com/HondaAo/go_blog_app/domain/entity"

type LikeRepository interface {
	CreateLike(like entity.Like) (err error)
	DeleteLike(id int) (err error)
	GetLikes(id int) (videoIds []int, err error)
}
