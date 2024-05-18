package repository

import (
	"quizzotic-backend/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
func (repo *mysqlDBQuizzoticRepository) UpdateQuiz(_ int, quiz *domain.Quiz) error {
	return repo.DB.Save(quiz).Error
}

// Create inserts a new user into the database
func (repo *mysqlDBQuizzoticRepository) CreateUser(email string, password string, name string) (domain.User, error) {
	newUser := domain.User{
		ID:       uuid.New(), // Generate a new UUID
		Name:     name,
		Email:    email,
		Password: password, // Use the hashed password
		Verified: false,    // Default to false unless you have an email verification system in place
	}
	// Use GORM's Create method to add the new user to the database
	result := repo.DB.Create(&newUser)
	if result.Error != nil {
		return domain.User{}, result.Error // Return an empty User object and the error
	}

	return newUser, nil // Return the newly created user and nil error
}

// FindByEmail retrieves a user by their email
func (repo *mysqlDBQuizzoticRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	result := repo.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

// Update modifies an existing user record
func (repo *mysqlDBQuizzoticRepository) UpdateUser(user domain.User) error {
	result := repo.DB.Save(&user)
	return result.Error
}
