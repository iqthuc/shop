package auth

import (
	"context"
	"log/slog"
	"shop/pkg/utils"
	errs "shop/pkg/utils/errors"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	CreateUser(user createUserParams) (*createUserResult, error)
	GetUser(ctx context.Context, user loginRequest) (*User, error)
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
func (u useCase) Login(ctx context.Context, input loginRequest) (*loginResponse, error) {
	if err := u.validator.Struct(input); err != nil {
		return nil, errs.ErrVaidationFailed
	}

	user, err := u.repo.GetUser(ctx, input)
	if err != nil {
		slog.Debug("get user by email failed", slog.String("error", err.Error()))
		return nil, err
	}

	result := &loginResponse{
		UserID: user.id,
	}

	return result, nil
}

func (u useCase) SignUp(ctx context.Context, input signUpInput) (*signUpResult, error) {
	if err := u.validator.Struct(input); err != nil {
		return nil, errs.ErrVaidationFailed
	}

	passwordHash, err := utils.HashPassword(input.Password)
	if err != nil {
		slog.Debug("failed to hash password", slog.String("error", err.Error()))
		return nil, err
	}

	params := createUserParams{
		email:        input.Email,
		passwordHash: passwordHash,
	}

	user, err := u.repo.CreateUser(params)
	if err != nil {
		slog.Debug("failed to create user", slog.String("error", err.Error()))
		return nil, err
	}

	result := &signUpResult{
		email:     user.email,
		createdAt: user.createdAt,
	}

	return result, nil
}
