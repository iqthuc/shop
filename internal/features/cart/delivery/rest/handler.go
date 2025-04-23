package rest

import (
	"errors"
	"log/slog"
	"shop/internal/features/cart/core"
	"shop/internal/features/cart/core/dto"
	"shop/internal/middleware"
	"shop/pkg/token"
	"shop/pkg/utils/errorx"
	"shop/pkg/utils/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	useCase   core.CartUseCase
	validator validator.Validate
}

func NewCartHandler(useCase core.CartUseCase, validator validator.Validate) handler {
	return handler{
		useCase:   useCase,
		validator: validator,
	}
}

func (h handler) AddToCart(c *fiber.Ctx) error {
	var req dto.AddToCartRequest
	err := c.BodyParser(&req)
	if err != nil {
		slog.Debug("invalid request", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	claims, ok := c.Locals(middleware.AuthorizationPayloadKey).(*token.TokenClaims)
	if !ok || claims.UserID != req.UserID.String() {
		slog.Debug("invalid user", slog.String("userID", claims.UserID))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusUnauthorized)
	}

	err = h.useCase.AddToCart(c.Context(), req)
	if err != nil {
		if errors.Is(err, core.ErrVariantOrStockInvalid) {
			return response.ErrorJson(c, core.ErrVariantOrStockInvalid, fiber.StatusBadRequest)
		}

		slog.Debug("failed to add to cart", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrSomethingWrong, fiber.StatusInternalServerError)
	}

	slog.Info("user added product to cart",
		slog.String("user", req.UserID.String()),
		slog.Int("product variant", int(req.VariantID)),
		slog.Int("quantity", int(req.Quantity)),
	)

	return response.SuccessJson(c, nil, "Added to cart")
}
