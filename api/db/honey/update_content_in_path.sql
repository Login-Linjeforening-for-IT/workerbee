UPDATE honey
SET 
    text = :text,
    service = :service,
    language = :language,
    page = :page
WHERE id = :id
RETURNING *;