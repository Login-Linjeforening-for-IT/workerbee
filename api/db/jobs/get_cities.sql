-- name: get_cities :many
SELECT 
    c.id,
    c.name
FROM cities c
WHERE 
    c.id IN (SELECT city_id FROM ad_city_relation);
