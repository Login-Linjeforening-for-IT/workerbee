SELECT a.*,
    COUNT(*) OVER() AS total_count
FROM audiences AS a
WHERE 
    $1 = '' OR to_json(a)::text ILIKE '%' || $1 || '%'