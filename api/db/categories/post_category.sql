INSERT INTO categories
(
    color,
    name_no,
    name_en,
    description_no,
    description_en
)
VALUES (
    :color,
    :name_no,
    :name_en,
    :description_no,
    :description_en
)
RETURNING *;