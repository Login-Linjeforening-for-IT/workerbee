SELECT 
    s.id,
    s.submitted_at,
    s.updated_at,
    json_build_object(
        'id', u.id,
        'full_name', u.full_name,
        'email', u.email
    ) AS user,
    COALESCE(json_agg(
        json_build_object(
            'id', q.id,
            'question_title', q.question_title,
            'question_description', q.question_description,
            'question_type', q.question_type,
            'required', q.required,
            'position', q.position,
            'max', q.max,
            'options', COALESCE((
                SELECT json_agg(json_build_object(
                    'id', o.id,
                    'option_text', o.option_text,
                    'position', o.position
                ) ORDER BY o.position)
                FROM question_options o
                WHERE o.question_id = q.id
            ), '[]'),
            'answer', COALESCE((
                SELECT json_agg(json_build_object(
                    'id', a.id,
                    'answer_text', a.answer_text,
                    'selected_options', COALESCE((
                        SELECT json_agg(ao.option_id)
                        FROM answer_options ao
                        WHERE ao.answer_id = a.id
                    ), '[]')
                ))
                FROM answers a
                WHERE a.submission_id = s.id AND a.question_id = q.id
            ), '[]')
        ) ORDER BY q.position
    ) FILTER (WHERE q.id IS NOT NULL), '[]') AS questions
FROM submissions s
JOIN users u ON s.user_id = u.id
JOIN forms f ON s.form_id = f.id
LEFT JOIN questions q ON f.id = q.form_id
WHERE s.id = $1 AND s.form_id = $2
GROUP BY s.id, u.id;