-- name: GetProducts :many
SELECT id, name, base_price
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

-- name: GetProductDetails :one
SELECT
    p.id,
    p.name,
    p.slug,
    p.desciprtion,
    p.main_image_url,
    p.base_price,
    c.id AS category_id,
    c.name AS category_name,
    b.id AS brand_id,
    b.name AS brand_name
FROM
    products p
LEFT JOIN
    categories c ON p.category_id = c.id
LEFT JOIN
    brands b ON p.brand_id = b.id
WHERE
    p.id = $1
GROUP BY
    p.id, c.id, c.name, b.id, b.name
;

-- name: GetProductVariants :many
SELECT
	pv.id,
	pv.product_id,
	pv.sku,
	pv.price,
	pv.stock_quantity,
	pv.sold,
	pv.image_url,
	pv.is_default
FROM
	product_variants pv
	LEFT JOIN variant_attribute_values AS vav ON pv.id = vav.variant_id
	LEFT JOIN attribute_values AS av ON vav.value_id = av.id
	LEFT JOIN "attributes" AS a ON av.attribute_id = a.id
WHERE product_id = $1
GROUP BY pv.id
;
