UPDATE question_options
SET 
    option_text = $2,
    position = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;