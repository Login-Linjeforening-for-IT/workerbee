SELECT 
    c.id, 
    c.name_en,
    COUNT(e.id) as event_count
FROM categories c
LEFT JOIN events e ON e.category_id = c.id
  AND e.created_at >= now() - interval '3 months'
GROUP BY c.id
ORDER BY event_count DESC
LIMIT 1;