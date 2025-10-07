UPDATE categories
SET 
    color = :color,
    name_no = :name_no,
    name_en = :name_en,
    description_no = :description_no,
    description_en = :description_en,
    updated_at = NOW()
WHERE id = :id
RETURNING *;