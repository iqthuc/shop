package dto

import "github.com/shopspring/decimal"

type Product struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	BasePrice decimal.Decimal `json:"base_price"`
}

type GetProductsResult[T any] struct {
	Items      []T `json:"items"`
	*Filter    `json:"filters,omitempty"`
	*SortBy    `json:"sort_by,omitempty"`
	Pagination `json:"pagination"`
}

type GetProductsRequest struct {
	Page    int
	Filters Filter
	SortBy  SortBy
}

type GetProductsParams struct {
	Limit   int
	Offset  int
	Filters Filter
	SortBy  SortBy
}

type Filter struct {
	Keyword    string  `json:"key_word,omitempty"`
	CategoryID int     `json:"category_id,omitempty"`
	PriceMin   float64 `json:"price_min,omitempty"`
	PriceMax   float64 `json:"price_max,omitempty"`
}

type SortBy struct {
	Field string `json:"field,omitempty"`
	Order string `json:"order,omitempty"`
}

type Pagination struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
}
