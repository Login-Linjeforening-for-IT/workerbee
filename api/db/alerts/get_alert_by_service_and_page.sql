SELECT 
    a.*
FROM alerts a
WHERE 
    a.service = $1 AND a.page = $2;