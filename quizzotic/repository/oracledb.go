package repository

import (
	"database/sql"
	"quizzotic-backend/domain"
)

type oracleDBQuizzoticRepository struct {
	*sql.DB
}

// NewOracleDBQuizzoticRepository will create an object that represent the domain.QuizzoticRepository interface
func NewOracleDBQuizzoticRepository(db *sql.DB) domain.QuizzoticRepository {
	return &oracleDBQuizzoticRepository{
		db,
	}
}

// CheckDBConnection check the health status of the API server
func (repo *oracleDBQuizzoticRepository) CheckDBConnection() (status string, err error) {
	err = repo.Ping()
	if err != nil {
		status = "Down"
		return
	}
	status = "Up"
	return status, nil
}
