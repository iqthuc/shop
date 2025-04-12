package app

import (
	"log/slog"
	"shop/internal/infrastructure/config"
	"shop/internal/infrastructure/logger"
)

func NewLogger(cfg *config.Logger) *slog.Logger {
	return logger.ConfigureLogger(cfg)
}
