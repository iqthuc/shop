package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"shop/internal/features/auth"
	"shop/internal/features/cart"
	"shop/internal/features/order"
	"shop/internal/features/product"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/server"
	"shop/pkg/token"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	Server     *server.Server
	store      store.Store
	validator  *validator.Validate
	tokenMaker token.TokenMaker
	redis      *redis.Client
}

func NewApp(
	server *server.Server,
	store store.Store,
	validator *validator.Validate,
	tokenMaker token.TokenMaker,
	redis *redis.Client,
) *Application {
	return &Application{
		Server:     server,
		store:      store,
		validator:  validator,
		tokenMaker: tokenMaker,
		redis:      redis,
	}
}

func (a *Application) RegisterRoutes() {
	const (
		maxRequestsPerPeriod = 5
		rateLimitDuration    = 30 * time.Second
	)
	a.Server.Fiber.Use(
		logger.New(),
		limiter.New(limiter.Config{
			Max:        maxRequestsPerPeriod,
			Expiration: rateLimitDuration,
		}),
	)

	auth.SetupModule(a.Server.Fiber, a.store, *a.validator, a.tokenMaker)
	product.SetupModule(a.Server.Fiber, a.store, *a.validator, a.redis)
	cart.SetupModule(a.Server.Fiber, a.store, *a.validator, a.tokenMaker)
	order.SetupModule(a.Server.Fiber, a.store, *a.validator, a.tokenMaker)
}

func (a *Application) run() {
	a.RegisterRoutes()
	a.startServer() // place the end
}

func (a *Application) startServer() {
	go func() {
		a.Server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	a.cleanup()
}

func (a *Application) cleanup() {
	// releases resources before the app exits.
	const serverShutdownTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	err := a.Server.Fiber.ShutdownWithContext(ctx)
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
