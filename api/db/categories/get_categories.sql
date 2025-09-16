-- name: get_categories :many
SELECT * FROM categories WHERE id = $1;