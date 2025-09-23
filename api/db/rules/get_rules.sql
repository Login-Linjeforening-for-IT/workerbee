-- name: get_rules :many
SELECT *,
    COUNT(*) OVER() AS total_count
FROM rules as r
WHERE
    (
        $1 = '' OR
        to_json(r)::text ILIKE '%' || $1 || '%'
    )
