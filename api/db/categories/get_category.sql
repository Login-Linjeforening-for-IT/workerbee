-- name: get_category :many
SELECT *,
    COUNT(*) OVER() AS total_count
FROM categories
WHERE
    (
        $1 = '' OR
        to_json(categories)::text ILIKE '%' || $1 || '%'
    )
LIMIT $2 OFFSET $3;