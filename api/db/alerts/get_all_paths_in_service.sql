SELECT 
    a.page,
    array_agg(a.language ORDER BY a.language) AS languages
FROM alerts a
WHERE a.service = $1
GROUP BY a.page
ORDER BY a.page;