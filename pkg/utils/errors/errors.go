package errs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// AppError để đảm bảo chỉ trả về những error này cho người dùng
type AppError error

var (
	InvalidRequest  AppError = errors.New("Invalid request.")
	VaidationFailed AppError = errors.New("Validation failed.")
	EmailAlready    AppError = errors.New("Email already exists.")
	SomethingWrong  AppError = errors.New("Something wrong.")
	SignUpFailed    AppError = errors.New("Sign up failed.")
)

func WrapValidationFailed(err error) error {
	if _, ok := err.(validator.ValidationErrors); ok {
		return fmt.Errorf("%w", err)
	}
	return VaidationFailed
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
