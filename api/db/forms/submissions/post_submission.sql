INSERT INTO submissions 
(
    form_id,
    user_id
) 
VALUES ($1, $2)
RETURNING *;