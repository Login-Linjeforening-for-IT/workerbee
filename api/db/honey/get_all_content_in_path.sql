SELECT 
    h.language,
    h.text
FROM honey h
WHERE 
    h.service = $1 AND h.page = $2
ORDER BY h.language;