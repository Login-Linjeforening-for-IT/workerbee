SELECT 
    a.id AS id,
    a.name_en AS name_en,
    a.name_no AS name_no,
    a.description_en AS description_en,
    a.description_no AS description_no,
    a.year,
    a.created_at,
    a.updated_at,
    e.id AS "event.id",
    e.name_en AS "event.name_en",
    e.name_no AS "event.name_no",
    e.time_start AS "event.time_start",
    e.time_end AS "event.time_end"
FROM albums AS a
LEFT JOIN events AS e ON a.event_id = e.id
WHERE a.id = $1;