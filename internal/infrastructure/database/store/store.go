package store

// package này nên đặt trong foldler db
// nhưng đặt đây cho dễ sửa/xóa
// vì folder db được auto generate
import (
	"context"
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

func NewPostgresStore(conn *pgx.Conn) Store {
	return &PostgresStore{
		conn:    conn,
		Queries: db.New(conn),
	}
}

func (s *PostgresStore) CloseDB(ctx context.Context) error {
	return s.conn.Close(ctx)
}
