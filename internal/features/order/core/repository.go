package core

import (
	"context"
	"errors"
	"shop/internal/features/order/core/dto"
)

var (
	ErrVariantInvalid error = errors.New("variant invalid")
	ErrStockInvalid   error = errors.New("some products are out of stock")
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req dto.OrderRequest) error
	// VerifyOrder(ctx context.Context, orderId []int) error
	// DeleteItemsFromCart(ctx context.Context, cartId int) error
}
