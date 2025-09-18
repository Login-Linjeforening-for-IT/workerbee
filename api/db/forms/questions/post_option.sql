INSERT INTO question_options
(
    question_id,
    position,
    option_text,
)
VALUES ($1, $2, $3)
RETURNING *;
