package rest

import (
	"errors"
	"log/slog"
	"shop/internal/features/cart/core"
	"shop/internal/features/cart/core/dto"
	"shop/internal/middleware"
	"shop/pkg/utils/errorx"
	"shop/pkg/utils/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	userID, ok := c.Locals(middleware.UserIDKey).(uuid.UUID)
	if !ok || userID != req.UserID {
		slog.Debug("invalid user", slog.String("userID", userID.String()))
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
		slog.Int("product variant", req.VariantID),
		slog.Int("quantity", req.Quantity),
	)

	return response.SuccessJson(c, nil, "Added to cart")
}

func (h handler) UpdateCart(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("product_variant_id")
	if err != nil {
		slog.Debug("invalid product ID", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	var req dto.UpdateCartRequest
	req.VariantID = productID

	err = c.BodyParser(&req)
	if err != nil {
		slog.Debug("invalid request", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	userID, ok := c.Locals(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		slog.Debug("invalid user", slog.String("userID", userID.String()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusUnauthorized)
	}

	req.UserID = userID
	slog.Info("check req", slog.Any("req", req))
	err = h.useCase.UpdateCart(c.Context(), req)
	if err != nil {
		if errors.Is(err, core.ErrVariantOrStockInvalid) {
			return response.ErrorJson(c, core.ErrVariantOrStockInvalid, fiber.StatusBadRequest)
		}

		slog.Debug("failed to update cart", slog.String("error", err.Error()))

		return response.ErrorJson(c, errorx.ErrSomethingWrong, fiber.StatusInternalServerError)
	}

	slog.Info("user updated cart",
		slog.String("user", req.UserID.String()),
		slog.Int("product variant", req.VariantID),
		slog.Int("new quantity", req.Quantity),
	)

	return response.SuccessJson(c, nil, "Updated cart")
}

func (h handler) DeleteCart(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("product_variant_id")
	if err != nil {
		slog.Debug("invalid product ID", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	userID, ok := c.Locals(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		slog.Debug("invalid user", slog.String("userID", userID.String()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusUnauthorized)
	}

	req := dto.DeleteCartItemRequest{
		UserID:    userID,
		VariantID: productID,
	}
	err = h.useCase.DeleteCartItem(c.Context(), req)
	if err != nil {
		if errors.Is(err, core.ErrVariantOrStockInvalid) {
			return response.ErrorJson(c, core.ErrVariantOrStockInvalid, fiber.StatusBadRequest)
		}

		slog.Debug("failed to delete cart", slog.String("error", err.Error()))

		return response.ErrorJson(c, errorx.ErrSomethingWrong, fiber.StatusInternalServerError)
	}

	slog.Info("user deleted cart",
		slog.String("user", userID.String()),
		slog.Int("product variant", productID),
	)

	return response.SuccessJson(c, nil, "Deleted cart")
}
