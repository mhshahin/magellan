package config

import (
	"github.com/mhshahin/magellan/models"
	"github.com/spf13/viper"
)

func LoadConfig() (*models.Config, error) {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")

	config.AddConfigPath(".")
	config.AddConfigPath("./config")

	config.SetEnvPrefix("MAGELLAN")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg models.Config

	if err := config.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
