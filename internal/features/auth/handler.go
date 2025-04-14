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
}

type handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h handler) SignUp(c *fiber.Ctx) error {
	var req signUpRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorJson(c, errs.InvalidRequest, fiber.StatusBadRequest)
	}
	input := signUpInput{
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := h.useCase.SignUp(c.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, errs.VaidationFailed):
			return response.ErrorJson(c, err, fiber.StatusBadRequest)

		case errors.Is(err, errs.EmailAlready):
			return response.ErrorJson(c, err, fiber.StatusBadRequest)

		default:
			slog.Error("signup failed something", slog.String("details", err.Error()))
			return response.ErrorJson(c, errs.SomethingWrong, fiber.StatusInternalServerError)
		}
	}

	resp := signUpResponse{
		Email:     result.email,
		CreatedAt: result.createdAt,
	}
	return response.SuccessJson(c, resp, messages.SignUpSuccess)
}
