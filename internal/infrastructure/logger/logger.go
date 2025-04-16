package logger

import (
	"log/slog"
	"os"
	"shop/internal/infrastructure/config"
	"time"
)

func ConfigureLogger(cfg *config.Logger) *slog.Logger {
	var level slog.Level
	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	prodOpts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   slog.TimeKey,
					Value: slog.StringValue(a.Value.Time().Format(time.RFC3339)),
				}
			}
			if a.Key == "password" || a.Key == "token" {
				return slog.Attr{Key: a.Key, Value: slog.StringValue("[REDACTED]")}
			}

			return a
		},
		AddSource: level > slog.LevelInfo,
	}

	var handler slog.Handler
	if cfg.Environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout, prodOpts)
	} else {
		handler = NewPrettyHandler(&slog.HandlerOptions{
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
			AddSource:   true,
		},
			true, // disable AddSource for Info log?
		)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Info("Using log format", slog.String("format", cfg.Format))
	slog.Info("Logger initialized successfully", slog.String("log_level", cfg.Level))
	slog.Debug("This is a debug message (will only show if log level is Debug)")

	return logger
}
