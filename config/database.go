package config

import (
	"fmt"
	"log"

	"quizzotic-backend/domain"

	_ "github.com/sijms/go-ora/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DNS string

func InitializeMysqlDB() (*gorm.DB, error) {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	DNS = user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(DNS, "This is my server from viper")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate will create the table if it doesn't exist
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("failed to auto-migrate User schema: %v", err)
	}
	return db,nil
}
