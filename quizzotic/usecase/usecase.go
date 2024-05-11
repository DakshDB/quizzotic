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

func (h *quizzoticUsecase) CreateQuiz(quiz *domain.Quiz) error {
	return h.quizzoticRepo.CreateQuiz(quiz)
}

func (h *quizzoticUsecase) GetQuizzes() ([]domain.Quiz, error) {
	return h.quizzoticRepo.GetQuizzes()
}

func (h *quizzoticUsecase) GetQuizByID(id int) (domain.Quiz, error) {
	return h.quizzoticRepo.GetQuizByID(id)
}

func (h *quizzoticUsecase) UpdateQuiz(id int, quiz *domain.Quiz) error {
	return h.quizzoticRepo.UpdateQuiz(id, quiz)
}
