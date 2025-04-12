package app

import (
	"context"
	"log"
	"shop/internal/infrastructure/config"
	"shop/internal/infrastructure/database/store"

	"github.com/jackc/pgx/v5"
)

func newStore(cfg *config.Database) store.Store {
	conn, err := pgx.Connect(context.Background(), cfg.DataSourceName())
	if err != nil {
		log.Panicf("cannot open postgresql: %s ", err)
	}
	store := store.NewPostgresStore(conn)
	return store
}
