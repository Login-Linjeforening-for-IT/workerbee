-- name: get_organizations :many
SELECT *,
    COUNT(*) OVER() AS total_count
FROM organizations
WHERE
    (
        $1 = '' OR
        to_json(organizations)::text ILIKE '%' || $1 || '%'
    )
LIMIT $2 OFFSET $3;