package services

import (
	"github.com/goodcodeguy/honest-truth-api/repositories"
	"github.com/goodcodeguy/honest-truth-api/repositories/models"
	"github.com/goodcodeguy/honest-truth-api/services/dtos"
	"sync"
)

var videoServiceSingleton sync.Once
var videoServiceInstance *videoService

type videoService struct {
	videoRepository repositories.VideosRepository
}

type VideoService interface {
	GetAllVideos() []dtos.Video
	FindVideoById(id string) (dtos.Video, error)
	CreateVideo(v dtos.Video) dtos.Video
}

func NewVideoService(videoRepository repositories.VideosRepository) VideoService {
	videoServiceSingleton.Do(func() {
		videoServiceInstance = &videoService{
			videoRepository: videoRepository,
		}
	})
	return videoServiceInstance
}

func (s *videoService) GetAllVideos() []dtos.Video {
	videos := s.videoRepository.AllVideos()

	d := make([]dtos.Video, len(videos))
	for vIndex, v := range videos {

		notations := make([]dtos.Notation, len(v.Notations))
		for nIndex, n := range v.Notations {
			notations[nIndex] = dtos.Notation{
				ID:          n.ID,
				VideoID:     n.VideoID,
				Description: n.Description,
				Time:        n.Time,
			}
		}

		d[vIndex] = dtos.Video{
			ID:          v.ID,
			Title:       v.Title,
			UserHash:    v.UserHash,
			Description: v.Description,
			YouTubeID:   v.YouTubeID,
			Notations:   notations,
		}
	}

	return d
}

func (s *videoService) FindVideoById(id string) (dtos.Video, error) {
	video, err := s.videoRepository.FindVideoById(id)
	if err != nil {
		return dtos.Video{}, err
	}

	notations := make([]dtos.Notation, len(video.Notations))
	for nIndex, n := range video.Notations {
		notations[nIndex] = dtos.Notation{
			ID:          n.ID,
			VideoID:     n.VideoID,
			Description: n.Description,
			Time:        n.Time,
		}
	}

	return dtos.Video{
		ID:          video.ID,
		Title:       video.Title,
		UserHash:    video.UserHash,
		Description: video.Description,
		YouTubeID:   video.YouTubeID,
		EditKey:     video.EditKey,
		Notations:   notations,
	}, nil
}

func (s *videoService) CreateVideo(v dtos.Video) dtos.Video {
	video := s.videoRepository.CreateVideo(models.Video{
		Title:       v.Title,
		UserHash:    v.UserHash,
		Description: v.Description,
		YouTubeID:   v.YouTubeID,
	})

	return dtos.Video{
		ID:          video.ID,
		Title:       video.Title,
		UserHash:    video.UserHash,
		Description: video.Description,
		YouTubeID:   video.YouTubeID,
		EditKey:     video.EditKey,
	}
}
