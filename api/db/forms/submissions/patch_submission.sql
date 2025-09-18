UPDATE submissions
SET
    form_id = $1,
    user_id = $2,
    updated_at = NOW()
WHERE
    id = $3;