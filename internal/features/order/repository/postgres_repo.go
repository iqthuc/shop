package repository

import (
	"context"
	"shop/internal/features/order/core"
	"shop/internal/features/order/core/dto"
	"shop/internal/features/order/core/entity"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"

	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

type OrderPostgresRepo struct {
	Store store.Store
}

func NewCartPostgresRepo(store store.Store) OrderPostgresRepo {
	return OrderPostgresRepo{
		Store: store,
	}
}

/*
Process selected cart items to create an order:
1. Lock product variants for update.
2. Validate selected product IDs and quantities.
3. Check and lock product stock (FOR UPDATE).
4. Create the order and related order items.
5. Deduct product stock accordingly.
6. Remove purchased items from the cart.
*/
func (r OrderPostgresRepo) CreateOrder(ctx context.Context, req dto.OrderRequest) error {
	variantIDs := make([]int32, len(req.Items))
	copy(variantIDs, req.Items)
	slices.Sort(variantIDs)

	return r.Store.ExecTx(ctx, pgx.TxOptions{}, func(q *db.Queries) error {
		if err := lockVariants(ctx, q, variantIDs); err != nil {
			return err
		}

		products, totalAmount, err := getCartProducts(ctx, q, req.UserID, variantIDs)
		if err != nil {
			return err
		}

		if err := checkStock(ctx, q, req.UserID, variantIDs); err != nil {
			return err
		}

		err = createOrderAndItems(ctx, q, req.UserID, products, totalAmount)
		if err != nil {
			return err
		}

		if err := deleteCartItems(ctx, q, req.UserID, variantIDs); err != nil {
			return err
		}

		return nil
	})
}
func lockVariants(ctx context.Context, q *db.Queries, ids []int32) error {
	_, err := q.LockProductVariantsForUpdate(ctx, ids)

	return err
}

func getCartProducts(
	ctx context.Context,
	q *db.Queries,
	userID uuid.UUID,
	ids []int32,
) ([]entity.CartProduct, decimal.Decimal, error) {
	raw, err := q.GetUserCartProducts(ctx, db.GetUserCartProductsParams{
		UserID:     userID,
		VariantIds: ids,
	})

	if err != nil || len(raw) != len(ids) {
		return nil, decimal.Zero, core.ErrVariantInvalid
	}

	products := make([]entity.CartProduct, len(raw))
	var total decimal.Decimal
	for i, p := range raw {
		quantity := decimal.NewFromInt(int64(p.Quantity))
		total = total.Add(p.Price.Mul(quantity))
		products[i] = entity.CartProduct{
			ProductVariantID: p.ProductVariantID,
			Quantity:         p.Quantity,
			Price:            p.Price,
		}
	}

	return products, total, nil
}

func checkStock(ctx context.Context, q *db.Queries, userID uuid.UUID, ids []int32) error {
	items, err := q.FindOutOfStockItems(ctx, db.FindOutOfStockItemsParams{
		UserID:     userID,
		VariantIds: ids,
	})

	if err != nil || len(items) > 0 {
		return core.ErrStockInvalid
	}

	return nil
}

func createOrderAndItems(
	ctx context.Context,
	q *db.Queries,
	userID uuid.UUID,
	products []entity.CartProduct,
	total decimal.Decimal,
) error {
	orderID, err := q.CreateOrder(ctx, db.CreateOrderParams{
		UserID:      userID,
		TotalAmount: total,
	})

	if err != nil {
		return err
	}

	for _, item := range products {
		if err := q.CreateOrderItem(ctx, db.CreateOrderItemParams{
			OrderID:          orderID,
			ProductVariantID: item.ProductVariantID,
			Quantity:         item.Quantity,
			Price:            item.Price,
		}); err != nil {
			return err
		}

		if err := q.DecreaseProductStock(ctx, db.DecreaseProductStockParams{
			ID:            item.ProductVariantID,
			StockQuantity: item.Quantity,
		}); err != nil {
			return err
		}
	}

	return nil
}

func deleteCartItems(ctx context.Context, q *db.Queries, userID uuid.UUID, ids []int32) error {
	return q.DeleteCartItemsByProductIDs(ctx, db.DeleteCartItemsByProductIDsParams{
		UserID:     userID,
		VariantIds: ids,
	})
}
