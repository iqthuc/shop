package app

import (
	"net/http"
	"shop/internal/infrastructure/config"
)

func newServer(cfg *config.Server) *http.Server {
	return &http.Server{
		Addr: cfg.Address(),
	}
}
