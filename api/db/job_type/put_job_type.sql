UPDATE job_types
SET name_no = :name_no, name_en = :name_en
WHERE id = :id
RETURNING id;