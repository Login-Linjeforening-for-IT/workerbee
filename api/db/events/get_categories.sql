-- name: get_categories :many
SELECT DISTINCT c.*
FROM events as e
LEFT JOIN categories AS c ON e.category_id = c.id
WHERE 
    e.time_end > NOW() AND
    e.time_publish < NOW() AND
    e.visible = TRUE AND
    e.canceled = FALSE;