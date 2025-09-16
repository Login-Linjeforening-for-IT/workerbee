-- name: delete_rule :one
DELETE FROM rules WHERE id = $1
RETURNING *;