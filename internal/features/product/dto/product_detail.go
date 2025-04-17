package dto

import "shop/internal/features/product/entity"

type GetProductDetailResult struct {
	Detail   *entity.ProductDetail   `json:"detail"`
	Variants []entity.ProductVariant `json:"variants"`
}
