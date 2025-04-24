package core

import (
	"context"
	"errors"
	"shop/internal/features/cart/core/dto"
)

var (
	ErrVariantOrStockInvalid error = errors.New("variant or stock invalid")
)

type CartRepository interface {
	SaveCartItem(ctx context.Context, item dto.AddToCartRequest) error
	UpdateCart(ctx context.Context, input dto.UpdateCartRequest) error
	DeleteCartItem(ctx context.Context, input dto.DeleteCartItemRequest) error
}
