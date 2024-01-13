package http

import (
	"github.com/labstack/echo/v4"
	"quizzotic-backend/domain"
)

type quizzoticHandler struct {
	quizzoticUsecase domain.QuizzoticUsecase
}

func NewQuizzoticHandler(e *echo.Echo, quizzoticUsecase domain.QuizzoticUsecase) {
	handler := &quizzoticHandler{
		quizzoticUsecase: quizzoticUsecase,
	}

	// Health check endpoint
	e.GET("/healthCheck", handler.HealthCheck)
}

// HealthCheck is the handler for health check endpoint
func (h *quizzoticHandler) HealthCheck(c echo.Context) error {

	var response = make(map[string]string)

	healthCheckStatus, err := h.quizzoticUsecase.HealthCheck()
	var status = "Failure"
	if healthCheckStatus == "Up" {
		status = "Success"
	}

	response["status"] = status

	if err != nil {
		return c.JSON(503, response)
	}
	return c.JSON(200, response)
}
