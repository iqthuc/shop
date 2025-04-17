package entity

type Product struct {
	ID           int
	Name         string
	Description  string
	MainImageUrl string
	BasePrice    float64
}

type ProductDetail struct {
	ID           int
	Name         string
	Slug         string
	Desciprtion  string
	MainImageUrl string
	BasePrice    float64
	CategoryID   int
	CategoryName string
	BrandID      int
	BrandName    string
}

type ProductVariant struct {
	ID            int
	ProductID     int
	Sku           string
	Price         float64
	StockQuantity int
	Sold          int
	ImageUrl      string
	IsDefault     bool
}
