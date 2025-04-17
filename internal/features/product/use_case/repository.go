package use_case

import (
	"context"
	"shop/internal/features/product/dto"
	"shop/internal/features/product/entity"
)

type repository interface {
	FetchProducts(ctx context.Context, params dto.GetProductsParams) ([]entity.Product, int, error)
}
