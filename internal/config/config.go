package config

import (
	"github.com/bluewave-labs/checkmate-cli/pkg/logger"
	"github.com/spf13/viper"
)

var AppConfig Credentials

type Credentials struct {
	APIBaseURL string
	APIKey     string
	UserID     string
	TeamID     string
}

func init() {
	viper.SetConfigName("config")            // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.checkmate/") // call multiple times to add many search paths

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logger.Error(err.Error())
	}

	AppConfig = NewConfig()
}

func NewConfig() Credentials {
	return Credentials{
		APIBaseURL: viper.GetString("base_url"),
		APIKey:     viper.GetString("api_key"),
		UserID:     viper.GetString("user_id"),
		TeamID:     viper.GetString("team_id"),
	}
}
