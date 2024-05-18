package main

import (
	"gorm.io/gorm/logger"
	"log"
	"quizzotic-backend/config"
	"quizzotic-backend/domain"
	_handler "quizzotic-backend/quizzotic/delivery/http"
	_repository "quizzotic-backend/quizzotic/repository"
	_usecase "quizzotic-backend/quizzotic/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	e *echo.Echo
)

func init() {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	config.InitializeConfig()
}

func main() {

	config.InitializeMysqlDB()
	db, err := gorm.Open(mysql.Open(config.DNS), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate will create the table if it doesn't exist
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("failed to auto-migrate User schema: %v", err)
	}

	quizzoticRepo := _repository.NewMysqlDBQuizzoticRepository(db)
	quizzoticUsecase := _usecase.NewQuizzoticUsecase(quizzoticRepo)
	_handler.NewQuizzoticHandler(e, quizzoticUsecase)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}
