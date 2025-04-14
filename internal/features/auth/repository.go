package auth

import (
	"context"
	"fmt"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type repository struct {
	store store.Store
}

func NewRepository(store store.Store) Repository {
	return repository{
		store: store,
	}
}

func (r repository) CreateUser(u createUserParams) (*createUserResult, error) {
	params := db.CreateUserParams{
		Email: u.email,
		PasswordHash: pgtype.Text{
			String: u.passwordHash,
			Valid:  true,
		},
	}
	user, err := r.store.CreateUser(context.Background(), params)
	// TODO: check loi duplicate
	if err != nil {
		return nil, fmt.Errorf("db query failed: %w", err)
	}
	result := &createUserResult{
		email:     user.Email,
		createdAt: user.CreatedAt.Time,
	}
	return result, nil
}
