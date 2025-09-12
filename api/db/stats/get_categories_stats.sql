-- name: GetCategoriesStats :many
WITH counts AS (
  SELECT
    c.id AS category_id,
    c.name_en,
    c.name_no,
    COUNT(e.*) AS event_count
  FROM category c
  LEFT JOIN event e ON e.category = c.id
    AND e.time_start >= now() - interval '3 months'
    AND e.time_start <= now()
  GROUP BY c.id, c.name_en, c.name_no
)
SELECT * FROM counts
WHERE event_count > 0
ORDER BY event_count DESC;