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
    l.coordinate_lat,
    l.coordinate_lon,
    l.url,
    l.created_at,
    l.updated_at,

    c.id AS "cities.id",
    c.name AS "cities.name",

    COUNT (*) OVER() AS total_count
FROM locations AS l
LEFT JOIN cities c ON c.id = l.city_id
WHERE
    (
        $1 = '' OR
        to_json(l)::text ILIKE '%' || $1 || '%'
    )
    AND (
        cardinality($2::text[]) = 0
        OR l.type::text = ANY($2::text[])
    )
