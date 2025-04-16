package use_case

import (
	"context"
	"shop/internal/features/auth/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, email string) (*entity.User, error)
}
