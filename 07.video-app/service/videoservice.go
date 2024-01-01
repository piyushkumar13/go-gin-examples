package service

import (
	"video-app/domain/entity"
	"video-app/repository"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	Find(id uint) *entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	vr repository.VideoRepository
}

func NewVideoService(videoRepository repository.VideoRepository) VideoService {

	return &videoService{
		vr: videoRepository,
	}
}

func (v videoService) Save(video entity.Video) entity.Video {

	return v.vr.Save(video)
}

func (v videoService) Find(id uint) *entity.Video {

	return v.vr.Find(id)
}

func (v videoService) FindAll() []entity.Video {
	return v.vr.FindAll()
}
