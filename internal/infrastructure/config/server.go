package config

import "fmt"

type Server struct {
	Host string
	Port int
}

func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
