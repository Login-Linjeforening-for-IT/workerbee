INSERT INTO audiences (name_no, name_en)
VALUES (:name_no, :name_en)
RETURNING *;