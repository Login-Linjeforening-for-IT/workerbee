-- name: get_categories :many
SELECT DISTINCT e.category
FROM events as e
WHERE 
    e.time_end > NOW() AND
    e.time_publish < NOW() AND
    e.canceled = FALSE;
    
