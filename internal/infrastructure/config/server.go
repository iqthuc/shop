package config

import "fmt"

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
