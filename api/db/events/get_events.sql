-- name: GetEvents :many
SELECT e.id,
       e.visible,
       e.name_no,
       e.name_en,
       e.time_type,
       e.time_start,
       e.time_end,
       e.time_publish,
       e.canceled,
       e.link_signup,
       e.capacity,
       e.full,
       c.name_no AS category_name_no,
       c.name_en AS category_name_en,
       l.name_no AS location_name_no,
       l.name_en AS location_name_en,
       e.updated_at,
       (e.deleted_at IS NOT NULL)::bool AS is_deleted,
       a.name_no AS audience_name_no,
       a.name_en AS audience_name_en,
       o.name_no AS organizer_name_no,
       o.name_en AS organizer_name_en
FROM event AS e
INNER JOIN category AS c ON e.category = c.id
LEFT JOIN location AS l ON e.location = l.id
LEFT JOIN audience AS a ON e.audience = a.id
LEFT JOIN organization AS o ON e.organization = o.id
WHERE ($1::bool
       OR ((e.time_end IS NOT NULL AND e.time_end > now())
           OR (e.time_start > now() - interval '1 day')))
ORDER BY e.id
LIMIT $2::int
OFFSET $3::int;
