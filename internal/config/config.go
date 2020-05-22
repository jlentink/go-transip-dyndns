package config

import (
	"github.com/jlentink/go-transip-dyndns/internal/logger"
	"github.com/spf13/viper"
)

var _config *viper.Viper

// Init the config setup
func Init() {
	_config = viper.New()
	_config.SetConfigType("toml")
	_config.AddConfigPath("/etc")
	_config.AddConfigPath(".")
	_config.SetConfigName("go-transip-dyndns")

	if err := _config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Get().Fatalf("%s\n", err.Error())
		} else {
			logger.Get().Fatalf("Error parsing config file. (%s)\n", err.Error())
		}
	}
}

// Get the config instance
func Get() *viper.Viper {
	return _config
}
