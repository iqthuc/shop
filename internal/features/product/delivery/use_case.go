package delivery

import (
	"context"
	"shop/internal/features/product/dto"
	"shop/internal/features/product/entity"
)

type UseCase interface {
	GetProducts(ctx context.Context, input dto.GetProductsRequest) (*dto.GetProductsResult[entity.Product], error)
}
