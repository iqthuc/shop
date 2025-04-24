package store

import (
	"context"
	"log"
	"shop/internal/infrastructure/config"
	"shop/internal/infrastructure/database/store/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	db.Querier
	CloseDB(ctx context.Context)
	ExecTx(ctx context.Context, opts pgx.TxOptions, fn func(*db.Queries) error) error
}

type PostgresStore struct {
	pool *pgxpool.Pool
	*db.Queries
}

func NewPostgresStore(cfg *config.Database) *PostgresStore {
	pool, err := pgxpool.New(context.Background(), cfg.DataSourceName())
	if err != nil {
		log.Panicf("cannot open postgresql: %s ", err)
	}

	store := &PostgresStore{
		pool:    pool,
		Queries: db.New(pool),
	}

	return store
}

func (s *PostgresStore) CloseDB(ctx context.Context) {
	s.pool.Close()
}

func (store *PostgresStore) ExecTx(ctx context.Context, opts pgx.TxOptions, fn func(*db.Queries) error) error {
	tx, err := store.pool.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return rbErr
		}

		return err
	}

	return tx.Commit(ctx)
}
