UPDATE alerts
SET service = :service,
    page = :page,
    title_en = :title_en,
    title_no = :title_no,
    description_en = :description_en,
    description_no = :description_no
WHERE id = :id
RETURNING *;