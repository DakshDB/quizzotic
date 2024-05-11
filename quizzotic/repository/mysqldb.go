package repository

import (
	"gorm.io/gorm"
	"quizzotic-backend/domain"
)

type mysqlDBQuizzoticRepository struct {
	*gorm.DB
}

// NewMysqlDBQuizzoticRepository will create an object that represent the domain.QuizzoticRepository interface
func NewMysqlDBQuizzoticRepository(db *gorm.DB) domain.QuizzoticRepository {
	return &mysqlDBQuizzoticRepository{
		db,
	}
}

// CheckDBConnection check the health status of the API server
func (repo *mysqlDBQuizzoticRepository) CheckDBConnection() (status string, err error) {
	status = repo.DB.Name()

	return status, nil
}

// CreateQuiz will create a new quiz
func (repo *mysqlDBQuizzoticRepository) CreateQuiz(quiz *domain.Quiz) error {
	return repo.DB.Create(quiz).Error
}

// GetQuizzes will return all quizzes
func (repo *mysqlDBQuizzoticRepository) GetQuizzes() (quizzes []domain.Quiz, err error) {
	err = repo.DB.
		Preload("Question.Choices").
		Preload("Question").
		Find(&quizzes).Error
	return
}

// GetQuizByID will return a quiz by its ID
func (repo *mysqlDBQuizzoticRepository) GetQuizByID(id int) (quiz domain.Quiz, err error) {
	err = repo.DB.
		Preload("Question.Choices").
		Preload("Question").
		First(&quiz, id).Error
	return
}

// UpdateQuiz will update a quiz by its ID
func (repo *mysqlDBQuizzoticRepository) UpdateQuiz(id int, quiz *domain.Quiz) error {
	return repo.DB.Save(quiz).Error
}
