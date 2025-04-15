package errs

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
	//token
	ErrInvalidToken AppError = errors.New("token is invalid")
)

func WrapValidationFailed(err error) error {
	if _, ok := err.(validator.ValidationErrors); ok {
		return fmt.Errorf("%w", err)
	}

	return ErrValidationFailed
}

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
