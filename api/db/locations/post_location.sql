INSERT INTO locations
(
    name_no,
    name_en,
    type,
    mazemap_campus_id,
    mazemap_poi_id,
    address_street,
    address_postcode,
    city_id,
    coordinate_lat,
    coordinate_lon,
    url
)
VALUES (
    :name_no,
    :name_en,
    :type,
    :mazemap_campus_id,
    :mazemap_poi_id,
    :address_street,
    :address_postcode,
    :city_id,
    :coordinate_lat,
    :coordinate_lon,
    :url
)
RETURNING *;
