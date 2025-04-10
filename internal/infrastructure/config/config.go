package config

// config đại diện cho cấu hình toàn bộ ứng dụng
type AppConfig struct {
	Server   Server
	Database Database
	// Logger   LoggerConfig
}

func Load() (*AppConfig, error) {
	return &AppConfig{}, nil
}
