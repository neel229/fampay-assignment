package utils

import (
	"log"

	"github.com/spf13/viper"
)

// Config contains project's
// configuration variables
type Config struct {
	GoogleAPIKey string `mapstructure:"GOOGLE_API_KEY"`
	Port         string `mapstructure:"PORT"`
	TimeInterval string `mapstructure:"TIME_INTERVAL"`
}

// LoadConfig loads the config variables and returns
// a config struct
func LoadConfig(path string) (config *Config, err error) {
	// set config file's path, name and type
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// load environment variables if they
	// match name
	viper.AutomaticEnv()

	// find and read the config file
	if err = viper.ReadInConfig(); err != nil {
		log.Printf("error reading config file: %v\n", err)
		return
	}
	err = viper.Unmarshal(&config)
	return
}
