package messages

// AppMessage để đảm bảo chỉ trả về những message này cho người dùng
type AppMessage string

const (
	SignUpSuccess AppMessage = "Sign up success"
)
