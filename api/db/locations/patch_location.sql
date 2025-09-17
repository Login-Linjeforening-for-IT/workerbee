-- name: patch_location :one
UPDATE locations
SET
    name_no = $2,
    name_en = $3,
    type = $4,
    mazemap_campus_id = $5,
    mazemap_poi_id = $6,
    address_street = $7,
    address_postcode = $8,
    city_id = $9,
    coordinate_lat = $10,
    coordinate_lon = $11,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
