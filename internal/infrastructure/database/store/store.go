package store

import (
	"context"
	"log"
	"shop/internal/infrastructure/config"
	"shop/internal/infrastructure/database/store/db"

	"github.com/jackc/pgx/v5"
)

type Store interface {
	db.Querier
	CloseDB(ctx context.Context) error
}
type PostgresStore struct {
	conn *pgx.Conn
	*db.Queries
}

func New(cfg *config.Database) Store {
	conn, err := pgx.Connect(context.Background(), cfg.DataSourceName())
	if err != nil {
		log.Panicf("cannot open postgresql: %s ", err)
	}

	store := &PostgresStore{
		conn:    conn,
		Queries: db.New(conn),
	}

	return store
}

func (s *PostgresStore) CloseDB(ctx context.Context) error {
	return s.conn.Close(ctx)
}
