package core

import (
	"context"
	"shop/internal/features/order/core/dto"
)

type OrderUseCase interface {
	PlaceOrder(ctx context.Context, req dto.OrderRequest) error
}

type orderUC struct {
	repo OrderRepository
}

func NewOrderUseCase(repo OrderRepository) orderUC {
	return orderUC{repo: repo}
}

func (u orderUC) PlaceOrder(ctx context.Context, req dto.OrderRequest) error {
	return u.repo.CreateOrder(ctx, req)
}
