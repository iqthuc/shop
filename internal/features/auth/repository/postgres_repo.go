package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"shop/internal/features/auth/core/entity"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"
	"shop/pkg/utils/errorx"

	"github.com/jackc/pgx/v5/pgconn"
)

type authPostgreRepo struct {
	store store.Store
}

func NewAuthPostgreRepo(store store.Store) authPostgreRepo {
	return authPostgreRepo{
		store: store,
	}
}

func (r authPostgreRepo) GetUser(ctx context.Context, email string) (*entity.User, error) {
	user, err := r.store.GetUserByEmail(ctx, email)
	if err != nil {
		slog.Debug("failed to get user by email", slog.String("error", err.Error()))
		return nil, err
	}

	u := &entity.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash.String,
		Role:         user.Role,
		CreatedAt:    user.CreatedAt,
	}

	return u, nil
}

func (r authPostgreRepo) CreateUser(ctx context.Context, u entity.User) error {
	params := db.CreateUserParams{
		Email: u.Email,
		PasswordHash: sql.NullString{
			String: u.PasswordHash,
			Valid:  u.PasswordHash != "",
		},
	}

	_, err := r.store.CreateUser(ctx, params)
	if err != nil {
		slog.Debug("failed to create user", slog.String("error", err.Error()))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.SQLState() == "23505" {
			return errorx.ErrEmailAlready
		}

		return errorx.ErrDatabaseQueryFailed
	}

	return nil
}
