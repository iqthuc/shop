package app

import (
	"log/slog"
	"os"
	"shop/internal/infrastructure/config"
)

func Bootstrap() {
	var appConfig *config.AppConfig

	env := os.Getenv("ENV")
	if env == "production" {
		appConfig = loadConfig("configs", "config", "yaml")
	}
	appConfig = loadConfig("configs", "config.dev", "yaml")

	_ = NewLogger(appConfig.Logger)
	slog.Info("Application running...", slog.String("env", env))

	srv := newServer(appConfig.Server)
	store := newStore(appConfig.Database)
	validator := newValidator()
	app := &application{
		server:    srv,
		store:     store,
		validator: validator,
	}
	app.run()
}
