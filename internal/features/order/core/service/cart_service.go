package service

import (
	"context"

	"github.com/shopspring/decimal"
)

type CartItem struct {
	VariantID string
	Quantity  int
	Price     decimal.Decimal
}

type CartService struct {
	GetCartItems func(ctx context.Context, userID string) ([]*CartItem, error)
}
