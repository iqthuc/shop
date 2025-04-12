package app

import (
	"fmt"
	"shop/internal/infrastructure/config"

	"github.com/spf13/viper"
)

func loadConfig(configPath, configName, configType string) *config.AppConfig {
	viper := viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Cannot load config file: %w", err))
	}
	config := &config.AppConfig{}
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("Cannot load config file: %w", err))
	}

	return config
}
