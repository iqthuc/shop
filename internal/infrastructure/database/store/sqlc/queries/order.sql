-- name: LockProductVariantsForUpdate :many
SELECT id
FROM product_variants
WHERE id = ANY(sqlc.arg(variant_ids)::int[])
ORDER BY id
FOR UPDATE;

-- name: GetUserCartProducts :many
SELECT
    pv.id AS product_variant_id,
    pv.price,
    pv.stock_quantity,
    ci.quantity
FROM cart_items AS ci
JOIN carts AS c ON ci.cart_id = c.id
JOIN product_variants AS pv ON ci.product_variant_id = pv.id
WHERE c.user_id = $1
AND ci.product_variant_id = ANY(sqlc.arg(variant_ids)::int[]);

-- name: FindOutOfStockItems :many
SELECT ci.product_variant_id, ci.quantity, pv.stock_quantity
FROM cart_items AS ci
JOIN carts AS c ON ci.cart_id = c.id
JOIN product_variants AS pv ON ci.product_variant_id = pv.id
WHERE c.user_id = $1
AND ci.product_variant_id = ANY(sqlc.arg(variant_ids)::int[])
AND ci.quantity > pv.stock_quantity;

-- name: CalculateTotalPriceByProductIDs :one
SELECT COALESCE(SUM(price), 0)::numeric AS total
FROM product_variants
WHERE product_variant_id = ANY(sqlc.arg(variant_ids)::int[]);

-- name: CreateOrder :one
INSERT INTO orders (user_id, total_amount)
VALUES ($1, $2)
RETURNING id;

-- name: CreateOrderItem :exec
INSERT INTO order_items (order_id, product_variant_id, quantity, price)
VALUES ($1, $2, $3, $4);

-- name: DecreaseProductStock :exec
UPDATE product_variants
SET stock_quantity = stock_quantity - $2
WHERE id = $1 AND stock_quantity >= $2;

-- name: DeleteCartItemsByProductIDs :exec
DELETE FROM cart_items
WHERE cart_id IN (
    SELECT id FROM carts WHERE user_id = $1
)
AND product_variant_id = ANY(sqlc.arg(variant_ids)::int[]);
