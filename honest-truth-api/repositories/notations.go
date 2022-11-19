package repositories

import (
	"github.com/goodcodeguy/honest-truth-api/repositories/models"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type notationsRepository Repository

type NotationsRepository interface {
	AllNotations() []models.Notation
	CreateNotation(d models.Notation) models.Notation
}

func NewNotationsRepository(db *gorm.DB) NotationsRepository {
	return &notationsRepository{
		DB: db,
	}
}

func (r *notationsRepository) AllNotations() []models.Notation {
	var notations []models.Notation
	r.DB.Find(&notations)
	return notations
}

func (r *notationsRepository) CreateNotation(n models.Notation) models.Notation {
	n.ID = ksuid.New().String()

	r.DB.Create(&n)

	return n
}
