package main

import (
	"log"
	"quizzotic-backend/config"
	_handler "quizzotic-backend/quizzotic/delivery/http"
	_repository "quizzotic-backend/quizzotic/repository"
	_usecase "quizzotic-backend/quizzotic/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	e *echo.Echo
)

func init() {
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize Viper
    viper.AutomaticEnv() // Read environment variables
    viper.SetConfigFile(".env") // Specify the .env file to Viper

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file, %s", err)
    }
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
