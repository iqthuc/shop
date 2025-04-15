package auth

import (
	"shop/internal/infrastructure/database/store"
	"shop/pkg/token"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, s store.Store, v validator.Validate, tk token.TokenMaker) {
	repo := NewRepository(s)
	useCase := NewUseCase(repo, v, tk)
	handler := NewHandler(useCase)

	auth := r.Group("/auth")
	auth.Get("/sign-up", handler.SignUp)
	auth.Get("/login", handler.Login)
}
