package services

import (
	"github.com/goodcodeguy/honest-truth-api/repositories"
	"github.com/goodcodeguy/honest-truth-api/repositories/models"
	"github.com/goodcodeguy/honest-truth-api/services/dtos"
	"sync"
)

var notationServiceSingleton sync.Once
var notationServiceInstance *notationService

type notationService struct {
	notationRepository repositories.NotationsRepository
}

type NotationService interface {
	AllNotations() []dtos.Notation
	CreateNotation(notation dtos.Notation) dtos.Notation
}

func NewNotationService(notationRepository repositories.NotationsRepository) NotationService {
	notationServiceSingleton.Do(func() {
		notationServiceInstance = &notationService{
			notationRepository: notationRepository,
		}
	})
	return notationServiceInstance
}

func (s notationService) AllNotations() []dtos.Notation {
	notations := s.notationRepository.AllNotations()

	d := make([]dtos.Notation, len(notations))
	for i, n := range notations {
		d[i] = dtos.Notation{
			ID:          n.ID,
			VideoID:     n.VideoID,
			Description: n.Description,
			Time:        n.Time,
		}
	}

	return d
}

func (s notationService) CreateNotation(n dtos.Notation) dtos.Notation {
	notation := s.notationRepository.CreateNotation(models.Notation{
		VideoID:     n.VideoID,
		Time:        n.Time,
		Description: n.Description,
	})

	return dtos.Notation{
		ID:          notation.ID,
		VideoID:     notation.VideoID,
		Time:        notation.Time,
		Description: notation.Description,
	}
}
