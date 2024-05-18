package config

import (
	"github.com/spf13/viper"
	"strconv"
)

var PORT = "8080"

func InitializeConfig() {
	// Set configuration file which will be used to get/set config values
	viper.SetConfigFile(".env")
	// Ask viper to overwrite any configuration values with their corresponding environment counterparts
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Get PORT from env or use default if not set
	if viper.IsSet("PORT") {
		_port := viper.GetInt("PORT")
		PORT = strconv.Itoa(_port)
	}
}
