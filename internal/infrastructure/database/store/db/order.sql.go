// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const calculateTotalPriceByProductIDs = `-- name: CalculateTotalPriceByProductIDs :one
SELECT COALESCE(SUM(price), 0)::numeric AS total
FROM product_variants
WHERE product_variant_id = ANY($1::int[])
`

func (q *Queries) CalculateTotalPriceByProductIDs(ctx context.Context, variantIds []int32) (decimal.Decimal, error) {
	row := q.db.QueryRow(ctx, calculateTotalPriceByProductIDs, variantIds)
	var total decimal.Decimal
	err := row.Scan(&total)
	return total, err
}

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (user_id, total_amount)
VALUES ($1, $2)
RETURNING id
`

type CreateOrderParams struct {
	UserID      uuid.UUID       `json:"user_id"`
	TotalAmount decimal.Decimal `json:"total_amount"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (int32, error) {
	row := q.db.QueryRow(ctx, createOrder, arg.UserID, arg.TotalAmount)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createOrderItem = `-- name: CreateOrderItem :exec
INSERT INTO order_items (order_id, product_variant_id, quantity, price)
VALUES ($1, $2, $3, $4)
`

type CreateOrderItemParams struct {
	OrderID          int32       `json:"order_id"`
	ProductVariantID int32       `json:"product_variant_id"`
	Quantity         int32       `json:"quantity"`
	Price            interface{} `json:"price"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) error {
	_, err := q.db.Exec(ctx, createOrderItem,
		arg.OrderID,
		arg.ProductVariantID,
		arg.Quantity,
		arg.Price,
	)
	return err
}

const decreaseProductStock = `-- name: DecreaseProductStock :exec
UPDATE product_variants
SET stock_quantity = stock_quantity - $2
WHERE id = $1 AND stock_quantity >= $2
`

type DecreaseProductStockParams struct {
	ID            int32 `json:"id"`
	StockQuantity int32 `json:"stock_quantity"`
}

func (q *Queries) DecreaseProductStock(ctx context.Context, arg DecreaseProductStockParams) error {
	_, err := q.db.Exec(ctx, decreaseProductStock, arg.ID, arg.StockQuantity)
	return err
}

const deleteCartItemsByProductIDs = `-- name: DeleteCartItemsByProductIDs :exec
DELETE FROM cart_items
WHERE cart_id IN (
    SELECT id FROM carts WHERE user_id = $1
)
AND product_variant_id = ANY($2::int[])
`

type DeleteCartItemsByProductIDsParams struct {
	UserID     uuid.UUID `json:"user_id"`
	VariantIds []int32   `json:"variant_ids"`
}

func (q *Queries) DeleteCartItemsByProductIDs(ctx context.Context, arg DeleteCartItemsByProductIDsParams) error {
	_, err := q.db.Exec(ctx, deleteCartItemsByProductIDs, arg.UserID, arg.VariantIds)
	return err
}

const findOutOfStockItems = `-- name: FindOutOfStockItems :many
SELECT ci.product_variant_id, ci.quantity, pv.stock_quantity
FROM cart_items AS ci
JOIN carts AS c ON ci.cart_id = c.id
JOIN product_variants AS pv ON ci.product_variant_id = pv.id
WHERE c.user_id = $1
AND ci.product_variant_id = ANY($2::int[])
AND ci.quantity > pv.stock_quantity
`

type FindOutOfStockItemsParams struct {
	UserID     uuid.UUID `json:"user_id"`
	VariantIds []int32   `json:"variant_ids"`
}

type FindOutOfStockItemsRow struct {
	ProductVariantID int32 `json:"product_variant_id"`
	Quantity         int32 `json:"quantity"`
	StockQuantity    int32 `json:"stock_quantity"`
}

func (q *Queries) FindOutOfStockItems(ctx context.Context, arg FindOutOfStockItemsParams) ([]FindOutOfStockItemsRow, error) {
	rows, err := q.db.Query(ctx, findOutOfStockItems, arg.UserID, arg.VariantIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindOutOfStockItemsRow{}
	for rows.Next() {
		var i FindOutOfStockItemsRow
		if err := rows.Scan(&i.ProductVariantID, &i.Quantity, &i.StockQuantity); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserCartProducts = `-- name: GetUserCartProducts :many
SELECT
    pv.id AS product_variant_id,
    pv.price,
    pv.stock_quantity,
    ci.quantity
FROM cart_items AS ci
JOIN carts AS c ON ci.cart_id = c.id
JOIN product_variants AS pv ON ci.product_variant_id = pv.id
WHERE c.user_id = $1
AND ci.product_variant_id = ANY($2::int[])
`

type GetUserCartProductsParams struct {
	UserID     uuid.UUID `json:"user_id"`
	VariantIds []int32   `json:"variant_ids"`
}

type GetUserCartProductsRow struct {
	ProductVariantID int32           `json:"product_variant_id"`
	Price            decimal.Decimal `json:"price"`
	StockQuantity    int32           `json:"stock_quantity"`
	Quantity         int32           `json:"quantity"`
}

func (q *Queries) GetUserCartProducts(ctx context.Context, arg GetUserCartProductsParams) ([]GetUserCartProductsRow, error) {
	rows, err := q.db.Query(ctx, getUserCartProducts, arg.UserID, arg.VariantIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserCartProductsRow{}
	for rows.Next() {
		var i GetUserCartProductsRow
		if err := rows.Scan(
			&i.ProductVariantID,
			&i.Price,
			&i.StockQuantity,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const lockProductVariantsForUpdate = `-- name: LockProductVariantsForUpdate :many
SELECT id
FROM product_variants
WHERE id = ANY($1::int[])
ORDER BY id
FOR UPDATE
`

func (q *Queries) LockProductVariantsForUpdate(ctx context.Context, variantIds []int32) ([]int32, error) {
	rows, err := q.db.Query(ctx, lockProductVariantsForUpdate, variantIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
