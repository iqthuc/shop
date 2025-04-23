package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"shop/internal/features/auth"
	"shop/internal/features/cart"
	"shop/internal/features/product"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/server"
	"shop/pkg/token"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type application struct {
	server     *server.Server
	store      store.Store
	validator  *validator.Validate
	tokenMaker token.TokenMaker
}

func NewApp(
	server *server.Server,
	store store.Store,
	validator *validator.Validate,
	tokenMaker token.TokenMaker,
) *application {
	return &application{
		server:     server,
		store:      store,
		validator:  validator,
		tokenMaker: tokenMaker,
	}
}

func (a *application) run() {
	a.registerRoutes()
	a.startServer() // place the end
}

func (a *application) registerRoutes() {
	a.server.Fiber.Use(logger.New())
	auth.SetupModule(a.server.Fiber, a.store, *a.validator, a.tokenMaker)
	product.SetupModule(a.server.Fiber, a.store, *a.validator)
	cart.SetupModule(a.server.Fiber, a.store, *a.validator, a.tokenMaker)
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
	// releases resources before the app exits.
	const serverShutdownTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	err := a.server.Fiber.ShutdownWithContext(ctx)
	if err != nil {
		slog.Warn("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server exited")

	const dbShutdownTimeout = 2 * time.Second
	dbCtx, cancel := context.WithTimeout(context.Background(), dbShutdownTimeout)
	defer cancel()

	a.store.CloseDB(dbCtx)
	slog.Info("Database closed")
}
