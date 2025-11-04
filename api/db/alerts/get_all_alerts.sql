SELECT a.*, COUNT (*) OVER() AS total_count
FROM alerts as a
WHERE 
    (
        $1 = '' OR
        to_json(a)::text ILIKE '%' || $1 || '%'
    )