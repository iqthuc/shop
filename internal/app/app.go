package app

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"shop/internal/features/auth"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/server"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// app đại diện cho ứng dụng với tất cả các depencencies
type application struct {
	// router *http.ServeMux
	// logger *logger.Logger
	server    *server.Server
	store     store.Store
	validator *validator.Validate
}

func (a *application) run() {

	a.registerRoutes()
	a.startServer() // place the end
}

func (a *application) registerRoutes() {
	a.server.Fiber.Use(logger.New())
	auth.RegisterRoutes(a.server.Fiber, a.store, *a.validator)

}

func (a *application) startServer() {
	go func() {
		a.server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.cleanup()
}

func (a *application) cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.server.Fiber.ShutdownWithContext(ctx)
	slog.Info("Server exited")
	dbCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := a.store.CloseDB(dbCtx); err != nil {
		log.Printf("Close database failed: %v", err)
	}
	slog.Info("Database closed")
}
