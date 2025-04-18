package use_case

import (
	"context"
	"shop/internal/features/product/dto"
)

type repository interface {
	FetchProducts(ctx context.Context, params dto.GetProductsParams) ([]dto.Product, int, error)
	GetProductByID(ctx context.Context, productID int) (*dto.ProductDetail, error)
	FetchProductVariantByID(ctx context.Context, productID int) ([]dto.ProductVariant, error)
}
