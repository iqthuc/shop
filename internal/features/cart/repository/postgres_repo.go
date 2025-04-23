package repository

import (
	"context"
	"database/sql"
	"errors"
	"shop/internal/features/cart/core"
	"shop/internal/features/cart/core/dto"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"

	"github.com/jackc/pgx/v5"
)

type CartPostgresRepo struct {
	store store.Store
}

func NewCartPostgresRepo(store store.Store) CartPostgresRepo {
	return CartPostgresRepo{
		store: store,
	}
}

func (r CartPostgresRepo) AddToCart(ctx context.Context, item dto.AddToCartRequest) error {
	err := r.store.ExecTx(ctx, pgx.TxOptions{}, func(q *db.Queries) error {
		_, err := q.CheckIfVariantStockSufficient(ctx,
			db.CheckIfVariantStockSufficientParams{
				ID:            item.VariantID,
				StockQuantity: item.Quantity,
			})
		if err != nil {
			return err
		}

		cart_id, err := q.UpsertCarts(ctx, item.UserID)
		if err != nil {
			return err
		}

		err = q.UpsertToCartItems(ctx,
			db.UpsertToCartItemsParams{
				CartID:           cart_id,
				ProductVariantID: item.VariantID,
				Quantity:         item.Quantity,
			})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.ErrVariantOrStockInvalid
		}

		return err
	}

	return nil
}
