package core

import (
	"context"
	"shop/internal/features/product/core/dto"
	"shop/internal/features/product/core/entity"
)

type ProductRepository interface {
	FetchProducts(ctx context.Context, params dto.GetProductsParams) ([]entity.Product, int, error)
	GetProductByID(ctx context.Context, productID int) (*entity.ProductDetail, error)
	FetchProductVariantByID(ctx context.Context, productID int) ([]entity.ProductVariant, error)
}
