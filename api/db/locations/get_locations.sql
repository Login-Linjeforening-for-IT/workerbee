-- name: get_locations :many
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
    l.updated_at,
    COUNT(*) OVER() AS total_count
FROM locations AS l
INNER JOIN cities c ON c.id = l.city_id
WHERE
    (
        $1 = '' OR
        to_json(l)::text ILIKE '%' || $1 || '%'
    )
