package database

import (
	"database/sql"
	"fmt"
	"shop/internal/infrastructure/config"
)

func NewPostPresDB(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DataSourceName())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil

}
