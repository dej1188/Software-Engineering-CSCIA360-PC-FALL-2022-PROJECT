package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodcodeguy/honest-truth-api/controllers/api"
	"github.com/goodcodeguy/honest-truth-api/services"
	"github.com/goodcodeguy/honest-truth-api/services/dtos"
	"net/http"
)

type notationsController struct {
	notationService services.NotationService
	videoService    services.VideoService
}

type NotationsController interface {
	FindNotations(c *gin.Context)
	CreateNotation(c *gin.Context)
}

func NewNotationsController(notationService services.NotationService, videoService services.VideoService) NotationsController {
	return &notationsController{
		notationService: notationService,
		videoService:    videoService,
	}
}

func (ctl notationsController) FindNotations(c *gin.Context) {
	notations := ctl.notationService.AllNotations()

	d := make([]api.NotationResponse, len(notations))
	for nIndex, n := range notations {
		d[nIndex] = api.NotationResponse{
			ID:          n.ID,
			VideoID:     n.VideoID,
			Description: n.Description,
			Time:        n.Time,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": d})
}

func (ctl notationsController) CreateNotation(c *gin.Context) {
	var d api.CreateNotationRequest
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video, _ := ctl.videoService.FindVideoById(d.VideoID) // TODO: fix error check
	if video.EditKey != d.EditKey {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Edit Key"})
		return
	}

	notation := dtos.Notation{Description: d.Description, VideoID: d.VideoID, Time: d.Time}
	n := ctl.notationService.CreateNotation(notation)

	c.JSON(http.StatusOK, gin.H{"data": api.NotationResponse{
		ID:          n.ID,
		VideoID:     n.VideoID,
		Description: n.Description,
		Time:        n.Time,
	}})
}
