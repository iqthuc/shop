package core

import (
	"context"
	"shop/internal/features/cart/core/dto"
)

type CartUseCase interface {
	AddToCart(ctx context.Context, item dto.AddToCartRequest) error
}

type cartUC struct {
	repo CartRepository
}

func NewCartUseCase(repo CartRepository) cartUC {
	return cartUC{
		repo: repo,
	}
}

func (uc cartUC) AddToCart(ctx context.Context, item dto.AddToCartRequest) error {
	err := uc.repo.AddToCart(ctx, item)
	if err != nil {
		return err
	}
	return nil
}
