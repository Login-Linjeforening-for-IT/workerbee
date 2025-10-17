SELECT * 
FROM job_types as jt
WHERE 
    $1 = '' OR
    to_json(jt)::text ILIKE '%' || $1 || '%'