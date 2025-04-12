package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shop/internal/features/auth"
	"shop/internal/infrastructure/database/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

// app đại diện cho ứng dụng với tất cả các depencencies
type application struct {
	// router *http.ServeMux
	// logger *logger.Logger
	server    *http.Server
	store     store.Store
	validator *validator.Validate
}

func (a *application) run() {
	a.registerRoutes()
	a.startServer() // place the end
}

func (a *application) registerRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	a.server.Handler = r

	authModule := auth.NewModule(a.store, a.validator)
	r.Mount("/auth", authModule.InitRoutes())
}

func (a *application) startServer() {
	go func() {
		fmt.Println("Start server at", a.server.Addr)
		err := a.server.ListenAndServe()
		if err != nil {
			log.Fatal("Sever startup failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shuting down server")
	a.cleanup()
}

func (a *application) cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown server failed: %v", err)
	}

	dbCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := a.store.CloseDB(dbCtx); err != nil {
		log.Printf("Close database failed: %v", err)
	}
	log.Println("Server exiting")
}
