package http

import (
	"github.com/labstack/echo/v4"
	"quizzotic-backend/domain"
	"strconv"
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
	//Quiz endpoints
	e.POST("/quiz", handler.CreateQuiz)
	e.GET("/quiz", handler.GetQuizzes)
	e.GET("/quiz/:id", handler.GetQuizByID)
	e.PUT("/quiz/:id", handler.UpdateQuiz)

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

// CreateQuiz is the handler for creating a new quiz
func (h *quizzoticHandler) CreateQuiz(c echo.Context) error {
	quiz := domain.Quiz{}
	err := c.Bind(&quiz)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	err = h.quizzoticUsecase.CreateQuiz(&quiz)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(201, quiz)
}

// GetQuizzes is the handler for getting all quizzes
func (h *quizzoticHandler) GetQuizzes(c echo.Context) error {
	quizzes, err := h.quizzoticUsecase.GetQuizzes()
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, quizzes)
}

// GetQuizByID is the handler for getting a quiz by ID
func (h *quizzoticHandler) GetQuizByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	quiz, err := h.quizzoticUsecase.GetQuizByID(id)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, quiz)
}

// UpdateQuiz is the handler for updating a quiz
func (h *quizzoticHandler) UpdateQuiz(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	quiz := domain.Quiz{}
	err = c.Bind(&quiz)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	err = h.quizzoticUsecase.UpdateQuiz(id, &quiz)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, quiz)
}
