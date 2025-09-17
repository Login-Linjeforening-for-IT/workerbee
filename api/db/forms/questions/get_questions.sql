SELECT 
    q.*,
    COALESCE(array_agg(o.id) FILTER (WHERE o.id IS NOT NULL), '{}') AS option_ids,
    COALESCE(array_agg(o.option_text) FILTER (WHERE o.option_text IS NOT NULL), '{}') AS option_texts
FROM questions q
LEFT JOIN question_options o ON q.id = o.question_id
WHERE q.form_id = $1
GROUP BY q.id
ORDER BY q.position;