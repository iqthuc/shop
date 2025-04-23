package dto

import "github.com/google/uuid"

type AddToCartRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	VariantID int32     `json:"variant_id"`
	Quantity  int32     `json:"quantity"`
}
