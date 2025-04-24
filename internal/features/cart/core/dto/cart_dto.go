package dto

import "github.com/google/uuid"

type AddToCartRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	VariantID int       `json:"variant_id"`
	Quantity  int       `json:"quantity"`
}

type UpdateCartRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	VariantID int       `json:"variant_id"`
	Quantity  int       `json:"quantity"`
}

type DeleteCartItemRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	VariantID int       `json:"variant_id"`
}
