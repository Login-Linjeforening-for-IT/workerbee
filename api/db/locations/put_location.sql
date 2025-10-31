UPDATE locations
SET
    name_no = :name_no,
    name_en = :name_en,
    type = :type,
    mazemap_campus_id = :mazemap_campus_id,
    mazemap_poi_id = :mazemap_poi_id,
    address_street = :address_street,
    address_postcode = :address_postcode,
    city_id = :city_id,
    coordinate_lat = :coordinate_lat,
    coordinate_lon = :coordinate_lon,
    url = :url,
    updated_at = NOW()
WHERE id = :id
RETURNING *;
