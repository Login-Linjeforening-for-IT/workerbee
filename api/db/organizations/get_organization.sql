-- name: get_organization :one
SELECT * FROM organizations WHERE id = $1;