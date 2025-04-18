package server

import (
	"errors"
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
		var fe *fiber.Error
		if errors.As(err, &fe) {
			code = fe.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
