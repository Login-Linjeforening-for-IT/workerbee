-- name: get_category :many
SELECT * 
FROM categories
WHERE id = $1;