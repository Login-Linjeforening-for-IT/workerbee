-- name: get_location :one
SELECT
    l.id,
    l.name_no,
    l.name_en,
    l.type,
    l.mazemap_campus_id,
    l.mazemap_poi_id,
    l.address_street,
    l.address_postcode,
    c.name AS city_name,
    l.coordinate_lat,
    l.coordinate_long,
    l.url,
    l.created_at,
    l.updated_at
FROM locations AS l
INNER JOIN cities c ON c.id = l.city_id
WHERE l.id = $1;