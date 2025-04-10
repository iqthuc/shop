package config

type Database struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}
