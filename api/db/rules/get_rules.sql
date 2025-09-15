SELECT *,
    COUNT(*) OVER() AS total_count
FROM rules
WHERE
    (
        $1 = '' OR
        to_json(rules)::text ILIKE '%' || $1 || '%'
    )
LIMIT $2 OFFSET $3;