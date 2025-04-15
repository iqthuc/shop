package auth

import (
	"context"
	"errors"
	"log/slog"
	errs "shop/pkg/utils/errors"
	"shop/pkg/utils/messages"
	"shop/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

type UseCase interface {
	SignUp(ctx context.Context, input signUpInput) (*signUpResult, error)
	Login(ctx context.Context, input loginRequest) (*loginResponse, error)
}

type handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h handler) Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorJson(c, errs.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	result, err := h.useCase.Login(c.Context(), req)
	if err != nil {
		slog.Error("login failed", slog.String("error", err.Error()))
		return response.ErrorJson(c, errs.ErrSomethingWrong, fiber.StatusInternalServerError)
	}

	slog.Info("user login", slog.String("user", result.UserID.String()))

	return response.SuccessJson(c, result, messages.LoginSuccess)
}

func (h handler) SignUp(c *fiber.Ctx) error {
	var req signUpRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorJson(c, errs.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	input := signUpInput(req)
	result, err := h.useCase.SignUp(c.Context(), input)
	if err != nil {
		slog.Error("signup failed something", slog.String("details", err.Error()))
		switch {
		case errors.Is(err, errs.ErrVaidationFailed):
			return response.ErrorJson(c, errs.ErrVaidationFailed, fiber.StatusBadRequest)
		case errors.Is(err, errs.ErrEmailAlready):
			return response.ErrorJson(c, errs.ErrEmailAlready, fiber.StatusBadRequest)
		default:
			return response.ErrorJson(c, errs.ErrSomethingWrong, fiber.StatusInternalServerError)
		}
	}

	resp := signUpResponse{
		Email:     result.email,
		CreatedAt: result.createdAt,
	}

	slog.Info("sign up success")

	return response.SuccessJson(c, resp, messages.SignUpSuccess)
}
