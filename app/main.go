package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	db := config.InitializeOracleDBCollection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	quizzoticRepo := _repository.NewOracleDBQuizzoticRepository(db)
	quizzoticUsecase := _usecase.NewQuizzoticUsecase(quizzoticRepo)
	_handler.NewQuizzoticHandler(e, quizzoticUsecase)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}
