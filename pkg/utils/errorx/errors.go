package errorx

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// AppError để đảm bảo chỉ trả về những error này cho người dùng.
type AppError error

var (
	ErrInvalidRequest      AppError = errors.New("invalid request")
	ErrValidationFailed    AppError = errors.New("validation failed")
	ErrEmailAlready        AppError = errors.New("email already exists")
	ErrDatabaseQueryFailed AppError = errors.New("database query failed")
	ErrSomethingWrong      AppError = errors.New("something wrong")
	ErrSignUpFailed        AppError = errors.New("sign up failed")

	ErrPasswordNotMatch AppError = errors.New("password is not match")

	// token.
	ErrRefreshTokenNotFound AppError = errors.New("refresh token not found")
	ErrInvalidRefreshToken  AppError = errors.New("refresh token is invalid")

	// products.
	ErrGetProductsFailed                 AppError = errors.New("failed to get products")
	ErrGetProductDetailFailed            AppError = errors.New("failed to get product detail")
	ErrGetProductDetailConvertParamError AppError = errors.New("get product detail convert param error")

	// safe type.
	ErrOverflow AppError = errors.New("value over flow")
)

//nolint:err113
func PrettyValidationErrors(err error) error {
	var sb strings.Builder
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range ve {
			sb.WriteString(fmt.Sprintf(
				"Field '%s' failed validation rule '%s': ", e.Field(), e.Tag()),
			)
			if e.Param() != "" {
				sb.WriteString(fmt.Sprintf("expect '%s=%s', actual '%v'. ", e.Tag(), e.Param(), e.Value()))
			} else {
				sb.WriteString(fmt.Sprintf("expect a valid %s. ", e.Tag()))
			}
		}
	} else {
		sb.WriteString(err.Error())
	}

	return errors.New(strings.TrimSpace(sb.String()))
}
