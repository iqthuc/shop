package auth

import (
	"context"
	"fmt"
	"shop/pkg/utils"
	errs "shop/pkg/utils/errors"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	CreateUser(user createUserParams) (*createUserResult, error)
}

type useCase struct {
	repo      Repository
	validator validator.Validate
}

func NewUsecase(repo Repository, v validator.Validate) useCase {
	return useCase{
		repo:      repo,
		validator: v,
	}
}

func (u useCase) SignUp(ctx context.Context, input signUpInput) (*signUpResult, error) {
	if err := u.validator.Struct(input); err != nil {
		return nil, errs.VaidationFailed
	}
	passwordHash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	params := createUserParams{
		email:        input.Email,
		passwordHash: passwordHash,
	}
	user, err := u.repo.CreateUser(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	result := &signUpResult{
		email:     user.email,
		createdAt: user.createdAt,
	}
	return result, nil
}
