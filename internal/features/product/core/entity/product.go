package entity

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID        int
	Name      string
	BasePrice decimal.Decimal
}

type ProductDetail struct {
	ID           int
	Name         string
	Slug         string
	Desciprtion  string
	MainImageUrl string
	BasePrice    decimal.Decimal
	CategoryID   int
	CategoryName string
	BrandID      int
	BrandName    string
}

type ProductVariant struct {
	ID            int32
	ProductID     int32
	Sku           string
	Price         decimal.Decimal
	StockQuantity int32
	Sold          int32
	ImageUrl      sql.NullString
	IsDefault     bool
}
