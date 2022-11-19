package repositories

import (
	"fmt"
	"github.com/goodcodeguy/honest-truth-api/config"
	"github.com/goodcodeguy/honest-truth-api/repositories/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Repository struct {
	DB *gorm.DB
}

var NotationsRepo NotationsRepository
var VideosRepo VideosRepository

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv(config.ENV_DB_HOST_KEY), os.Getenv(config.ENV_DB_PORT_KEY), os.Getenv(config.ENV_DB_USER_KEY), os.Getenv(config.ENV_DB_PASS_KEY), os.Getenv(config.ENV_DB_NAME_KEY))

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("error: %s", err.Error())
		return
	}

	if os.Getenv(config.ENV_DB_AUTO_MIGRATE_KEY) == "1" {
		log.Print("Database Auto Migrations Enabled")
		_ = database.AutoMigrate(&models.Video{}, &models.Notation{})
	}

	DB := database

	NotationsRepo = NewNotationsRepository(DB)
	VideosRepo = NewVideosRepository(DB)
}
