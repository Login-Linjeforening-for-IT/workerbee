-- Update a question
UPDATE questions
SET 
    question_title = $2,
    question_description = $3,
    question_type = $4,
    required = $5,
    position = $6,
    max = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
