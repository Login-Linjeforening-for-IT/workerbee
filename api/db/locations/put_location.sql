UPDATE locations
SET
    name_no = :name_no,
    name_en = :name_en,
    type = :type,
    mazemap_campus_id = :mazemap_campus_id,
    mazemap_poi_id = $6,
    address_street = $7,
    address_postcode = $8,
    city_id = $9,
    coordinate_lat = $10,
    coordinate_lon = $11,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
