package gateways

import "github.com/HondaAo/go_blog_app/domain/entity"

type VideoRepository struct {
	SQLHandler
}

func (repo *VideoRepository) GetVideo(id int) (video entity.Video, err error) {
	if err = repo.Find(&video, id).Error; err != nil {
		return
	}
	return
}

func (repo *VideoRepository) GetVideos() (videos []entity.Video, err error) {
	if err = repo.Find(&videos).Error; err != nil {
		return
	}
	return
}

func (repo *VideoRepository) CreateVideo(video entity.Video) (err error) {
	if err = repo.Create(&video).Error; err != nil {
		return
	}
	return
}

func (repo *VideoRepository) UpdateVideo(video entity.Video) (err error) {
	if err = repo.Updates(video).Error; err != nil {
		return
	}
	return
}

func (repo *VideoRepository) SearchVideo(category string) (videos []entity.Video, err error) {
	if err = repo.Find(&videos, category).Error; err != nil {
		return
	}
	return
}
