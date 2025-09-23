-- name: get_organizations :many
SELECT *,
    COUNT(*) OVER() AS total_count
FROM organizations as o
WHERE
    (
        $1 = '' OR
        to_json(o)::text ILIKE '%' || $1 || '%'
    )
