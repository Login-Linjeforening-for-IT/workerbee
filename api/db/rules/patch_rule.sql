UPDATE rules 
SET 
    name_no = $2,
    name_en = $3,
    description_no = $4,
    description_en = $5
WHERE id = $1
RETURNING *;