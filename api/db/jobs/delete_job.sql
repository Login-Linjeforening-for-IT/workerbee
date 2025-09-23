-- name: delete_job :one
DELETE FROM jobs 
WHERE id = $1
RETURNING *;