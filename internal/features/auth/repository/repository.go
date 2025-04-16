package repository

import (
	"context"
	"errors"
	"shop/internal/features/auth/entity"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"
	errs "shop/pkg/utils/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type repository struct {
	store store.Store
}

func NewRepository(store store.Store) repository {
	return repository{
		store: store,
	}
}

func (r repository) GetUser(ctx context.Context, email string) (*entity.User, error) {
	user, err := r.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
		ID:           user.ID.Bytes,
		Email:        user.Email,
		PasswordHash: user.PasswordHash.String,
	}

	return u, nil
}

func (r repository) CreateUser(ctx context.Context, u entity.User) error {
	params := db.CreateUserParams{
		Email: u.Email,
		PasswordHash: pgtype.Text{
			String: u.PasswordHash,
			Valid:  true,
		},
	}

	_, err := r.store.CreateUser(context.Background(), params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.SQLState() == "23505" {
			return errs.ErrEmailAlready
		}

		return errs.ErrDatabaseQueryFailed
	}

	return nil
}
