-- name: delete_event :one
DELETE FROM events 
WHERE id = $1
    AND now() - interval '3 days' < time_end
RETURNING id;