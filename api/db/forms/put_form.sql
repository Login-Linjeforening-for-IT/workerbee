UPDATE forms
SET 
    title = $2,
    description = $3,
    capacity = $4,
    open_at = $5,
    close_at = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
