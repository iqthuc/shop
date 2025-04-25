package dto

import "github.com/google/uuid"

type OrderRequest struct {
	UserID uuid.UUID `json:"cart_id"`
	Items  []int32   `json:"items"`
}
