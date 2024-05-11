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

	// TODO: Implement the following endpoints
	////Quiz endpoints
	//e.POST("/quiz", handler.CreateQuiz)
	//e.GET("/quiz", handler.GetQuizzes)
	//e.GET("/quiz/:id", handler.GetQuizByID)
	//e.PUT("/quiz/:id", handler.UpdateQuiz)
	//e.DELETE("/quiz/:id", handler.DeleteQuiz)
	//
	//// Question endpoints
	//e.POST("/question", handler.CreateQuestion)
	//e.GET("/question", handler.GetQuestions)
	//e.GET("/question/:id", handler.GetQuestionByID)
	//e.PUT("/question/:id", handler.UpdateQuestion)
	//e.DELETE("/question/:id", handler.DeleteQuestion)
}

// HealthCheck is the handler for health check endpoint
func (h *quizzoticHandler) HealthCheck(c echo.Context) error {

	var response = make(map[string]string)
	response["status"] = "Failure"

	_, err := h.quizzoticUsecase.HealthCheck()
	if err != nil {
		return c.JSON(503, response)
	}

	response["status"] = "Success"
	return c.JSON(200, response)
}
