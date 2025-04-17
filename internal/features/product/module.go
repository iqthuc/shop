package product

import (
	"shop/internal/features/product/delivery"
	"shop/internal/features/product/repository"
	"shop/internal/features/product/use_case"
	"shop/internal/infrastructure/database/store"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupModule(r fiber.Router, s store.Store, v validator.Validate) {
	repo := repository.NewRepository(s)
	useCase := use_case.NewUseCase(repo)
	handler := delivery.NewHandler(useCase, v)

	api := r.Group("/api")
	api.Get("/products", handler.GetProducts)
}
