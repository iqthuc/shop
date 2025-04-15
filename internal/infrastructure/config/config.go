package config

// config đại diện cho cấu hình toàn bộ ứng dụng.
type AppConfig struct {
	Logger   *Logger   `mapstructure:"logger"`
	Database *Database `mapstructure:"database"`
	Server   *Server   `mapstructure:"server"`
	// Logger   LoggerConfig
}
