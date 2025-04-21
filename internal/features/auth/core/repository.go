package core

import (
	"context"
	"shop/internal/features/auth/core/entity"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, email string) (*entity.User, error)
}
