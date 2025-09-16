-- name: get_location :one
SELECT * FROM locations WHERE id = $1;