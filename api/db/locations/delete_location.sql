-- name: delete_location :one
DELETE FROM locations WHERE id = $1
RETURNING id;
