package messages

// AppMessage để đảm bảo chỉ trả về những message này cho người dùng.
type AppMessage string

const (
	SignUpSuccess  AppMessage = "sign up success"
	LoginSuccess   AppMessage = "login success"
	LogoutSuccess  AppMessage = "logout success"
	RefreshSuccess AppMessage = "refresh token success"
)
