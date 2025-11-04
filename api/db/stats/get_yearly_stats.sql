SELECT activity.*
FROM daily_history AS activity
WHERE activity.insert_date >= CURRENT_DATE - INTERVAL '1 year'
ORDER BY activity.insert_date DESC;