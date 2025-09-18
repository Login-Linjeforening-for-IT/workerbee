SELECT 
	f.id,
	f.title,
	f.description,
	f.capacity,
	f.open_at,
	f.close_at,
	f.created_at,
	f.updated_at,
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
			), '[]')
		)
	) FILTER (WHERE q.id IS NOT NULL), '[]') AS questions
FROM forms f
JOIN users u ON f.user_id = u.id
LEFT JOIN questions q ON f.id = q.form_id
WHERE f.id = $1
GROUP BY f.id, u.id;