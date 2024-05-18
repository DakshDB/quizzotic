package usecase

import (
	"errors"
	"quizzotic-backend/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type quizzoticUsecase struct {
	quizzoticRepo domain.QuizzoticRepository
}

func NewQuizzoticUsecase(quizzoticRepo domain.QuizzoticRepository) domain.QuizzoticUsecase {
	return &quizzoticUsecase{
		quizzoticRepo: quizzoticRepo,
	}
}

func (h *quizzoticUsecase) HealthCheck() (string, error) {
	return h.quizzoticRepo.CheckDBConnection()
}

func (h *quizzoticUsecase) CreateQuiz(quiz *domain.Quiz) error {
	return h.quizzoticRepo.CreateQuiz(quiz)
}

func (h *quizzoticUsecase) GetQuizzes() ([]domain.Quiz, error) {
	return h.quizzoticRepo.GetQuizzes()
}

func (h *quizzoticUsecase) GetQuizByID(id int) (domain.Quiz, error) {
	return h.quizzoticRepo.GetQuizByID(id)
}

func (h *quizzoticUsecase) UpdateQuiz(id int, quiz *domain.Quiz) error {
	return h.quizzoticRepo.UpdateQuiz(id, quiz)
}

// Signup handles new user registration and returns a JWT token for immediate login
func (u *quizzoticUsecase) Signup(email string, password string, name string) (string, error) {
    if _, err := u.quizzoticRepo.FindUserByEmail(email); err == nil {
        return "", errors.New("email already in use")
    }
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    password = string(hashedPassword)
	// Declare the user variable outside the if block
	var user domain.User

    // Now use the variable without the :=, so you don't redeclare it
	user, err = u.quizzoticRepo.CreateUser(email, password, name)
	if err != nil {
    	return "", err
	}

	// Now you can use user in the u.GenerateJWT
	token, err := u.GenerateJWT(user)
	if err != nil {
    	return "", err
	}

	return token, nil
}

// Login validates user credentials with bcrypt password verification
func (u *quizzoticUsecase) Login(email string, password string) (domain.User, string, error) {
    user, err := u.quizzoticRepo.FindUserByEmail(email)
    if err != nil {
        return domain.User{}, "", err
    }

    // Verify the password
    if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return domain.User{}, "", errors.New("invalid credentials")
    }

    // Generate JWT token
    token, err := u.GenerateJWT(user)
    if err != nil {
        return domain.User{}, "", err
    }

    return user, token, nil
}

func (u *quizzoticUsecase) GenerateJWT(user domain.User) (string, error) {
    claims := jwt.MapClaims{}
    claims["user_id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires after 72 hours

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte("your_secret_key")) // Replace "your_secret_key" with a secure key
    if err != nil {
        return "", err
    }

    return signedToken, nil
}