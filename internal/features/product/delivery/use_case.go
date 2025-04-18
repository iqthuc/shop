package delivery

import (
	"context"
	"shop/internal/features/product/dto"
)

type UseCase interface {
	GetProducts(
		ctx context.Context,
		input dto.GetProductsRequest,
	) (*dto.GetProductsResult[dto.Product], error)

	GetProductDetail(
		ctx context.Context,
		productID int,
	) (*dto.GetProductDetailResult, error)
}
