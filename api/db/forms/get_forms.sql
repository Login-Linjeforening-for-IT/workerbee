SELECT *,
    COUNT(*) OVER() AS total_count
FROM forms
WHERE 
    (
      $1 = '' OR
      to_json(forms)::text ILIKE '%' || $1 || '%'
    )