-- name: get_cities :many
SELECT 
    c.*,
    COUNT(*) OVER() as total_count
FROM cities as c
WHERE ($1 = '' OR to_json(c)::text ILIKE '%' || $1 || '%')
