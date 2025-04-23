package messages

// AppMessage để đảm bảo chỉ trả về những message này cho người dùng.
type AppMessage string

func (m AppMessage) String() string {
	return string(m)
}

const (
	SignUpSuccess  AppMessage = "sign up success"
	LoginSuccess   AppMessage = "login success"
	LogoutSuccess  AppMessage = "logout success"
	RefreshSuccess AppMessage = "refresh token success"

	// product.
	GetProductsSuccess      AppMessage = "get products success"
	GetProductDetailSuccess AppMessage = "get product detail success"

	// cart.
	GetCartsSuccess  AppMessage = "get carts success"
	AddToCartSuccess AppMessage = "add to cart success"
)
