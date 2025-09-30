-- name: get_event :one
SELECT e.id,
    e.visible,
    e.name_no,
    e.name_en,
    e.description_no,
    e.description_en,
    e.informational_no,
    e.informational_en,
    e.time_type,
    e.time_start,
    e.time_end,
    e.time_publish,
    e.canceled,
    e.link_signup,
    e.capacity,
    e.is_full,
    e.parent_id,
    e.category_id,
    e.location_id,
    e.organization_id,
    c.name_no AS category_name_no,
    c.name_en AS category_name_en,
    l.name_no AS location_name_no,
    l.name_en AS location_name_en,
    e.updated_at,
    e.created_at,
    a.name_no AS audience_name_no,
    a.name_en AS audience_name_en,
    o.name_no AS organizer_name_no,
    o.name_en AS organizer_name_en
FROM events AS e
    INNER JOIN categories AS c ON e.category_id = c.id
    LEFT JOIN locations AS l ON e.location_id = l.id
    LEFT JOIN audiences AS a ON e.audience_id = a.id
    LEFT JOIN organizations AS o ON e.organization_id = o.id
WHERE e.id = $1;