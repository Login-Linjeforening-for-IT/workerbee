SELECT 
    h.id,
    h.page,
    h.language,
    COUNT(*) OVER() AS total_count
FROM honey h
WHERE h.service = $1
GROUP BY h.id
ORDER BY h.page;