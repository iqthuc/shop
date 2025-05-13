package product

import (
	"shop/internal/features/product/core"
	"shop/internal/features/product/delivery/rest"
	"shop/internal/features/product/repository"
	"shop/internal/infrastructure/database/store"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func SetupModule(r fiber.Router, s store.Store, v validator.Validate, redis *redis.Client) {
	repo := repository.NewProductPostgreRepo(s)
	useCase := core.NewProductUseCase(repo, redis)
	handler := rest.NewHandler(useCase, v)

	api := r.Group("/api")
	api.Get("/products", handler.GetProducts)
	api.Get("/product/:id", handler.GetProductDetail)
}
