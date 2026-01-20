-- name: GetNextPublishTime :one
SELECT 
    e.time_publish
FROM 
    events e
WHERE 
    e.time_publish > NOW()
ORDER BY 
    e.time_publish ASC
LIMIT 1;
