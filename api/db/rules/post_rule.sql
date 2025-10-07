INSERT INTO rules
(
    name_no,
    name_en,
    description_no,
    description_en
)
VALUES (
    :name_no,
    :name_en,
    :description_no,
    :description_en
)
RETURNING *;