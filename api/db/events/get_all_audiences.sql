SELECT DISTINCT a.*
FROM events as e
LEFT JOIN audiences as a ON e.audience_id = a.id
WHERE 
    e.time_end > NOW() AND
    e.time_publish < NOW() AND
    e.visible = TRUE AND
    e.canceled = FALSE;