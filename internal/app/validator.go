package app

import "github.com/go-playground/validator/v10"

func newValidator() *validator.Validate {
	return validator.New()
}
