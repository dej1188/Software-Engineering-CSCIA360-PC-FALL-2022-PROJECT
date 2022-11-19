package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goodcodeguy/honest-truth-api/config"
	"github.com/goodcodeguy/honest-truth-api/config/environments"
	"github.com/goodcodeguy/honest-truth-api/controllers"
	"github.com/goodcodeguy/honest-truth-api/repositories"
	"github.com/goodcodeguy/honest-truth-api/services"
	"github.com/joho/godotenv"
	"os"
)

// TODO: Have controllers return DTOs instead of models
func main() {

	_ = godotenv.Load(".env")

	if os.Getenv("ENV") == environments.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	repositories.ConnectDatabase()

	// Initialize Services
	notationService := services.NewNotationService(repositories.NotationsRepo)
	videoService := services.NewVideoService(repositories.VideosRepo)

	// Initialize Controllers
	notationsController := controllers.NewNotationsController(notationService, videoService)
	videosController := controllers.NewVideosController(videoService)

	// Routes
	r.GET("/videos", videosController.FindVideos)
	r.GET("/videos/:id", videosController.FindVideoById)
	r.POST("/videos", videosController.CreateVideo)

	r.GET("/notations", notationsController.FindNotations)
	r.POST("/notations", notationsController.CreateNotation)

	r.GET("/health", controllers.HealthCheck)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "", gin.H{})
	})

	listenerAddress := fmt.Sprintf(":%s", os.Getenv(config.ENV_LISTENER_PORT))
	_ = r.Run(listenerAddress)
}
