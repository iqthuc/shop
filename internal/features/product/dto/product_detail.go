package dto

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

type GetProductDetailResult struct {
	Detail   *ProductDetail   `json:"product"`
	Variants []ProductVariant `json:"variants"`
}

type ProductDetail struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Slug         string          `json:"slug"`
	Desciprtion  string          `json:"description"`
	MainImageUrl string          `json:"main_image_url"`
	BasePrice    decimal.Decimal `json:"base_price"`
	CategoryID   int             `json:"category_id"`
	CategoryName string          `json:"category_name"`
	BrandID      int             `json:"brand_id"`
	BrandName    string          `json:"brand_name"`
}

type ProductVariant struct {
	ID            int32           `json:"id"`
	ProductID     int32           `json:"product_id"`
	Sku           string          `json:"sku"`
	Price         decimal.Decimal `json:"price"`
	StockQuantity int32           `json:"stock_quantity"`
	Sold          int32           `json:"sold"`
	ImageUrl      sql.NullString  `json:"image_url"`
	IsDefault     bool            `json:"is_default"`
}
