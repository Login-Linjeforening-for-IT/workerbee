-- name: get_locations :many
SELECT *,
    COUNT(*) OVER() AS total_count
FROM locations
WHERE
    (
        $1 = '' OR
        to_json(locations)::text ILIKE '%' || $1 || '%'
    )
LIMIT $2 OFFSET $3;