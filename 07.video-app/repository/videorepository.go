package repository

import (
	"video-app/config"
	"video-app/domain/entity"
)

type VideoRepository interface {
	Save(video entity.Video) entity.Video
	Find(id uint) *entity.Video
	FindAll() []entity.Video
}

type videoRepository struct {
	//videos []entity.Video
}

func NewVideoRepository() VideoRepository {

	return &videoRepository{}
}

func (videoRepo *videoRepository) Save(video entity.Video) entity.Video {

	//videoRepo.videos = append(videoRepo.videos, video)

	config.DB.Save(&video)

	return video
}

func (videoRepo *videoRepository) Find(id uint) *entity.Video {

	var video entity.Video
	config.DB.First(&video, id) // Here I am not preloading with Author(whcih is of type Person)

	//for _, video := range videoRepo.videos {
	//
	//	if video.Id == id {
	//		return &video
	//	}
	//}

	return &video
}

func (videoRepo *videoRepository) FindAll() []entity.Video {

	var videos []entity.Video

	config.DB.Preload("Author").Find(&videos) // Author is the name of the field in Video struct.

	return videos
}
