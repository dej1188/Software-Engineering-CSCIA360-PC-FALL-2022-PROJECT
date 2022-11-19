package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodcodeguy/honest-truth-api/controllers/api"
	"github.com/goodcodeguy/honest-truth-api/services"
	"github.com/goodcodeguy/honest-truth-api/services/dtos"
	"net/http"
)

type videosController struct {
	videosService services.VideoService
}
type VideosController interface {
	FindVideos(c *gin.Context)
	FindVideoById(c *gin.Context)
	CreateVideo(c *gin.Context)
}

func NewVideosController(videosService services.VideoService) VideosController {
	return &videosController{
		videosService: videosService,
	}
}

func (ctl videosController) FindVideos(c *gin.Context) {
	videos := ctl.videosService.GetAllVideos()

	d := make([]api.VideoResponse, len(videos))
	for i, v := range videos {

		notations := make([]api.NotationResponse, len(v.Notations))
		for nIndex, n := range v.Notations {
			notations[nIndex] = api.NotationResponse{
				ID:          n.ID,
				VideoID:     n.VideoID,
				Description: n.Description,
				Time:        n.Time,
			}
		}

		d[i] = api.VideoResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			UserHash:    v.UserHash,
			YouTubeID:   v.YouTubeID,
			Notations:   notations,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": d})
}

func (ctl videosController) FindVideoById(c *gin.Context) {
	video, err := ctl.videosService.FindVideoById(c.Param("id"))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"data": nil})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	notations := make([]api.NotationResponse, len(video.Notations))
	for nIndex, n := range video.Notations {
		notations[nIndex] = api.NotationResponse{
			ID:          n.ID,
			VideoID:     n.VideoID,
			Description: n.Description,
			Time:        n.Time,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": api.VideoResponse{
		ID:          video.ID,
		Title:       video.Title,
		Description: video.Description,
		UserHash:    video.UserHash,
		YouTubeID:   video.YouTubeID,
		Notations:   notations,
	}})
}

func (ctl videosController) CreateVideo(c *gin.Context) {
	var d api.CreateVideoRequest
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video := dtos.Video{
		Title:       d.Title,
		Description: d.Description,
		UserHash:    d.UserHash,
		YouTubeID:   d.YouTubeID,
	}

	v := ctl.videosService.CreateVideo(video)

	notations := make([]api.NotationResponse, len(v.Notations))
	for _, n := range v.Notations {
		notations = append(notations, api.NotationResponse{
			ID:          n.ID,
			VideoID:     n.VideoID,
			Description: n.Description,
			Time:        n.Time,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": api.CreateVideoResponse{
		ID:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		UserHash:    v.UserHash,
		YouTubeID:   v.YouTubeID,
		EditKey:     v.EditKey,
		Notations:   notations,
	}})
}
