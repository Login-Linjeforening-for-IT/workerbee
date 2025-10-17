INSERT INTO job_types (name_no, name_en)
VALUES (:name_no, :name_en)
RETURNING id;