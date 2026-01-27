SELECT 
    *,
    COUNT(*) OVER() AS total_count
FROM quotes
