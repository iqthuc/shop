package app

import (
	"log/slog"
	"os"
	"shop/internal/infrastructure/config"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/logger"
	"shop/internal/infrastructure/server"
	"shop/pkg/token"

	"github.com/go-playground/validator/v10"
)

func Bootstrap() {
	var appConfig *config.AppConfig
	env := os.Getenv("ENV")
	if env == "production" {
		appConfig = loadConfig("configs", "config", "yaml")
	} else {
		appConfig = loadConfig("configs", "config.dev", "yaml")
	}

	logger.ConfigureLogger(appConfig.Logger)
	slog.Info("Application running...", slog.String("env", env))

	server := server.New(appConfig.Server)
	store := store.NewPostgresStore(appConfig.Database)
	validator := validator.New()
	tkMaker := token.NewJwtMaker(appConfig.Token.SecretKey)

	app := NewApp(server, store, validator, tkMaker)
	app.run()
}
