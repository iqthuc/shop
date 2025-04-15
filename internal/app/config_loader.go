package app

import (
	"log/slog"
	"os"
	"shop/internal/infrastructure/config"

	"github.com/spf13/viper"
)

func loadConfig(configPath, configName, configType string) *config.AppConfig {
	viper := viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("cannot read file config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	config := &config.AppConfig{}
	if err := viper.Unmarshal(config); err != nil {
		slog.Error("cannot unmarshal config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return config
}
