UPDATE albums
SET name_no = :name_no,
    name_en = :name_en,
    description_no = :description_no,
    description_en = :description_en,
    event_id = :event_id,
    year = :year,
    updated_at = NOW()
WHERE id = :id
RETURNING *;