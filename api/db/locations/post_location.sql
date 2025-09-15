INSERT INTO rules 
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
    coordinate_lon
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;
