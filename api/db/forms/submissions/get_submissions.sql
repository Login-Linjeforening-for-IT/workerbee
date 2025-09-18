SELECT *,
    COUNT(*) OVER() AS total_count
FROM submissions
WHERE 
    (
      $1 = '' OR
      to_json(submissions)::text ILIKE '%' || $1 || '%'
    )