package delivery

import (
	"context"
	"shop/internal/features/auth/dto"
)

type UseCase interface {
	SignUp(ctx context.Context, req dto.SignUpInput) error
	Login(ctx context.Context, input dto.LoginInput) (*dto.LoginResult, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}
