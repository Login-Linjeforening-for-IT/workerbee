-- name: get_event :one
SELECT
    e.id,
    e.name_en,
    e.name_no,
    e.time_start,
    e.time_end,
    c.id AS "category.id",
    c.name_no AS "category.name_no",
    c.name_en AS "category.name_en",
    l.id AS "location.id",
    l.name_no AS "location.name_no",
    l.name_en AS "location.name_en",
    o.id AS "organization.id",
    o.name_no AS "organization.name_no",
    o.name_en AS "organization.name_en"
FROM events AS e
LEFT JOIN categories AS c ON e.category_id = c.id
LEFT JOIN locations AS l ON e.location_id = l.id
LEFT JOIN organizations AS o ON e.organization_id = o.id
WHERE e.id = $1;