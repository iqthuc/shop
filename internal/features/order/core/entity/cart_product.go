package entity

import "github.com/shopspring/decimal"

type CartProduct struct {
	ProductVariantID int32
	Quantity         int32
	Price            decimal.Decimal
}
