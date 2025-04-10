package config

import "fmt"

type Database struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func (p *Database) DataSourceName() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.Username, p.Password, p.DBName, p.SSLMode)

}
