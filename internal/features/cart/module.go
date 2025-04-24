package cart

import (
	"shop/internal/features/cart/core"
	"shop/internal/features/cart/delivery/rest"
	"shop/internal/features/cart/repository"
	"shop/internal/infrastructure/database/store"
	"shop/internal/middleware"
	"shop/pkg/token"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupModule(r fiber.Router, s store.Store, v validator.Validate, tk token.TokenMaker) {
	repo := repository.NewCartPostgresRepo(s)
	useCase := core.NewCartUseCase(repo)
	handler := rest.NewCartHandler(useCase, v)

	api := r.Group("/api")
	api.Use(middleware.JWTAuth(tk))
	api.Post("/carts", handler.AddToCart)
	api.Patch("/cart/items/:product_variant_id", handler.UpdateCart)
	api.Delete("/cart/items/:product_variant_id", handler.DeleteCart)
}
