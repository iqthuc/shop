package delivery

import (
	"errors"
	"log/slog"
	"shop/internal/features/auth/core"
	"shop/internal/features/auth/core/dto"
	"shop/pkg/utils/errorx"
	"shop/pkg/utils/messages"
	"shop/pkg/utils/response"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	useCase   core.AuthUseCase
	validator validator.Validate
}

func NewHandler(useCase core.AuthUseCase, validator validator.Validate) handler {
	return handler{
		useCase:   useCase,
		validator: validator,
	}
}

func (h handler) Logout(c *fiber.Ctx) error {
	// add token to blacklist later

	c.ClearCookie("refresh_token")
	return response.SuccessJson(c, nil, messages.LogoutSuccess.String())
}

func (h handler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return response.ErrorJson(c, errorx.ErrRefreshTokenNotFound, fiber.StatusUnauthorized)
	}

	result, err := h.useCase.RefreshToken(c.Context(), refreshToken)
	if err != nil {
		return response.ErrorJson(c, errorx.ErrInvalidRefreshToken, fiber.StatusUnauthorized)
	}

	resp := dto.RefreshTokenReponse{
		AccessToken: result,
	}

	return response.SuccessJson(c, resp, messages.RefreshSuccess.String())
}

func (h handler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	if err := h.validator.Struct(req); err != nil {
		return response.ErrorJson(c, errorx.ErrValidationFailed, fiber.StatusBadRequest)
	}

	input := dto.LoginInput(req)

	result, err := h.useCase.Login(c.Context(), input)
	if err != nil {
		slog.Error("login failed", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrSomethingWrong, fiber.StatusInternalServerError)
	}

	resp := dto.LoginResponse{
		UserID:      result.UserID,
		AccessToken: result.AccessToken,
		ExpiresIn:   result.ExpiresIn,
		TokenType:   result.TokenType,
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    result.RefreshToken,
		Expires:  time.Now().Add(result.RefreshExpiresIn),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	slog.Info("user login", slog.String("user", result.UserID.String()))

	return response.SuccessJson(c, resp, messages.LoginSuccess.String())
}

func (h handler) SignUp(c *fiber.Ctx) error {
	var req dto.SignUpRequest

	if err := c.BodyParser(&req); err != nil {
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	if err := h.validator.Struct(req); err != nil {
		return response.ErrorJson(c, errorx.PrettyValidationErrors(err), fiber.StatusBadRequest)
	}

	input := dto.SignUpInput(req)
	err := h.useCase.SignUp(c.Context(), input)
	if err != nil {
		slog.Error("signup failed something", slog.String("details", err.Error()))
		switch {
		case errors.Is(err, errorx.ErrValidationFailed):
			return response.ErrorJson(c, errorx.ErrValidationFailed, fiber.StatusBadRequest)
		case errors.Is(err, errorx.ErrEmailAlready):
			return response.ErrorJson(c, errorx.ErrEmailAlready, fiber.StatusBadRequest)
		default:
			return response.ErrorJson(c, errorx.ErrSomethingWrong, fiber.StatusInternalServerError)
		}
	}

	slog.Info("sign up success")

	return response.SuccessJson(c, nil, messages.SignUpSuccess.String())
}
