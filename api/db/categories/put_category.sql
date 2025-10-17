UPDATE categories
SET
    name_no = :name_no,
    name_en = :name_en,
    color = :color
WHERE id = :id
RETURNING *;