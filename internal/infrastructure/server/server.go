package server

import (
	"log/slog"
	"shop/internal/infrastructure/config"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Fiber  *fiber.App
	Config *config.Server
}

func New(cfg *config.Server) *Server {
	f := fiber.New(
		fiber.Config{
			AppName:      "shop",
			ErrorHandler: NewErrorHandler(),
			Prefork:      false,
		},
	)
	server := Server{
		Fiber:  f,
		Config: cfg,
	}

	return &server
}

func (s *Server) ListenAndServe() {
	err := s.Fiber.Listen(s.Config.Address())
	if err != nil {
		slog.Error("cannot start server", slog.String("error", err.Error()))
	}
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
