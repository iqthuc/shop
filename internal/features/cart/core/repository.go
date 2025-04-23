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
	AddToCart(ctx context.Context, item dto.AddToCartRequest) error
}
