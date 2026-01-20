SELECT 
    j.time_publish
FROM 
    jobs j
WHERE 
    j.time_publish > NOW()
ORDER BY 
    j.time_publish ASC
LIMIT 1;