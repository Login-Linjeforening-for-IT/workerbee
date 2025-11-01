INSERT INTO albums 
( 
    name_en, 
    name_no,
    description_en,
    description_no,
    created_at,
    updated_at
) VALUES 
(
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6
)
RETURNING id;