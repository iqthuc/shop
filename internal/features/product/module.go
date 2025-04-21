package product

import (
	"shop/internal/features/product/core"
	"shop/internal/features/product/delivery/rest"
	"shop/internal/features/product/repository"
	"shop/internal/infrastructure/database/store"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupModule(r fiber.Router, s store.Store, v validator.Validate) {
	repo := repository.NewProductPostgreRepo(s)
	useCase := core.NewProductUseCase(repo)
	handler := rest.NewHandler(useCase, v)

	api := r.Group("/api")
	api.Get("/products", handler.GetProducts)
	api.Get("/product/:id", handler.GetProductDetail)
}
