package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shop/internal/config"
	"shop/internal/features/auth"
	"time"
)

// app đại diện cho ứng dụng với tất cả các depencencies
type Application struct {
	config *config.AppConfig
	router *http.ServeMux
	// logger *logger.Logger
	// db *database.DB
}

func (a *Application) Run() {
	a.registerRoutes()
	a.startServer()
}

func (a *Application) registerRoutes() {
	auth.RegisterRoute(a.router)
}

func (a *Application) startServer() error {
	server := http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	go func() {
		fmt.Println("start server at", server.Addr)
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shuting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("Server exiting")
	return nil
}
