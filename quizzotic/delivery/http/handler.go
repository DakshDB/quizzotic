package http

import (
	"net/http"
	"quizzotic-backend/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type quizzoticHandler struct {
	quizzoticUsecase domain.QuizzoticUsecase
}

func NewQuizzoticHandler(e *echo.Echo, quizzoticUsecase domain.QuizzoticUsecase) {
	handler := &quizzoticHandler{
		quizzoticUsecase: quizzoticUsecase,
	}

	// Health check endpoint
	e.GET("/healthCheck", handler.HealthCheck, JWTMiddleware)

	// TODO: Implement the following endpoints
	//Quiz endpoints
	e.POST("/quiz", handler.CreateQuiz, JWTMiddleware)
	e.GET("/quiz", handler.GetQuizzes, JWTMiddleware)
	e.GET("/quiz/:id", handler.GetQuizByID, JWTMiddleware)
	e.PUT("/quiz/:id", handler.UpdateQuiz, JWTMiddleware)

	//User apis
	e.POST("/signup", handler.Signup)
    e.POST("/login", handler.Login)

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

// Signup is the handler for user registration
func (h *quizzoticHandler) Signup(c echo.Context) error {
	var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
		Name	 string `json:"name"`
    }

    if err := c.Bind(&credentials); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    token, err := h.quizzoticUsecase.Signup(credentials.Email, credentials.Password, credentials.Name)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, map[string]string{"token": token})
}

// Login is the handler for user login
func (h *quizzoticHandler) Login(c echo.Context) error {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.Bind(&credentials); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    user, token, err := h.quizzoticUsecase.Login(credentials.Email, credentials.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
    }

    return c.JSON(http.StatusOK, map[string]interface{}{"token": token, "user": user})
}

