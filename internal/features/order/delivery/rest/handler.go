package rest

import (
	"log/slog"
	"shop/internal/features/order/core"
	"shop/internal/features/order/core/dto"
	"shop/internal/middleware"
	"shop/pkg/utils/errorx"
	"shop/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type handler struct {
	useCase core.OrderUseCase
}

func NewCartHandler(uc core.OrderUseCase) handler {
	return handler{
		useCase: uc,
	}
}

func (h handler) PlaceOrder(c *fiber.Ctx) error {
	orderRequest := dto.OrderRequest{}
	if err := c.BodyParser(&orderRequest); err != nil {
		slog.Debug("invalid request", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusBadRequest)
	}

	userID, ok := c.Locals(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		slog.Debug("invalid user", slog.String("userID", userID.String()))
		return response.ErrorJson(c, errorx.ErrInvalidRequest, fiber.StatusUnauthorized)
	}

	orderRequest.UserID = userID

	err := h.useCase.PlaceOrder(c.Context(), orderRequest)
	if err != nil {
		slog.Debug("failed to order", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrSomethingWrong, fiber.StatusInternalServerError)
	}

	slog.Info("Order placed successfully", slog.String("userID", userID.String()))

	return response.SuccessJson(c, nil, "Order success")
}
