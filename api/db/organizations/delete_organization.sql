-- name: delete_organization :one
DELETE FROM organizations WHERE id = $1
RETURNING *;