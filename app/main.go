package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"quizzotic-backend/config"
	_handler "quizzotic-backend/quizzotic/delivery/http"
	_repository "quizzotic-backend/quizzotic/repository"
	_usecase "quizzotic-backend/quizzotic/usecase"
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
	db, err := gorm.Open(mysql.Open(config.DNS), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	quizzoticRepo := _repository.NewMysqlDBQuizzoticRepository(db)
	quizzoticUsecase := _usecase.NewQuizzoticUsecase(quizzoticRepo)
	_handler.NewQuizzoticHandler(e, quizzoticUsecase)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}
