package auth

import (
	"shop/internal/infrastructure/database/store"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, s store.Store, v validator.Validate) {
	repo := NewRepository(s)
	useCase := NewUsecase(repo, v)
	handler := NewHandler(useCase)

	auth := r.Group("/auth")
	auth.Get("/sign-up", handler.SignUp)
}
