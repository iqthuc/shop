-- name: GetProducts :many
SELECT id, name, base_price::float8
FROM products
WHERE (@key_word::text = '' OR name ILIKE '%' || @key_word || '%')
ORDER BY
    CASE WHEN sqlc.arg(sort_column)::varchar = 'price' AND sqlc.arg(sort_direction)::varchar = 'asc' THEN base_price END ASC,
    CASE WHEN sqlc.arg(sort_column)::varchar = 'price' AND sqlc.arg(sort_direction)::varchar = 'desc' THEN base_price END DESC,
    CASE WHEN sqlc.arg(sort_column)::varchar = 'name' AND sqlc.arg(sort_direction)::varchar = 'asc' THEN name END ASC,
    CASE WHEN sqlc.arg(sort_column)::varchar = 'name' AND sqlc.arg(sort_direction)::varchar = 'desc' THEN name END DESC
LIMIT $1 OFFSET $2;

-- name: GetProductsCount :one
SELECT COUNT(*) FROM products;
