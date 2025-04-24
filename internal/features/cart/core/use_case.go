package core

import (
	"context"
	"shop/internal/features/cart/core/dto"
)

type CartUseCase interface {
	AddToCart(ctx context.Context, item dto.AddToCartRequest) error
	UpdateCart(ctx context.Context, input dto.UpdateCartRequest) error
	DeleteCartItem(ctx context.Context, input dto.DeleteCartItemRequest) error
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
	err := uc.repo.SaveCartItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (uc cartUC) UpdateCart(ctx context.Context, input dto.UpdateCartRequest) error {
	err := uc.repo.UpdateCart(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (uc cartUC) DeleteCartItem(ctx context.Context, input dto.DeleteCartItemRequest) error {
	err := uc.repo.DeleteCartItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
