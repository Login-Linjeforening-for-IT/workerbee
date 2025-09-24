-- name: get_categories :many
SELECT *,
    COUNT(*) OVER() AS total_count
FROM categories as c
WHERE
    (
        $1 = '' OR
        to_json(c)::text ILIKE '%' || $1 || '%'
    )
