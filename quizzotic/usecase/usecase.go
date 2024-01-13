package usecase

import (
	"quizzotic-backend/domain"
)

type quizzoticUsecase struct {
	quizzoticRepo domain.QuizzoticRepository
}

func NewQuizzoticUsecase(quizzoticRepo domain.QuizzoticRepository) domain.QuizzoticUsecase {
	return &quizzoticUsecase{
		quizzoticRepo: quizzoticRepo,
	}
}

func (h *quizzoticUsecase) HealthCheck() (string, error) {
	return h.quizzoticRepo.CheckDBConnection()
}
