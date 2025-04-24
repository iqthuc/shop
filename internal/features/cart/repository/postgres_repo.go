package repository

import (
	"context"
	"database/sql"
	"errors"
	"shop/internal/features/cart/core"
	"shop/internal/features/cart/core/dto"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"
	"shop/pkg/utils/errorx"
	"shop/pkg/utils/safetype"

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

func (r CartPostgresRepo) SaveCartItem(ctx context.Context, item dto.AddToCartRequest) error {
	variantID, err1 := safetype.SafeIntToInt32(item.VariantID)
	quantity, err2 := safetype.SafeIntToInt32(item.Quantity)
	if err1 != nil || err2 != nil {
		return errorx.ErrOverflow
	}
	err := r.store.ExecTx(ctx, pgx.TxOptions{}, func(q *db.Queries) error {
		_, err := q.CheckIfVariantStockSufficient(ctx,
			db.CheckIfVariantStockSufficientParams{
				ID:            variantID,
				StockQuantity: quantity,
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
				ProductVariantID: variantID,
				Quantity:         quantity,
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

func (r CartPostgresRepo) UpdateCart(ctx context.Context, input dto.UpdateCartRequest) error {
	variantID, err1 := safetype.SafeIntToInt32(input.VariantID)
	quantity, err2 := safetype.SafeIntToInt32(input.Quantity)
	if err1 != nil || err2 != nil {
		return errorx.ErrOverflow
	}

	err := r.store.UpdateCartItem(ctx, db.UpdateCartItemParams{
		UserID:           input.UserID,
		ProductVariantID: variantID,
		Quantity:         quantity,
	})

	return err
}

func (r CartPostgresRepo) DeleteCartItem(ctx context.Context, input dto.DeleteCartItemRequest) error {
	variantID, err1 := safetype.SafeIntToInt32(input.VariantID)
	if err1 != nil {
		return errorx.ErrOverflow
	}

	err := r.store.DeleteCartItem(ctx, db.DeleteCartItemParams{
		UserID:           input.UserID,
		ProductVariantID: variantID,
	})

	return err
}
