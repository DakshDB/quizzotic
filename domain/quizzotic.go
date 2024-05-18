package domain

import (
	"github.com/google/uuid"
)

type Choice struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"questionId"`
	Text       string `json:"text"`
}

func (Choice) TableName() string {
	return "choice"
}

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Choices  []Choice `json:"choices" gorm:"foreignKey:QuestionID"`
	Answer   string   `json:"answer"`
	AnswerID int      `json:"answerId" gorm:"-"`
	QuizID   int      `json:"quizId"`
}

func (Question) TableName() string {
	return "question"
}

type Quiz struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description" gorm:"-"`
	Question       []Question `json:"questions" gorm:"foreignKey:QuizID"`
	TotalQuestions int        `json:"totalQuestions"`
	MaxTime        int        `json:"maxTime"`
}

func (Quiz) TableName() string {
	return "quiz"
}

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);primary_key"` // Store UUID as CHAR(36)
    Name     string `json:"name"`
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"password"`
	Verified bool   `json:"verified"`
}

func (User) TableName() string {
	return "user"
}

type QuizzoticUsecase interface {
	HealthCheck() (string, error) // HealthCheck checks the health status of the API server

	CreateQuiz(quiz *Quiz) error
	GetQuizzes() ([]Quiz, error)
	GetQuizByID(id int) (Quiz, error)
	UpdateQuiz(id int, quiz *Quiz) error
	Signup(email string, password string, name string) (string, error)  // Signup handles new user registration
    Login(email string, password string) (User, string, error)  // Login validates user credentials and returns the user
	GenerateJWT(user User) (string, error)
}

type QuizzoticRepository interface {
	CheckDBConnection() (string, error) // CheckDBConnection checks the database connection

	CreateQuiz(quiz *Quiz) error
	GetQuizzes() ([]Quiz, error)
	GetQuizByID(id int) (Quiz, error)
	UpdateQuiz(id int, quiz *Quiz) error
	CreateUser(email string, password string, name string) (User, error)  // Create inserts a new user into the database
    FindUserByEmail(email string) (User, error)  // FindByEmail retrieves a user by their email
    UpdateUser(user User) error  // Update modifies an existing user record
}
