SELECT *,
  COUNT(*) OVER() AS total_count
FROM forms
WHERE deleted_at IS NULL
  AND (
    $1 = '' OR
    title ILIKE '%' || $1 || '%' OR
    description ILIKE '%' || $1 || '%'
  )
LIMIT $2 OFFSET $3;