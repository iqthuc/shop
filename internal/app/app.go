package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shop/internal/features/auth"
	"shop/internal/infrastructure/config"
	"shop/internal/infrastructure/database"
	"time"
)

// app đại diện cho ứng dụng với tất cả các depencencies
type Application struct {
	// router *http.ServeMux
	// logger *logger.Logger
	server *http.Server
	db     *sql.DB
}

func NewApp(config *config.AppConfig) *Application {
	db, _ := database.NewPostPresDB(config.Database)

	server := &http.Server{
		Addr: config.Server.Address(),
	}
	app := &Application{
		db:     db,
		server: server,
	}
	return app

}

func (a *Application) Run() {
	a.startServer()
	a.registerRoutes()
}

func (a *Application) registerRoutes() {
	mainRouter := http.NewServeMux()
	auth.RegisterRoute(mainRouter, a.db)
}

func (a *Application) startServer() error {
	go func() {
		fmt.Println("start server at", a.server.Addr)
		a.server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shuting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("Server exiting")
	return nil
}
