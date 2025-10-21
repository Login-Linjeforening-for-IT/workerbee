SELECT 
    h.page,
    array_agg(h.language ORDER BY h.language) AS languages
FROM honey h
WHERE h.service = $1
GROUP BY h.page
ORDER BY h.page;