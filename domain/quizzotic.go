package domain

type Choice struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	Text       string `json:"text"`
}

func (Choice) TableName() string {
	return "choice"
}

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"Question"`
	Choices  []Choice `json:"choices" gorm:"foreignKey:QuestionID"`
	Answer   string   `json:"answer"`
	QuizID   int      `json:"quiz_id"`
}

func (Question) TableName() string {
	return "question"
}

type Quiz struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description" gorm:"-"`
	Question       []Question `json:"questions" gorm:"foreignKey:QuizID"`
	TotalQuestions int        `json:"total_questions"`
	MaxTime        int        `json:"max_time"`
}

func (Quiz) TableName() string {
	return "quiz"
}

type QuizzoticUsecase interface {
	HealthCheck() (string, error) // HealthCheck checks the health status of the API server

	CreateQuiz(quiz *Quiz) error
	GetQuizzes() ([]Quiz, error)
	GetQuizByID(id int) (Quiz, error)
	UpdateQuiz(id int, quiz *Quiz) error
}

type QuizzoticRepository interface {
	CheckDBConnection() (string, error) // CheckDBConnection checks the database connection

	CreateQuiz(quiz *Quiz) error
	GetQuizzes() ([]Quiz, error)
	GetQuizByID(id int) (Quiz, error)
	UpdateQuiz(id int, quiz *Quiz) error
}
