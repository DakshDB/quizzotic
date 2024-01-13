package domain

type QuizzoticUsecase interface {
	HealthCheck() (string, error) // HealthCheck checks the health status of the API server
}

type QuizzoticRepository interface {
	CheckDBConnection() (string, error) // CheckDBConnection checks the database connection
}
