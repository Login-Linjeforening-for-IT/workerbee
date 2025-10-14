-- name: delete_event :one
DELETE FROM events WHERE id = $1
RETURNING id;