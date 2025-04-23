-- name: UpsertCarts :one
INSERT INTO
    carts (user_id)
VALUES ($1)
ON CONFLICT (user_id) DO UPDATE
SET
    user_id = EXCLUDED.user_id
RETURNING id;

-- name: CheckIfVariantStockSufficient :one
SELECT *
FROM product_variants
WHERE id= $1 AND stock_quantity > $2;

-- name: UpsertToCartItems :exec
INSERT INTO cart_items(cart_id, product_variant_id, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (cart_id, product_variant_id) DO UPDATE
SET quantity = cart_items.quantity + EXCLUDED.quantity;
