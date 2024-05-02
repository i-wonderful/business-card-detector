package history

import "card_detector/internal/model"

type CardRepo interface {
	GetAll() []model.Card
}

type Service struct {
	cardRepo CardRepo
}

func NewService(personRepo CardRepo) *Service {
	return &Service{
		cardRepo: personRepo,
	}
}

func (s *Service) GetAll() []model.Card {
	return s.cardRepo.GetAll()
}
