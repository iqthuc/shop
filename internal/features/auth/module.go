package auth

import (
	"shop/internal/features/auth/delivery"
	"shop/internal/features/auth/repository"
	"shop/internal/features/auth/use_case"
	"shop/internal/infrastructure/database/store"
	"shop/internal/middleware"
	"shop/pkg/token"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupModule(r fiber.Router, s store.Store, v validator.Validate, tk token.TokenMaker) {
	repo := repository.NewRepository(s)
	useCase := use_case.NewUseCase(repo, tk)
	handler := delivery.NewHandler(useCase, v)

	auth := r.Group("/auth")
	auth.Get("/sign-up", handler.SignUp)
	auth.Get("/login", handler.Login)
	auth.Get("/refresh-token", middleware.JWTAuth(tk), handler.RefreshToken)
	auth.Get("/logout", handler.Logout)
}
