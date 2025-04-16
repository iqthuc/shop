package use_case

import (
	"context"
	"log/slog"
	"shop/internal/features/auth/dto"
	"shop/internal/features/auth/entity"
	"shop/pkg/token"
	errs "shop/pkg/utils/errors"
	"shop/pkg/utils/password"
	"time"
)

type useCase struct {
	repo       Repository
	tokenMaker token.TokenMaker
}

func NewUseCase(repo Repository, tk token.TokenMaker) useCase {
	return useCase{
		repo:       repo,
		tokenMaker: tk,
	}
}

func (u useCase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	rfToken, err := u.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		slog.Debug("verify token failed", slog.String("error", err.Error()))
		return "", err
	}

	if ok := rfToken.Valid(); ok != nil {
		return "", errs.ErrInvalidRefreshToken
	}

	const accessTokenLifetime = 15 * time.Minute
	newAccessToken, err := u.tokenMaker.CreateToken(rfToken.UserID, rfToken.Role, rfToken.TokenType, accessTokenLifetime)
	if err != nil {
		slog.Debug("create new access token failed", slog.String("error", err.Error()))
		return "", err
	}

	return newAccessToken, nil
}

func (u useCase) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginResult, error) {
	user, err := u.repo.GetUser(ctx, input.Email)
	if err != nil {
		slog.Debug("get user by email failed", slog.String("error", err.Error()))
		return nil, err
	}

	if !password.CheckPasswordHash(input.Password, user.PasswordHash) {
		return nil, errs.ErrPasswordNotMatch
	}

	const accessTokenLifetime = 15 * time.Minute
	accessToken, err := u.tokenMaker.CreateToken(user.ID.String(), user.Role, token.Access, accessTokenLifetime)
	if err != nil {
		slog.Debug("create access token failed", slog.String("error", err.Error()))
		return nil, err
	}

	const refreshTokenLifetime = 24 * time.Hour
	refreshToken, err := u.tokenMaker.CreateToken(user.ID.String(), user.Role, token.Refresh, refreshTokenLifetime)
	if err != nil {
		slog.Debug("create access token failed", slog.String("error", err.Error()))
		return nil, err
	}

	result := dto.LoginResult{
		UserID:           user.ID,
		AccessToken:      accessToken,
		ExpiresIn:        accessTokenLifetime,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: refreshTokenLifetime,
		TokenType:        "bearer",
	}

	return &result, nil
}

func (u useCase) SignUp(ctx context.Context, req dto.SignUpInput) error {
	passwordHash, err := password.HashPassword(req.Password)
	if err != nil {
		slog.Debug("failed to hash password", slog.String("error", err.Error()))
		return err
	}

	params := entity.User{
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	err = u.repo.CreateUser(ctx, params)
	if err != nil {
		slog.Debug("failed to create user", slog.String("error", err.Error()))
		return err
	}

	return nil
}
