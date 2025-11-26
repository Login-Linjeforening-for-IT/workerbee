-- name: delete_job :one
DELETE FROM jobs 
WHERE id = $1
    AND now() - interval '3 days' < created_at
RETURNING id;