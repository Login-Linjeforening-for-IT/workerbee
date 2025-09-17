INSERT INTO forms 
(
    user_id,
    title,
    description,
    capacity,
    open_at,
    close_at
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
