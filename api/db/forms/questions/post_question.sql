-- Insert a new question
INSERT INTO questions
(
    form_id,
    question_title,
    question_description,
    question_type,
    required,
    position,
    max
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
