package auth

import (
	"context"
	"errors"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"
	errs "shop/pkg/utils/errors"

	"github.com/jackc/pgx/v5/pgconn"
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

func (r repository) GetUser(ctx context.Context, input loginRequest) (*User, error) {
	user, err := r.store.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	u := &User{
		id:    user.ID.Bytes,
		email: user.Email,
	}

	return u, nil
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
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.SQLState() == "23505" {
			return nil, errs.ErrEmailAlready
		}

		return nil, errs.ErrDatabaseQueryFailed
	}

	result := &createUserResult{
		email:     user.Email,
		createdAt: user.CreatedAt.Time,
	}

	return result, nil
}
