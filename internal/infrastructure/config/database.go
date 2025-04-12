package config

import (
	"fmt"
	"net/url"
)

type Database struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

func (db *Database) DataSourceName() string {
	u := &url.URL{
		Scheme:   db.Driver,
		User:     url.UserPassword(db.Username, db.Password),
		Host:     fmt.Sprintf("%s:%d", db.Host, db.Port),
		Path:     "/" + db.DBName,
		RawQuery: "sslmode=" + db.SSLMode,
	}
	return u.String()
}
