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
