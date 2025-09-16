-- name: get_event :one
SELECT * FROM events WHERE id = $1;