SELECT 
    h.service,
    h.language,
    h.page,
    h.text
FROM honey h
WHERE 
    h.service = $1 AND 
    h.page = $2 AND 
    h.language = $3;