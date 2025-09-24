-- name: delete_category :one
DELETE FROM categories 
WHERE id = $1
RETURNING *;