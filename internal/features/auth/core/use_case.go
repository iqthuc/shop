package core

import (
	"context"
	"log/slog"
	"shop/internal/features/auth/core/dto"
	"shop/internal/features/auth/core/entity"
	"shop/pkg/token"
	"shop/pkg/utils/errorx"
	"shop/pkg/utils/password"
	"time"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, req dto.SignUpInput) error
	Login(ctx context.Context, input dto.LoginInput) (*dto.LoginResult, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type useCase struct {
	repo       AuthRepository
	tokenMaker token.TokenMaker
}

func NewUseCase(repo AuthRepository, tk token.TokenMaker) useCase {
	return useCase{
		repo:       repo,
		tokenMaker: tk,
	}
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

func (u useCase) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginResult, error) {
	user, err := u.repo.GetUser(ctx, input.Email)
	if err != nil {
		slog.Debug("get user by email failed", slog.String("error", err.Error()))
		return nil, err
	}

	if !password.CheckPasswordHash(input.Password, user.PasswordHash) {
		return nil, errorx.ErrPasswordNotMatch
	}

	const accessTokenLifetime = 15 * time.Minute
	accessTokenInfo := token.CreateTokenParams{
		UserID:    user.ID,
		Role:      user.Role,
		TokenType: token.Access,
		Duration:  accessTokenLifetime,
	}
	accessToken, err := u.tokenMaker.CreateToken(accessTokenInfo)
	if err != nil {
		slog.Debug("create access token failed", slog.String("error", err.Error()))
		return nil, err
	}

	const refreshTokenLifetime = 24 * time.Hour
	refreshTokenInfo := token.CreateTokenParams{
		UserID:    user.ID,
		Role:      user.Role,
		TokenType: token.Refresh,
		Duration:  refreshTokenLifetime,
	}
	refreshToken, err := u.tokenMaker.CreateToken(refreshTokenInfo)
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

func (u useCase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	rfToken, err := u.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		slog.Debug("verify token failed", slog.String("error", err.Error()))
		return "", err
	}

	if ok := rfToken.Valid(); ok != nil {
		return "", errorx.ErrInvalidRefreshToken
	}

	const accessTokenLifetime = 15 * time.Minute
	accessTokenInfo := token.CreateTokenParams{
		UserID:    rfToken.UserID,
		Role:      rfToken.Role,
		TokenType: token.Access,
		Duration:  accessTokenLifetime,
	}
	newAccessToken, err := u.tokenMaker.CreateToken(accessTokenInfo)
	if err != nil {
		slog.Debug("create new access token failed", slog.String("error", err.Error()))
		return "", err
	}

	return newAccessToken, nil
}
