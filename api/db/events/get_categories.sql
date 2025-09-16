-- name: get_categories :many
SELECT DISTINCT c.name_no, c.name_en
FROM categories c
JOIN events e ON e.category_id = c.id
WHERE 
    e.time_end > NOW() AND
    e.time_publish < NOW() AND
    e.canceled = FALSE;
