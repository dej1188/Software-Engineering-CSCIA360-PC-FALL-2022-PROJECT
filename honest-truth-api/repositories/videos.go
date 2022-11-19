package repositories

import (
	"github.com/goodcodeguy/honest-truth-api/repositories/models"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type videosRepository Repository

type VideosRepository interface {
	AllVideos() []models.Video
	FindVideoById(id string) (models.Video, error)
	CreateVideo(v models.Video) models.Video
}

func NewVideosRepository(db *gorm.DB) VideosRepository {
	return &videosRepository{
		DB: db,
	}
}

func (r *videosRepository) AllVideos() []models.Video {
	var videos []models.Video

	r.DB.Preload("Notations").Find(&videos)

	return videos
}

func (r *videosRepository) FindVideoById(id string) (models.Video, error) {
	var video models.Video

	if err := r.DB.Preload("Notations").First(&video, "id = ?", id).Error; err != nil {
		return video, err
	}

	return video, nil
}

func (r *videosRepository) CreateVideo(v models.Video) models.Video {
	v.ID = ksuid.New().String()
	v.EditKey = ksuid.New().String()

	r.DB.Create(&v)

	return v
}
