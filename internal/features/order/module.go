package order

import (
	"shop/internal/features/order/core"
	"shop/internal/features/order/delivery/rest"
	"shop/internal/features/order/repository"
	"shop/internal/infrastructure/database/store"
	"shop/internal/middleware"
	"shop/pkg/token"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupModule(r fiber.Router, s store.Store, v validator.Validate, tk token.TokenMaker) {
	repo := repository.NewCartPostgresRepo(s)
	useCase := core.NewOrderUseCase(repo)
	handler := rest.NewCartHandler(useCase)

	api := r.Group("/api")
	api.Use(middleware.JWTAuth(tk))
	api.Post("/order", handler.PlaceOrder)
}
