-- name: delete_event :one
DELETE FROM events 
WHERE id = $1
    AND now() - interval '3 days' < created_at
RETURNING id;