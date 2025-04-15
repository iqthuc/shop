package auth

import (
	"context"
	"log/slog"
	"shop/pkg/token"
	errs "shop/pkg/utils/errors"
	"shop/pkg/utils/password"
	"time"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	CreateUser(user createUserParams) (*createUserResult, error)
	GetUser(ctx context.Context, user loginRequest) (*User, error)
}

type useCase struct {
	repo       Repository
	validator  validator.Validate
	tokenMaker token.TokenMaker
}

func NewUseCase(repo Repository, v validator.Validate, tk token.TokenMaker) useCase {
	return useCase{
		repo:       repo,
		validator:  v,
		tokenMaker: tk,
	}
}
func (u useCase) Login(ctx context.Context, input loginRequest) (*loginResponse, error) {
	if err := u.validator.Struct(input); err != nil {
		return nil, errs.ErrValidationFailed
	}

	user, err := u.repo.GetUser(ctx, input)
	if err != nil {
		slog.Debug("get user by email failed", slog.String("error", err.Error()))
		return nil, err
	}
	slog.Debug("check password", slog.String("password", user.password))

	if !password.CheckPasswordHash(input.Password, user.password) {
		slog.Info("check password")
		return nil, errs.ErrPasswordNotMatch
	}

	const accessTokenLifetime = 15 * time.Minute
	accessToken, err := u.tokenMaker.CreateAccessToken(user.id.String(), user.role, token.Access, accessTokenLifetime)
	if err != nil {
		slog.Debug("create access token failed", slog.String("error", err.Error()))

		return nil, err
	}

	result := loginResponse{
		UserID:      user.id,
		AccessToken: accessToken,
	}

	return &result, nil
}

func (u useCase) SignUp(ctx context.Context, input signUpInput) (*signUpResult, error) {
	if err := u.validator.Struct(input); err != nil {
		return nil, errs.ErrValidationFailed
	}

	passwordHash, err := password.HashPassword(input.Password)
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
