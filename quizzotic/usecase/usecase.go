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
	quizzes, err := h.quizzoticRepo.GetQuizzes()
	if err != nil {
		return nil, err
	}

	// Update the answerID
	for i := range quizzes {
		for j := range quizzes[i].Question {
			for k := range quizzes[i].Question[j].Choices {
				if quizzes[i].Question[j].Choices[k].Text == quizzes[i].Question[j].Answer {
					quizzes[i].Question[j].AnswerID = quizzes[i].Question[j].Choices[k].ID
				}
			}
		}
	}

	return quizzes, nil
}

func (h *quizzoticUsecase) GetQuizByID(id int) (domain.Quiz, error) {
	quiz, err := h.quizzoticRepo.GetQuizByID(id)
	if err != nil {
		return domain.Quiz{}, err
	}

	// Update the answerID
	for i := range quiz.Question {
		for j := range quiz.Question[i].Choices {
			if quiz.Question[i].Choices[j].Text == quiz.Question[i].Answer {
				quiz.Question[i].AnswerID = quiz.Question[i].Choices[j].ID
			}
		}
	}

	return quiz, nil
}

func (h *quizzoticUsecase) UpdateQuiz(id int, quiz *domain.Quiz) error {
	return h.quizzoticRepo.UpdateQuiz(id, quiz)
}
