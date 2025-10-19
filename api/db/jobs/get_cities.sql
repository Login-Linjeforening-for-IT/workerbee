-- name: get_cities :many
SELECT 
    c.id,
    c.name
FROM cities c
WHERE 
    c.id IN (SELECT city_id FROM ad_city_relation WHERE job_id IN 
        (SELECT id FROM jobs WHERE visible = true AND time_expire > NOW() AND time_publish < NOW())
    );
