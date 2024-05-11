package config

import (
	"github.com/spf13/viper"
	"strconv"
)

var PORT = "8080"

func InitializeConfig() {
	viper.AutomaticEnv()

	// Get PORT from env or use default if not set
	if viper.IsSet("PORT") {
		_port := viper.GetInt("PORT")
		PORT = strconv.Itoa(_port)
	}
}
